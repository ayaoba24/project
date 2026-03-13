package main

import (
	"os"
	"strings"
	"testing"
)

func TestGenerateASCII_HelloStandard(t *testing.T) {
	result, err := GenerateASCII("hello", "standard")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(result, " _") {
		t.Error("expected ASCII art underscores, got:\n", result)
	}
}

func TestGenerateASCII_EmptyString(t *testing.T) {
	result, err := GenerateASCII("", "standard")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "\n" {
		t.Errorf("expected single newline, got %q", result)
	}
}

func TestGenerateASCII_UnknownBanner(t *testing.T) {
	_, err := GenerateASCII("hello", "unknown")
	if err == nil {
		t.Error("expected error for unknown banner, got nil")
	}
}

func TestGenerateASCII_UnsupportedChar(t *testing.T) {
	_, err := GenerateASCII(string([]byte{0x01}), "standard")
	if err == nil {
		t.Error("expected error for non-printable character")
	}
}
func TestOutputFile_Created(t *testing.T) {
	tmpFile := t.TempDir() + "/out.txt"

	result, err := GenerateASCII("Hi", "standard")
	if err != nil {
		t.Fatal(err)
	}

	err = os.WriteFile(tmpFile, []byte(result), 0o644)
	if err != nil {
		t.Fatal(err)
	}

	data, err := os.ReadFile(tmpFile)
	if err != nil {
		t.Fatal(err)
	}

	if string(data) != result {
		t.Errorf("file content mismatch\ngot:  %q\nwant: %q", string(data), result)
	}
}

func TestOutputFile_MatchesStdout(t *testing.T) {
	result1, _ := GenerateASCII("hello", "standard")
	result2, _ := GenerateASCII("hello", "standard")
	if result1 != result2 {
		t.Error("GenerateASCII is not deterministic")
	}
}
