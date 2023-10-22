package test

import (
	"os"
	"sbmf/internal"
	"testing"
)

func TestMainFunctions(t *testing.T) {
	file := "sample.yml"
	fileExpected := "sample_expected.yml"
	originalContent := mustReadFile(t, file)
	defer func() {
		err := os.WriteFile(file, originalContent, 0655)
		if err != nil {
			t.Fatal(err)
		}
	}()

	internal.IncreaseVersion(file)
	internal.Generate(file)

	rewrittenContent := mustReadFile(t, file)
	expectedContent := mustReadFile(t, fileExpected)

	if string(rewrittenContent) != string(expectedContent) {
		t.Fatalf("written file does not match expected file\nexpected:\n%s\nwritten:\n%s", string(expectedContent), string(rewrittenContent))
	}
}

func mustReadFile(t *testing.T, file string) []byte {
	t.Helper()
	content, err := os.ReadFile(file)
	if err != nil {
		t.Fatalf("could not read file %s: %v", file, err)
	}

	return content
}
