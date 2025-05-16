package parser

import (
	"fmt"
	"github.com/yoh0xff/senbonzakura/ast"
	"github.com/yoh0xff/senbonzakura/lexer"
)

// parseAdditiveExpression parses additive expressions
//
// AdditiveExpression
//
//	: FactorExpression
//	| AdditiveExpression ADDITIVE_OPERATOR FactorExpression
//	;
func parseAdditiveExpression(parser *Parser) ast.Expression {
	return parseBinaryExpression(
		parser,
		lexer.TokenAdditiveOperator,
		parseFactorExpression,
		func(op string) ast.BinaryOperator {
			switch op {
			case "+":
				return ast.OperatorAdd
			case "-":
				return ast.OperatorSubtract
			default:
				panic(fmt.Sprintf("Unknown additive operator %s", op))
			}
		},
	)
}

// parseFactorExpression parses factor expressions
//
// FactorExpression
//
//	: PrimaryExpression
//	| FactorExpression FACTOR_OPERATOR PrimaryExpression
//	;
func parseFactorExpression(parser *Parser) ast.Expression {
	return parseBinaryExpression(
		parser,
		lexer.TokenFactorOperator,
		parseUnaryExpression,
		func(op string) ast.BinaryOperator {
			switch op {
			case "*":
				return ast.OperatorMultiply
			case "/":
				return ast.OperatorDivide
			default:
				panic(fmt.Sprintf("Unknown factor operator %s", op))
			}
		},
	)
}
