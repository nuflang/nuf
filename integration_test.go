package main

import (
	"os"
	"testing"

	"github.com/nuflang/nuf/codegen"
	"github.com/nuflang/nuf/lexer"
	"github.com/nuflang/nuf/parser"
)

// func TestNufToHTMLLandmarks(t *testing.T) {
// 	inputFilePathname := "./test-data/landmarks/landmarks_input.nuf"
// 	inputBytes, err := os.ReadFile(inputFilePathname)
// 	if err != nil {
// 		t.Fatalf("Failed to read input file: %s", inputFilePathname)
// 	}
//
// 	// outputFilePathname := "./test-data/landmarks/landmarks_output.html"
// 	// outputBytes, err := os.ReadFile(outputFilePathname)
// 	// if err != nil {
// 	// 	t.Fatalf("Failed to read output file: %s", outputFilePathname)
// 	// }
//
// 	tokens := lexer.Tokenize(string(inputBytes))
// 	ast := parser.Parse(tokens)
// 	litter.Dump(ast)
//
// 	// html := codegen.GenerateHTML(ast)
//
// 	// if html != string(outputBytes) {
// 	// 	t.Errorf("Expected: %s, Got: %s", string(outputBytes), html)
// 	// }
// }

func TestNufToHTMLLandmarksWithHeadings(t *testing.T) {
	inputFilePathname := "./test-data/landmarks/landmarks_with_headings_input.nuf"
	inputBytes, err := os.ReadFile(inputFilePathname)
	if err != nil {
		t.Fatalf("Failed to read input file: %s", inputFilePathname)
	}

	outputFilePathname := "./test-data/landmarks/landmarks_with_headings_output.html"
	outputBytes, err := os.ReadFile(outputFilePathname)
	if err != nil {
		t.Fatalf("Failed to read output file: %s", outputFilePathname)
	}

	tokens := lexer.Tokenize(string(inputBytes))
	ast := parser.Parse(tokens)
	html := codegen.GenerateHTML(ast)

	if html != string(outputBytes) {
		t.Errorf("Expected: %s, Got: %s", string(outputBytes), html)
	}
}
