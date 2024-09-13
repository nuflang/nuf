package main

import (
	"log"
	"os"

	"github.com/nuflang/nuf/codegen"
	"github.com/nuflang/nuf/lexer"
	"github.com/nuflang/nuf/parser"
)

func main() {
	landmarkInputFile := "./test-data/landmarks/input.nuf"
	landmarkBytes, err := os.ReadFile(landmarkInputFile)
	if err != nil {
		log.Fatalf("Failed to read file: %s", landmarkInputFile)
	}

	tokens := lexer.Tokenize(string(landmarkBytes))
	// for _, token := range tokens {
	// 	token.Debug()
	// }

	ast := parser.Parse(tokens)
	// litter.Dump(ast)

	html := codegen.GenerateHTML(ast)
	landmarkGeneratedFile := "./test-data/landmarks/generated.html"
	err = os.WriteFile(landmarkGeneratedFile, []byte(html), 0400)
	if err != nil {
		log.Fatalf("Failed to write file: %s", landmarkGeneratedFile)
	}
}
