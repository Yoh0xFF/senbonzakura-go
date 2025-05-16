package parser

import (
	"fmt"
	"github.com/yoh0xff/senbonzakura/ast"
	"github.com/yoh0xff/senbonzakura/lexer"
)

// parseLogicalOrExpression parses logical OR expressions
//
// LogicalOrExpression
//
//	: LogicalAndExpression LOGICAL_OR_OPERATOR LogicalOrExpression
//	| LogicalAndExpression
//	;
func parseLogicalOrExpression(parser *Parser) ast.Expression {
	return parseLogicalExpression(
		parser,
		lexer.TokenLogicalOrOperator,
		parseLogicalAndExpression,
		func(op string) ast.LogicalOperator {
			switch op {
			case "||":
				return ast.OperatorOr
			default:
				panic(fmt.Sprintf("Unknown logical operator %s", op))
			}
		},
	)
}

// parseLogicalAndExpression parses logical AND expressions
//
// LogicalAndExpression
//
//	: EqualityExpression LOGICAL_AND_OPERATOR LogicalAndExpression
//	| EqualityExpression
//	;
func parseLogicalAndExpression(parser *Parser) ast.Expression {
	return parseLogicalExpression(
		parser,
		lexer.TokenLogicalAndOperator,
		parseEqualityExpression,
		func(op string) ast.LogicalOperator {
			switch op {
			case "&&":
				return ast.OperatorAnd
			default:
				panic(fmt.Sprintf("Unknown logical operator %s", op))
			}
		},
	)
}

// parseEqualityExpression parses equality expressions
//
// EqualityExpression
//
//	: RelationalExpression EQUALITY_OPERATOR EqualityExpression
//	| RelationalExpression
//	;
func parseEqualityExpression(parser *Parser) ast.Expression {
	return parseBinaryExpression(
		parser,
		lexer.TokenEqualityOperator,
		parseRelationalExpression,
		func(op string) ast.BinaryOperator {
			switch op {
			case "==":
				return ast.OperatorEqual
			case "!=":
				return ast.OperatorNotEqual
			default:
				panic(fmt.Sprintf("Unknown equality operator %s", op))
			}
		},
	)
}

// parseRelationalExpression parses relational expressions
//
// RelationalExpression
//
//	: AdditiveExpression
//	| AdditiveExpression RELATIONAL_OPERATOR AdditiveExpression
//	;
func parseRelationalExpression(parser *Parser) ast.Expression {
	return parseBinaryExpression(
		parser,
		lexer.TokenRelationalOperator,
		parseAdditiveExpression,
		func(op string) ast.BinaryOperator {
			switch op {
			case ">":
				return ast.OperatorGreaterThan
			case ">=":
				return ast.OperatorGreaterThanOrEqualTo
			case "<":
				return ast.OperatorLessThan
			case "<=":
				return ast.OperatorLessThanOrEqualTo
			default:
				panic(fmt.Sprintf("Unknown relational operator %s", op))
			}
		},
	)
}
