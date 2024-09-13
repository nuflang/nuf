package parser

import (
	"github.com/nuflang/nuf/ast"
	"github.com/nuflang/nuf/lexer"
)

type bindingPower int

const (
	defaultBP bindingPower = iota
)

type blockStatementHandler func(p *parser) ast.BlockStatement

type blockStatementLookup map[lexer.TokenKind]blockStatementHandler
type bindingPowerLookup map[lexer.TokenKind]bindingPower

var bindingPowerLookupTable = bindingPowerLookup{}
var blockStatementLookupTable = blockStatementLookup{}

func blockStatement(kind lexer.TokenKind, blockStatementFunc blockStatementHandler) {
	bindingPowerLookupTable[kind] = defaultBP
	blockStatementLookupTable[kind] = blockStatementFunc
}

func createTokenLookups() {
	blockStatement(lexer.GREATER, parseLandmarkBlockStatement)
}
