package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNewProjectAndCompiler(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "HelloWorld")
	if err := newProject(dir); err != nil {
		t.Fatal(err)
	}
	content, err := os.ReadFile(filepath.Join(dir, "main.dvl"))
	if err != nil {
		t.Fatal(err)
	}
	g, p, err := compile(string(content))
	if err != nil {
		t.Fatalf("starter must compile: %v", err)
	}
	if _, err := makePE(g, p); err != nil {
		t.Fatalf("starter must generate a native executable: %v", err)
	}
	if err := newProject(dir); err == nil {
		t.Fatal("must not overwrite an existing project")
	}
	if _, err := os.Stat(filepath.Join(dir, "README.md")); err != nil {
		t.Fatal(err)
	}
}

func TestNewProjectRejectsInvalidNames(t *testing.T) {
	for _, name := range []string{"12bad", ".", "bad.name", "has space", "_private"} {
		if err := newProject(filepath.Join(t.TempDir(), name)); err == nil {
			t.Errorf("expected rejection for %q", name)
		}
	}
}

func TestTokenDumpJSONAndLocations(t *testing.T) {
	filename := filepath.Join(t.TempDir(), "a.dvl")
	if err := os.WriteFile(filename, []byte("let x=42\nsay \"hello\"\n"), 0644); err != nil {
		t.Fatal(err)
	}
	var buf bytes.Buffer
	if err := writeTokens(filename, true, &buf); err != nil {
		t.Fatal(err)
	}
	var tokens []tokenRecord
	if err := json.Unmarshal(buf.Bytes(), &tokens); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if len(tokens) != 7 || tokens[0].Text != "let" || tokens[0].Line != 1 || tokens[0].Column != 1 {
		t.Fatalf("unexpected token sequence: %+v", tokens)
	}
	if tokens[4].Text != "say" || tokens[4].Line != 2 || tokens[4].Column != 1 {
		t.Fatalf("incorrect line/column: %+v", tokens[4])
	}
	buf.Reset()
	if err := writeTokens(filename, false, &buf); err != nil || !strings.Contains(buf.String(), "2:1\tident\t\"say\"") {
		t.Fatalf("human-readable token dump: %s, %v", buf.String(), err)
	}
}

func TestTokensRejectMalformedAndArgs(t *testing.T) {
	filename := filepath.Join(t.TempDir(), "invalid.dvl")
	if err := os.WriteFile(filename, []byte(`say "unterminated`), 0644); err != nil {
		t.Fatal(err)
	}
	if err := writeTokens(filename, true, &bytes.Buffer{}); err == nil {
		t.Fatal("invalid source accepted")
	}
	for _, args := range [][]string{{}, {"--json"}, {"a.dvl", "b.dvl"}, {"a.dvl", "--json", "--json"}, {"a.dvl", "--bad"}} {
		if _, _, err := parseTokenArgs(args); err == nil {
			t.Errorf("should reject %#v", args)
		}
	}
	if name, isJSON, err := parseTokenArgs([]string{"--json", "a.dvl"}); err != nil || name != "a.dvl" || !isJSON {
		t.Fatalf("valid alternate order rejected: %v", err)
	}
}
