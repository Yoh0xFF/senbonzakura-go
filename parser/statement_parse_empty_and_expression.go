package parser

import (
	"github.com/yoh0xff/senbonzakura/ast"
	"github.com/yoh0xff/senbonzakura/lexer"
)

// parseEmptyStatement parses empty statements
//
// EmptyStatement
//
//	: ';'
//	;
func parseEmptyStatement(parser *Parser) ast.Statement {
	eatToken(parser, lexer.TokenStatementEnd)
	return &ast.EmptyStatement{}
}

// parseExpressionStatement parses expression statements
//
// ExpressionStatement
//
//	: Expression ';'
//	;
func parseExpressionStatement(parser *Parser, consumeStatementEnd bool) ast.Statement {
	expression := ParseRootExpression(parser)

	if consumeStatementEnd {
		eatToken(parser, lexer.TokenStatementEnd)
	}

	return &ast.ExpressionStatement{
		Expression: expression,
	}
}
