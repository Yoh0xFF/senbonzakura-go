package parser

import (
	"strconv"
	"strings"

	"github.com/yoh0xff/senbonzakura/ast"
	"github.com/yoh0xff/senbonzakura/lexer"
)

// parseLiteralExpression parses different types of literals
//
// Literal
//
//	: BooleanLiteral
//	: NilLiteral
//	: NumericLiteral
//	| StringLiteral
//	;
func parseLiteralExpression(parser *Parser) ast.Expression {
	switch parser.lookahead.TokenType {
	case lexer.Boolean:
		return parseBooleanLiteralExpression(parser)
	case lexer.Nil:
		return parseNilLiteralExpression(parser)
	case lexer.Number:
		return parseNumericLiteralExpression(parser)
	case lexer.String:
		return parseStringLiteralExpression(parser)
	default:
		panic("Literal: unexpected literal production")
	}
}

// parseBooleanLiteralExpression parses boolean literals
//
// BooleanLiteral
//
//	: BOOLEAN
//	;
func parseBooleanLiteralExpression(parser *Parser) ast.Expression {
	token := eat(parser, lexer.Boolean)
	tokenValue := parser.source[token.Start:token.End]
	boolValue := tokenValue == "true"

	return &ast.BooleanLiteralExpression{
		Value: boolValue,
	}
}

// parseNilLiteralExpression parses nil literals
//
// NilLiteral
//
//	: NIL
//	;
func parseNilLiteralExpression(parser *Parser) ast.Expression {
	eat(parser, lexer.Nil)

	return &ast.NilLiteralExpression{}
}

// parseNumericLiteralExpression parses numeric literals
//
// NumericLiteral
//
//	: NUMBER
//	;
func parseNumericLiteralExpression(parser *Parser) ast.Expression {
	token := eat(parser, lexer.Number)
	tokenValue := parser.source[token.Start:token.End]
	tokenValue = strings.TrimSpace(tokenValue)

	// Parse the number as int32
	numValue, err := strconv.ParseInt(tokenValue, 10, 32)
	if err != nil {
		panic("Invalid numeric literal: " + err.Error())
	}

	return &ast.NumericLiteralExpression{
		Value: int32(numValue),
	}
}

// parseStringLiteralExpression parses string literals
//
// StringLiteral
//
//	: STRING
//	;
func parseStringLiteralExpression(parser *Parser) ast.Expression {
	token := eat(parser, lexer.String)
	// Remove the surrounding quotes
	tokenValue := parser.source[token.Start+1 : token.End-1]

	return &ast.StringLiteralExpression{
		Value: tokenValue,
	}
}
