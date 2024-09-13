package main

import (
	"log"
	"os"

	"github.com/nuflang/nuf/lexer"
)

func main() {
	landmarkInputFile := "./test-data/landmarks/input.nuf"
	landmarkBytes, err := os.ReadFile(landmarkInputFile)
	if err != nil {
		log.Fatalf("Failed to read file: %s", landmarkInputFile)
	}

	tokens := lexer.Tokenize(string(landmarkBytes))
	for _, token := range tokens {
		token.Debug()
	}
}
