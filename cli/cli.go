package cli

import (
	"flag"
	"log"
	"os"

	"github.com/nuflang/nuf/evaluator"
	"github.com/nuflang/nuf/lexer"
	"github.com/nuflang/nuf/parser"
)

func RunCLI() {
	generateCmd := flag.NewFlagSet("generate", flag.ExitOnError)

	if len(os.Args) < 2 {
		os.Exit(1)
	}

	switch os.Args[1] {
	case "generate":
		generateCmd.Parse(os.Args[2:])
		inputFilename := generateCmd.Args()[0]
		outputFilename := generateCmd.Args()[1]

		inputBytes, err := os.ReadFile(inputFilename)
		if err != nil {
			log.Fatalf("Failed to read file: %s", inputFilename)
		}

		lex := lexer.NewLexer(string(inputBytes))
		p := parser.NewParser(lex)
		program := p.ParseProgram()
		output := evaluator.NewOutput()
		output.Eval(program)

		err = os.WriteFile(outputFilename, []byte(output.Value), 0600)
		if err != nil {
			log.Fatalf("Failed to write file %s: %s", outputFilename, err)
		}
	}
}
