package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// tokenRecord is a stable, JSON-friendly view of the native compiler lexer.
// These records refer to the specified input file, not expanded imports.
type tokenRecord struct {
	Kind   string `json:"kind"`
	Text   string `json:"text"`
	Line   int    `json:"line"`
	Column int    `json:"column"`
}

func writeTokens(filename string, jsonFormat bool, dst io.Writer) error {
	b, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("read source: %w", err)
	}
	tokens, err := lex(string(b))
	if err != nil {
		return err
	}
	if jsonFormat {
		out := make([]tokenRecord, 0, len(tokens))
		for _, tok := range tokens {
			out = append(out, tokenRecord{Kind: tok.kind, Text: tok.text, Line: tok.line, Column: tok.col})
		}
		enc := json.NewEncoder(dst)
		enc.SetIndent("", "  ")
		return enc.Encode(out)
	}
	for _, tok := range tokens {
		if _, err := fmt.Fprintf(dst, "%d:%d\t%s\t%q\n", tok.line, tok.col, tok.kind, tok.text); err != nil {
			return err
		}
	}
	return nil
}

var projectNamePattern = regexp.MustCompile(`^[A-Za-z][A-Za-z0-9_-]{0,63}$`)

// newProject creates a minimal, compileable DVILR program in a *new* directory.
// It deliberately refuses to touch an existing directory, preventing clobbering.
func newProject(target string) error {
	name := filepath.Base(filepath.Clean(target))
	if !projectNamePattern.MatchString(name) {
		return fmt.Errorf("project directory name %q must start with a letter and contain only letters, digits, '-' or '_' (max 64 characters)", name)
	}
	if err := os.Mkdir(target, 0755); err != nil {
		return fmt.Errorf("create new project directory: %w", err)
	}
	const source = "# Your first DVILR program\nlet answer = 40 + 2\nsay \"Hello from DVILR!\"\nsay answer\n"
	const instructions = "# New DVILR project\n\nThis project targets the experimental Windows x64 native compiler.\n\nCompile and run on Windows:\n\n```powershell\ndvl.exe check main.dvl\ndvl.exe build main.dvl -o app.exe\n.\\app.exe\n```\n\nThe compiler is currently Go-bootstrapped, not self-hosted.\n"
	if err := os.WriteFile(filepath.Join(target, "main.dvl"), []byte(source), 0644); err != nil {
		return fmt.Errorf("write starter source: %w", err)
	}
	if err := os.WriteFile(filepath.Join(target, "README.md"), []byte(instructions), 0644); err != nil {
		return fmt.Errorf("write starter README: %w", err)
	}
	return nil
}

// Parse a fixed command syntax; allow the JSON switch before or after the file.
func parseTokenArgs(args []string) (string, bool, error) {
	filename := ""
	jsonMode := false
	for _, arg := range args {
		if arg == "--json" {
			if jsonMode {
				return "", false, fmt.Errorf("duplicate --json option")
			}
			jsonMode = true
		} else if strings.HasPrefix(arg, "-") || filename != "" {
			return "", false, fmt.Errorf("unexpected tokens argument %q", arg)
		} else {
			filename = arg
		}
	}
	if filename == "" {
		return "", false, fmt.Errorf("tokens requires a .dvl file")
	}
	return filename, jsonMode, nil
}
