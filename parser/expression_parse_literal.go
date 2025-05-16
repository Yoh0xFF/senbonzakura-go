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
	case lexer.TokenBoolean:
		return parseBooleanLiteralExpression(parser)
	case lexer.TokenNil:
		return parseNilLiteralExpression(parser)
	case lexer.TokenNumber:
		return parseNumericLiteralExpression(parser)
	case lexer.TokenString:
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
	token := eatToken(parser, lexer.TokenBoolean)
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
	eatToken(parser, lexer.TokenNil)

	return &ast.NilLiteralExpression{}
}

// parseNumericLiteralExpression parses numeric literals
//
// NumericLiteral
//
//	: NUMBER
//	;
func parseNumericLiteralExpression(parser *Parser) ast.Expression {
	token := eatToken(parser, lexer.TokenNumber)
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
	token := eatToken(parser, lexer.TokenString)
	// Remove the surrounding quotes
	tokenValue := parser.source[token.Start+1 : token.End-1]

	return &ast.StringLiteralExpression{
		Value: tokenValue,
	}
}
