package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func writeReviewFixture(t *testing.T, root, name string, data []byte) {
	t.Helper()
	path := filepath.Join(root, name)
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, data, 0644); err != nil {
		t.Fatal(err)
	}
}

func readReviewFixture(t *testing.T, path string) []byte {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return data
}

func TestSolutionDiff(t *testing.T) {
	for _, tt := range []struct {
		path, solution, target, fixture string
	}{
		{"01-language/01-variables", "solution.go", "exercise.go", "swap.go.txt"},
		{"07-testing/01-basic-test", "solution_test.go", "exercise_test.go", "abs_test.go.txt"},
	} {
		t.Run(tt.path, func(t *testing.T) {
			root := t.TempDir()
			item := exercise{Path: tt.path}
			answer := readReviewFixture(t, filepath.Join("testdata/review", tt.fixture))
			writeReviewFixture(t, root, filepath.Join("solutions", tt.path, tt.solution), answer)
			// Use the real solution with different formatting as the learner's answer.
			own := bytes.TrimPrefix(answer, []byte("//go:build ignore\n\n"))
			own = bytes.Replace(own, []byte("package "), []byte("package    "), 1)
			target := filepath.Join(tt.path, tt.target)
			writeReviewFixture(t, root, target, own)
			var out bytes.Buffer
			changed, err := solutionDiff(root, item, &out)
			if err != nil || changed || out.Len() != 0 {
				t.Fatalf("format-only difference: changed=%v, err=%v, diff=%s", changed, err, &out)
			}
			own = append(own, []byte("\n// learner's note\n")...)
			writeReviewFixture(t, root, target, own)
			changed, err = solutionDiff(root, item, &out)
			if err != nil || !changed || !strings.Contains(out.String(), "-// learner's note") || !strings.Contains(out.String(), "+++ solution/"+tt.target) {
				t.Fatalf("changed=%v, err=%v, diff=%s", changed, err, &out)
			}
			if !bytes.Equal(readReviewFixture(t, filepath.Join(root, target)), own) {
				t.Fatal("diff modified learner's file")
			}
			if !bytes.Equal(readReviewFixture(t, filepath.Join(root, "solutions", tt.path, tt.solution)), answer) {
				t.Fatal("diff modified solution")
			}
		})
	}
}

func TestSolutionDiffAddedAndRemovedFiles(t *testing.T) {
	root := t.TempDir()
	item := exercise{Path: "05-idiomatic-go/10-remove-utils", Remove: []string{"utils", "already-removed"}}
	answer := readReviewFixture(t, "testdata/review/normalize.go.txt")
	writeReviewFixture(t, root, filepath.Join("solutions", item.Path, "solution.go"), answer)
	writeReviewFixture(t, root, filepath.Join(item.Path, "utils/old.go"), []byte("package utils\n"))
	var out bytes.Buffer
	changed, err := solutionDiff(root, item, &out)
	if err != nil || !changed {
		t.Fatalf("changed=%v, err=%v", changed, err)
	}
	for _, want := range []string{"+++ solution/exercise.go", "--- your/utils/old.go", "-package utils", "/dev/null"} {
		if !strings.Contains(out.String(), want) {
			t.Fatalf("missing %q in diff:\n%s", want, &out)
		}
	}
	if _, err := os.Stat(filepath.Join(root, item.Path, "utils/old.go")); err != nil {
		t.Fatal("diff deleted learner's helper:", err)
	}
}

func TestCheckFormatting(t *testing.T) {
	root := t.TempDir()
	writeReviewFixture(t, root, "exercise.go", []byte("package koan\n"))
	writeReviewFixture(t, root, "nested/exercise_test.go", []byte("package    koan\n"))
	writeReviewFixture(t, root, "mutants/broken.go.txt", []byte("not valid Go"))
	if err := checkFormatting(root); err == nil || !strings.Contains(err.Error(), "nested/exercise_test.go") {
		t.Fatalf("unformatted test file not reported: %v", err)
	}
	if got := string(readReviewFixture(t, filepath.Join(root, "nested/exercise_test.go"))); got != "package    koan\n" {
		t.Fatal("format check modified file")
	}
	writeReviewFixture(t, root, "nested/exercise_test.go", []byte("package koan\n"))
	if err := checkFormatting(root); err != nil {
		t.Fatal(err)
	}
	writeReviewFixture(t, root, "exercise.go", []byte("package koan\nfunc broken("))
	if err := checkFormatting(root); err == nil {
		t.Fatal("invalid syntax accepted")
	}
}

func TestCheckAndLint(t *testing.T) {
	repo, err := filepath.Abs("../..")
	if err != nil {
		t.Fatal(err)
	}
	root := t.TempDir()
	path := "01-language/01-variables"
	for _, name := range []string{"go.mod", "go.sum", "curriculum.json", path + "/exercise_test.go"} {
		writeReviewFixture(t, root, name, readReviewFixture(t, filepath.Join(repo, name)))
	}
	answer := bytes.TrimPrefix(readReviewFixture(t, "testdata/review/swap.go.txt"), []byte("//go:build ignore\n\n"))
	writeReviewFixture(t, root, path+"/exercise.go", answer)
	// An unrelated broken exercise must not prevent checking this answer.
	writeReviewFixture(t, root, "02-data/01-string/exercise.go", []byte("invalid Go"))
	t.Chdir(root)
	if err := run([]string{"check", path}); err != nil {
		t.Fatal("correct answer failed:", err)
	}
	writeReviewFixture(t, root, path+"/exercise.go", bytes.Replace(answer, []byte("package "), []byte("package    "), 1))
	if err := run([]string{"check", path}); err == nil || !strings.Contains(err.Error(), "gofmt") {
		t.Fatalf("check did not reject passing but unformatted answer: %v", err)
	}
	writeReviewFixture(t, root, path+"/exercise.go", answer)
	for _, tt := range []struct {
		name, code, diagnostic string
	}{
		{"vet", "package koan\n\nimport \"fmt\"\n\nfunc Print() { fmt.Printf(\"%d\", \"text\") }\n", "go vet "},
		{"staticcheck", "package koan\n\nfunc unused() {}\n", "go tool staticcheck "},
	} {
		t.Run(tt.name, func(t *testing.T) {
			writeReviewFixture(t, root, path+"/extra.go", []byte(tt.code))
			if err := run([]string{"lint", path}); err == nil || !strings.Contains(err.Error(), tt.diagnostic) {
				t.Fatalf("lint did not report %s failure: %v", tt.name, err)
			}
		})
	}
}
