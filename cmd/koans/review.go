package main

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func lintExercise(root string, item exercise) error {
	formatErr := checkFormatting(filepath.Join(root, item.Path))
	pattern := "./" + item.Path + "/..."
	vetErr := command(root, "go", "vet", pattern)
	staticErr := command(root, "go", "tool", "staticcheck", pattern)
	if err := errors.Join(formatErr, vetErr, staticErr); err != nil {
		return err
	}
	fmt.Println("LINT PASS — gofmt・go vet・staticcheck")
	return nil
}

func checkFormatting(dir string) error {
	var problems []error
	err := filepath.WalkDir(dir, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		formatted, err := format.Source(data)
		if err != nil {
			problems = append(problems, fmt.Errorf("%s: %w", path, err))
		} else if !bytes.Equal(data, formatted) {
			problems = append(problems, fmt.Errorf("gofmtが必要です: %s", path))
		}
		return nil
	})
	return errors.Join(append(problems, err)...)
}

func solutionTarget(rel string) string {
	switch filepath.Base(rel) {
	case "solution.go":
		return filepath.Join(filepath.Dir(rel), "exercise.go")
	case "solution_test.go":
		return filepath.Join(filepath.Dir(rel), "exercise_test.go")
	default:
		return rel
	}
}

func solutionDiff(root string, item exercise, out io.Writer) (bool, error) {
	dir, err := os.MkdirTemp("", "go-koans-diff-")
	if err != nil {
		return false, err
	}
	defer os.RemoveAll(dir)
	for _, side := range []string{"your", "solution"} {
		if err := os.Mkdir(filepath.Join(dir, side), 0755); err != nil {
			return false, err
		}
	}
	solutionRoot := filepath.Join(root, "solutions", item.Path)
	files := 0
	err = filepath.WalkDir(solutionRoot, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}
		rel, err := filepath.Rel(solutionRoot, path)
		if err != nil {
			return err
		}
		rel = solutionTarget(rel)
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		data = bytes.TrimPrefix(data, []byte("//go:build ignore\n\n"))
		if err := writeDiffFile(filepath.Join(dir, "solution", rel), data); err != nil {
			return err
		}
		files++
		return copyDiffFile(filepath.Join(root, item.Path, rel), filepath.Join(dir, "your", rel))
	})
	if err != nil {
		return false, err
	}
	if files == 0 {
		return false, fmt.Errorf("%s: 模範解答のGoファイルがありません", item.Path)
	}
	for _, name := range item.Remove {
		source := filepath.Join(root, item.Path, name)
		err := filepath.WalkDir(source, func(path string, entry fs.DirEntry, err error) error {
			if errors.Is(err, os.ErrNotExist) {
				return nil
			}
			if err != nil {
				return err
			}
			if entry.IsDir() {
				return nil
			}
			rel, err := filepath.Rel(filepath.Join(root, item.Path), path)
			if err != nil {
				return err
			}
			return copyDiffFile(path, filepath.Join(dir, "your", rel))
		})
		if err != nil {
			return false, err
		}
	}
	cmd := exec.Command("git", "--no-pager", "diff", "--no-index", "--no-ext-diff", "--no-textconv", "--color=never", "--src-prefix=", "--dst-prefix=", "--", "your", "solution")
	cmd.Dir, cmd.Stdout = dir, out
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		var exit *exec.ExitError
		if errors.As(err, &exit) && exit.ExitCode() == 1 {
			return true, nil
		}
		return false, fmt.Errorf("git diff: %w\n%s", err, stderr.String())
	}
	return false, nil
}

func copyDiffFile(source, target string) error {
	data, err := os.ReadFile(source)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		return err
	}
	return writeDiffFile(target, data)
}

func writeDiffFile(path string, data []byte) error {
	if strings.HasSuffix(path, ".go") {
		if formatted, err := format.Source(data); err == nil {
			data = formatted
		}
	}
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
