package types

import (
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"strings"
	"testing"
)

func TestTypesPackagesDoNotImportInfrastructurePackages(t *testing.T) {
	files, err := goFilesUnder(".")
	if err != nil {
		t.Fatalf("collect Go files: %v", err)
	}

	for _, file := range files {
		parsed, err := parser.ParseFile(token.NewFileSet(), file, nil, parser.ImportsOnly)
		if err != nil {
			t.Fatalf("parse %s imports: %v", file, err)
		}

		for _, spec := range parsed.Imports {
			path := strings.Trim(spec.Path.Value, `"`)
			if strings.HasPrefix(path, "github.com/rei0721/go-scaffold/pkg/") {
				t.Fatalf("types packages must not import lower infrastructure package %q from %s", path, file)
			}
		}
	}
}

func goFilesUnder(root string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, ".go") && !strings.HasSuffix(path, "_test.go") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
