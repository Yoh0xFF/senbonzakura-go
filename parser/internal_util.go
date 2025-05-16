package parser

import (
	"fmt"
	"github.com/yoh0xff/senbonzakura/ast"
	"github.com/yoh0xff/senbonzakura/lexer"
)

// Eat expects a token of a given type
func eat(parser *Parser, tokenType lexer.TokenType) lexer.Token {
	if parser.lookahead.TokenType != tokenType {
		panic(fmt.Sprintf(
			"Unexpected token: %s, expected token: '%s'",
			parser.lookahead.TokenType.String(), tokenType.String(),
		))
	}

	preToken := parser.lookahead
	parser.lookahead = parser.lexer.NextToken()
	return preToken
}

// EatAnyOf expects a token of any given types
func eatAnyOf(parser *Parser, tokenTypes []lexer.TokenType) lexer.Token {
	for _, tokenType := range tokenTypes {
		if parser.lookahead.TokenType == tokenType {
			preToken := parser.lookahead
			parser.lookahead = parser.lexer.NextToken()
			return preToken
		}
	}

	panic(fmt.Sprintf(
		"Unexpected token: %s, expected tokens: '%v'",
		parser.lookahead.TokenType.String(), tokenTypes,
	))
}

// IsToken checks the current token type
func isToken(parser *Parser, tokenType lexer.TokenType) bool {
	return parser.lookahead.TokenType == tokenType
}

// IsAnyOfToken checks if the current token is any of the given types
func isAnyOfToken(parser *Parser, tokenTypes []lexer.TokenType) bool {
	for _, tokenType := range tokenTypes {
		if parser.lookahead.TokenType == tokenType {
			return true
		}
	}

	return false
}

// IsValidAssignmentTarget checks if the expression is valid assignment target
func isValidAssignmentTarget(expression ast.Expression) bool {
	switch expression.NodeType() {
	case ast.NodeIdentifierExpression, ast.NodeMemberExpression:
		return true
	default:
		return false
	}
}

// IsLiteralToken checks if the current token is literal
func isLiteralToken(parser *Parser) bool {
	return isAnyOfToken(
		parser,
		[]lexer.TokenType{
			lexer.Boolean,
			lexer.Nil,
			lexer.Number,
			lexer.String,
		},
	)
}

// IsAssignmentOperatorToken checks if the current token is assignment operator
func isAssignmentOperatorToken(parser *Parser) bool {
	return isAnyOfToken(
		parser,
		[]lexer.TokenType{
			lexer.SimpleAssignmentOperator,
			lexer.ComplexAssignmentOperator,
		},
	)
}

// OperandParserFunc defines the function signature for parsing operands
type OperandParserFunc func(*Parser) ast.Expression

// BinaryOperatorMapperFunc defines the function signature for mapping binary operators
type BinaryOperatorMapperFunc func(string) ast.BinaryOperator

// LogicalOperatorMapperFunc defines the function signature for mapping logical operators
type LogicalOperatorMapperFunc func(string) ast.LogicalOperator

// ParseBinaryExpression parses generic binary expressions
func parseBinaryExpression(
	parser *Parser,
	tokenType lexer.TokenType,
	operandParser OperandParserFunc,
	operatorMapper BinaryOperatorMapperFunc,
) ast.Expression {
	left := operandParser(parser)

	for parser.lookahead.TokenType == tokenType {
		operatorToken := eat(parser, tokenType)
		operatorValue := parser.source[operatorToken.Start:operatorToken.End]
		operator := operatorMapper(operatorValue)

		right := operandParser(parser)

		left = &ast.BinaryExpression{
			Operator: operator,
			Left:     left,
			Right:    right,
		}
	}

	return left
}

// ParseLogicalExpression parses generic logical expressions
func parseLogicalExpression(
	parser *Parser,
	tokenType lexer.TokenType,
	operandParser OperandParserFunc,
	operatorMapper LogicalOperatorMapperFunc,
) ast.Expression {
	left := operandParser(parser)

	for parser.lookahead.TokenType == tokenType {
		operatorToken := eat(parser, tokenType)
		operatorValue := parser.source[operatorToken.Start:operatorToken.End]
		operator := operatorMapper(operatorValue)

		right := operandParser(parser)

		left = &ast.LogicalExpression{
			Operator: operator,
			Left:     left,
			Right:    right,
		}
	}

	return left
}
