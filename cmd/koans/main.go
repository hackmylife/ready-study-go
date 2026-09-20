package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type exercise struct {
	Path     string   `json:"path"`
	Goal     string   `json:"goal"`
	Database bool     `json:"database"`
	Kind     string   `json:"kind"`
	Remove   []string `json:"remove,omitempty"`
	Race     bool     `json:"race,omitempty"`
}

type event struct {
	Action  string
	Package string
	Test    string
	Output  string
}

type result struct {
	Action string
	Ran    bool
	Passed bool
	Failed bool
	Output strings.Builder
}

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string) error {
	root, err := findRoot()
	if err != nil {
		return err
	}
	items, err := catalog(root)
	if err != nil {
		return err
	}
	if len(args) == 0 {
		fmt.Println("koans list | check PATH | lint PATH | diff PATH | progress | next | hint PATH 1..3 | verify [--starters]")
		return nil
	}
	switch args[0] {
	case "list":
		for _, item := range items {
			suffix := ""
			if item.Database {
				suffix = " [PostgreSQL]"
			}
			fmt.Printf("%s  %s%s\n", item.Path, item.Goal, suffix)
		}
		return nil
	case "check", "lint", "diff", "hint":
		if len(args) < 2 {
			return errors.New("演習のパスを指定してください")
		}
		item, err := lookup(items, strings.TrimPrefix(strings.TrimSuffix(args[1], "/"), "./"))
		if err != nil {
			return err
		}
		if args[0] != "hint" && len(args) != 2 {
			return fmt.Errorf("%s PATH を指定してください", args[0])
		}
		if args[0] == "lint" {
			return lintExercise(root, item)
		}
		if args[0] == "diff" {
			changed, err := solutionDiff(root, item, os.Stdout)
			if err != nil {
				return err
			}
			if !changed {
				fmt.Println("整形後のコードは模範解答と同じです。")
			}
			fmt.Printf("差分の - は自分のコード、+ は模範解答です。違いがあっても不正解とは限りません。\n解説: solutions/%s/explanation.md\n", item.Path)
			return nil
		}
		if args[0] == "hint" {
			if len(args) != 3 {
				return errors.New("hint PATH 1..3 を指定してください")
			}
			n, err := strconv.Atoi(args[2])
			if err != nil || n < 1 || n > 3 {
				return errors.New("ヒント番号は1〜3です")
			}
			data, err := os.ReadFile(filepath.Join(root, item.Path, "HINTS.md"))
			if err != nil {
				return err
			}
			parts := strings.Split(string(data), "</summary>")
			if len(parts) <= n {
				return errors.New("ヒント形式が不正です")
			}
			fmt.Println(strings.TrimSpace(strings.SplitN(parts[n], "</details>", 2)[0]))
			return nil
		}
		if item.Database && os.Getenv("KOANS_DATABASE_URL") == "" {
			return errors.New("PostgreSQLが必要です。docs/database.md を参照し KOANS_DATABASE_URL を設定してください")
		}
		results, testErr := testResults(root, []exercise{item})
		for _, result := range results {
			fmt.Print(result.Output.String())
		}
		if testErr != nil {
			return testErr
		}
		if !passed(results["ready-study-go/"+item.Path]) {
			return errors.New("成功したテストがありません。テストの削除や全件skipを確認してください")
		}
		if item.Kind == "testing" {
			if err := checkMutants(root, item); err != nil {
				return err
			}
		}
		if err := lintExercise(root, item); err != nil {
			return err
		}
		fmt.Printf("PASS — テストとlintが成功しました。\n模範解答と比較する: go run ./cmd/koans diff %s\n", item.Path)
		return nil
	case "progress", "next":
		selected := available(items)
		results, testErr := testResults(root, selected)
		if len(results) == 0 && testErr != nil {
			return testErr
		}
		completed, waiting := 0, 0
		for _, item := range items {
			status := "TODO"
			r := results["ready-study-go/"+item.Path]
			if item.Database && os.Getenv("KOANS_DATABASE_URL") == "" {
				status = "NEEDS_DB"
				waiting++
			} else if passed(r) {
				status = "PASS"
				if item.Kind == "testing" {
					if err := checkMutants(root, item); err != nil {
						status = "TESTS_WEAK"
					}
				}
				if status == "PASS" {
					completed++
				}
			} else if r == nil || !r.Ran || r.Action == "pass" {
				status = "BUILD_ERROR"
			}
			if args[0] == "next" {
				if status != "PASS" && status != "NEEDS_DB" {
					fmt.Printf("%s\n%s\ngo run ./cmd/koans check %s\n", item.Path, item.Goal, item.Path)
					return nil
				}
			} else {
				fmt.Printf("%-12s %s\n", status, item.Path)
			}
		}
		fmt.Printf("実測: %d/%d 演習PASS、%d 演習はDB未接続\n", completed, len(items), waiting)
		return nil
	case "verify":
		if len(args) > 2 || (len(args) == 2 && args[1] != "--starters") {
			return errors.New("verify [--starters]")
		}
		return verify(root, items, len(args) == 2)
	default:
		return fmt.Errorf("不明なコマンド: %s", args[0])
	}
}

func findRoot() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(path, "curriculum.json")); err == nil {
			return path, nil
		}
		parent := filepath.Dir(path)
		if parent == path {
			return "", errors.New("curriculum.jsonがあるリポジトリ内で実行してください")
		}
		path = parent
	}
}

func catalog(root string) ([]exercise, error) {
	data, err := os.ReadFile(filepath.Join(root, "curriculum.json"))
	if err != nil {
		return nil, err
	}
	var items []exercise
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, err
	}
	seen := make(map[string]bool)
	for _, item := range items {
		if !filepath.IsLocal(item.Path) || strings.Contains(item.Path, "\\") || seen[item.Path] {
			return nil, fmt.Errorf("不正または重複した演習パス: %q", item.Path)
		}
		seen[item.Path] = true
		for _, name := range item.Remove {
			if !filepath.IsLocal(name) {
				return nil, fmt.Errorf("不正な削除パス: %q", name)
			}
		}
	}
	if len(items) == 0 {
		return nil, errors.New("演習一覧が空です")
	}
	return items, nil
}

func lookup(items []exercise, path string) (exercise, error) {
	for _, item := range items {
		if item.Path == path {
			return item, nil
		}
	}
	return exercise{}, fmt.Errorf("演習が見つかりません: %q (listで確認できます)", path)
}

func available(items []exercise) []exercise {
	var out []exercise
	for _, item := range items {
		if !item.Database || os.Getenv("KOANS_DATABASE_URL") != "" {
			out = append(out, item)
		}
	}
	return out
}

func passed(r *result) bool {
	return r != nil && r.Action == "pass" && r.Passed && !r.Failed
}

func command(dir, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s %s: %w", name, strings.Join(args, " "), err)
	}
	return nil
}

func testResults(root string, items []exercise) (map[string]*result, error) {
	args := []string{"test", "-json", "-race", "-count=1", "-timeout=15s"}
	for _, item := range items {
		args = append(args, "./"+item.Path+"/...")
	}
	cmd := exec.Command("go", args...)
	cmd.Dir = root
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	pipe, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	results, decodeErr := decodeResults(pipe)
	err = cmd.Wait()
	if decodeErr != nil {
		return results, decodeErr
	}
	if err != nil {
		return results, fmt.Errorf("go test: %w\n%s", err, stderr.String())
	}
	return results, nil
}

func decodeResults(reader io.Reader) (map[string]*result, error) {
	results := make(map[string]*result)
	decoder := json.NewDecoder(reader)
	for {
		var ev event
		if err := decoder.Decode(&ev); err != nil {
			if errors.Is(err, io.EOF) {
				return results, nil
			}
			return results, err
		}
		r := results[ev.Package]
		if r == nil {
			r = &result{}
			results[ev.Package] = r
		}
		if ev.Test != "" && ev.Action == "run" {
			r.Ran = true
		}
		if ev.Test != "" && ev.Action == "fail" {
			r.Failed = true
		}
		if ev.Test != "" && ev.Action == "pass" {
			r.Passed = true
		}
		if ev.Test == "" && (ev.Action == "pass" || ev.Action == "fail" || ev.Action == "skip") {
			r.Action = ev.Action
		}
		r.Output.WriteString(ev.Output)
	}
}

func scratch(root string) (string, error) {
	dir, err := os.MkdirTemp("", "go-koans-")
	if err != nil {
		return "", err
	}
	err = filepath.WalkDir(root, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		if rel == "." {
			return nil
		}
		if entry.IsDir() && (strings.HasPrefix(entry.Name(), ".") || rel == "solutions") {
			return filepath.SkipDir
		}
		if entry.Type()&os.ModeSymlink != 0 {
			return fmt.Errorf("検証コピーではsymlinkを扱いません: %s", rel)
		}
		target := filepath.Join(dir, rel)
		if entry.IsDir() {
			return os.MkdirAll(target, 0755)
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(target, data, 0644)
	})
	if err != nil {
		os.RemoveAll(dir)
		return "", err
	}
	return dir, nil
}

func overlay(root, dir string, item exercise) error {
	source := filepath.Join(root, "solutions", item.Path)
	err := filepath.WalkDir(source, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}
		rel, err := filepath.Rel(source, path)
		if err != nil {
			return err
		}
		rel = solutionTarget(rel)
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		data = bytes.TrimPrefix(data, []byte("//go:build ignore\n\n"))
		target := filepath.Join(dir, item.Path, rel)
		if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
			return err
		}
		return os.WriteFile(target, data, 0644)
	})
	if err != nil {
		return err
	}
	for _, name := range item.Remove {
		if err := os.RemoveAll(filepath.Join(dir, item.Path, name)); err != nil {
			return err
		}
	}
	return nil
}

func checkMutants(root string, item exercise) error {
	dir, err := scratch(root)
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)
	paths, err := filepath.Glob(filepath.Join(root, item.Path, "mutants", "*.go.txt"))
	if err != nil {
		return err
	}
	if len(paths) == 0 {
		return fmt.Errorf("%s: 欠陥実装がありません", item.Path)
	}
	for _, path := range paths {
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if err := os.WriteFile(filepath.Join(dir, item.Path, "exercise.go"), data, 0644); err != nil {
			return err
		}
		results, testErr := testResults(dir, []exercise{item})
		r := results["ready-study-go/"+item.Path]
		if testErr == nil || r == nil || !r.Failed || strings.Contains(r.Output.String(), "panic:") {
			return fmt.Errorf("%s: %sをテストのassertionで検出できません\n%v", item.Path, filepath.Base(path), testErr)
		}
	}
	fmt.Printf("欠陥検出PASS: %s (実測 %d件)\n", item.Path, len(paths))
	return nil
}

func verify(root string, items []exercise, starters bool) error {
	if os.Getenv("KOANS_DATABASE_URL") == "" {
		fmt.Println("DB未接続: PostgreSQL演習はSKIP。全体の検証にはdocs/database.mdの設定が必要です。")
	}
	dir, err := scratch(root)
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)
	for _, item := range items {
		if err := overlay(root, dir, item); err != nil {
			return fmt.Errorf("%s: %w", item.Path, err)
		}
	}
	if err := command(dir, "go", "vet", "./..."); err != nil {
		return err
	}
	if err := command(dir, "go", "tool", "staticcheck", "./..."); err != nil {
		return err
	}
	if err := command(dir, "go", "test", "-race", "-count=1", "-timeout=90s", "./..."); err != nil {
		return err
	}
	for _, item := range items {
		if item.Kind == "testing" {
			if err := checkMutants(dir, item); err != nil {
				return err
			}
		}
	}
	if starters {
		selected := available(items)
		results, testErr := testResults(root, selected)
		var invalid []string
		for _, item := range selected {
			r := results["ready-study-go/"+item.Path]
			if r == nil || r.Action != "fail" || !r.Failed {
				invalid = append(invalid, item.Path)
			}
		}
		if len(invalid) > 0 {
			sort.Strings(invalid)
			return fmt.Errorf("未着手状態でテストの失敗を確認できない演習:\n%s\n%v", strings.Join(invalid, "\n"), testErr)
		}
		fmt.Printf("未着手テストの失敗を確認: 実測 %d演習\n", len(selected))
	}
	fmt.Println("模範解答・vet・staticcheck・race・Testing演習の欠陥検出を検証しました。学習者のファイルは変更していません。")
	return nil
}
