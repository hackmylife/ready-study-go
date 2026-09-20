package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDecodeGoTestResults(t *testing.T) {
	file, err := os.Open("testdata/results.jsonl")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	results, err := decodeResults(file)
	if err != nil {
		t.Fatal(err)
	}
	for _, tt := range []struct {
		pkg    string
		passed bool
		failed bool
	}{
		{"koans-fixture/pass", true, false},
		{"koans-fixture/fail", false, true},
		{"koans-fixture/skip", false, false},
		{"koans-fixture/empty", false, false},
	} {
		t.Run(tt.pkg, func(t *testing.T) {
			r := results[tt.pkg]
			if r == nil {
				t.Fatal("missing package result")
			}
			if passed(r) != tt.passed || r.Failed != tt.failed {
				t.Fatalf("pass=%v, failed=%v; want pass=%v, failed=%v", passed(r), r.Failed, tt.passed, tt.failed)
			}
		})
	}
	if passed(nil) {
		t.Fatal("missing results cannot pass")
	}
}

func TestOverlayLeavesLearnerFilesUntouched(t *testing.T) {
	root := t.TempDir()
	write := func(name, content string) {
		t.Helper()
		path := filepath.Join(root, name)
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			t.Fatal(err)
		}
	}
	write("01/one/exercise.go", "learner answer\n")
	write("01/one/utils/helper.go", "old helper\n")
	write("solutions/01/one/solution.go", "//go:build ignore\n\npackage koan\n")
	write("solutions/01/one/explanation.md", "why\n")
	dir, err := scratch(root)
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)
	if err := overlay(root, dir, exercise{Path: "01/one", Remove: []string{"utils"}}); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "01/one/exercise.go"))
	if err != nil || string(data) != "package koan\n" {
		t.Fatalf("overlay=%q, err=%v", data, err)
	}
	if _, err := os.Stat(filepath.Join(dir, "01/one/utils")); !os.IsNotExist(err) {
		t.Fatal("removed package remains in scratch")
	}
	data, err = os.ReadFile(filepath.Join(root, "01/one/exercise.go"))
	if err != nil || string(data) != "learner answer\n" {
		t.Fatal("learner answer changed")
	}
	if _, err := os.Stat(filepath.Join(root, "01/one/utils/helper.go")); err != nil {
		t.Fatal("learner helper was deleted")
	}
}

func TestCatalogRejectsDuplicateAndEscapingPaths(t *testing.T) {
	for _, content := range []string{
		`[]`,
		`[{"path":"../outside"}]`,
		`[{"path":"01/one"},{"path":"01/one"}]`,
		`[{"path":"01/one","remove":["../two"]}]`,
	} {
		root := t.TempDir()
		if err := os.WriteFile(filepath.Join(root, "curriculum.json"), []byte(content), 0644); err != nil {
			t.Fatal(err)
		}
		if _, err := catalog(root); err == nil {
			t.Fatalf("accepted invalid catalog: %s", content)
		}
	}
}

func TestDecodeResultsRejectsMalformedStream(t *testing.T) {
	if _, err := decodeResults(strings.NewReader("not JSON")); err == nil {
		t.Fatal("invalid go test output accepted")
	}
}
