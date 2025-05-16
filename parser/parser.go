package parser

import (
	"github.com/yoh0xff/senbonzakura/ast"
	"github.com/yoh0xff/senbonzakura/lexer"
)

type Parser struct {
	source    string
	lexer     *lexer.Lexer
	lookahead lexer.Token
}

func NewParser(source string) *Parser {
	lexerInstance := lexer.NewLexer(source)
	lookahead := lexerInstance.NextToken()

	return &Parser{
		source:    source,
		lexer:     lexerInstance,
		lookahead: lookahead,
	}
}

// ParseRootStatement entry point to parse statement
// Parses a string into an AST
func ParseRootStatement(parser *Parser) ast.Statement {
	// TODO
	// parse_program_statement(parser)
	return nil
}

// ParseRootExpression entry point to parse expression
//
// Expression
//
//	: AssignmentExpression
//	;
func ParseRootExpression(parser *Parser) ast.Expression {
	return parseAssignmentExpression(parser)
}
