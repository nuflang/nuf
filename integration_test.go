package main

import (
	"os"
	"testing"

	"github.com/nuflang/nuf/codegen"
	"github.com/nuflang/nuf/lexer"
	"github.com/nuflang/nuf/parser"
)

func TestNufToHTML(t *testing.T) {
	landmarkInputFile := "./test-data/landmarks/input.nuf"
	landmarkInputBytes, err := os.ReadFile(landmarkInputFile)
	if err != nil {
		t.Fatalf("Failed to read input file: %s", landmarkInputFile)
	}

	landmarkOutputFile := "./test-data/landmarks/output.html"
	landmarkOutputBytes, err := os.ReadFile(landmarkOutputFile)
	if err != nil {
		t.Fatalf("Failed to read output file: %s", landmarkOutputFile)
	}

	tokens := lexer.Tokenize(string(landmarkInputBytes))

	ast := parser.Parse(tokens)

	html := codegen.GenerateHTML(ast)

	if html != string(landmarkOutputBytes) {
		t.Errorf("Expected: %s, Got: %s", string(landmarkOutputBytes), html)
	}
}
