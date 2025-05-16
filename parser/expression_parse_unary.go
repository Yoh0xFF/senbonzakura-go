package parser

import (
	"fmt"
	"github.com/yoh0xff/senbonzakura/ast"
	"github.com/yoh0xff/senbonzakura/lexer"
)

// parseUnaryExpression parses unary expressions
//
// UnaryExpression
//
//	: LeftHandSideExpression
//	| ADDITIVE_OPERATOR UnaryExpression
//	| LOGICAL_NOT_OPERATOR UnaryExpression
//	;
func parseUnaryExpression(parser *Parser) ast.Expression {
	// Check if the next token is an operator
	if isNextTokenAnyOfType(
		parser,
		[]lexer.TokenType{
			lexer.TokenAdditiveOperator,
			lexer.TokenLogicalNotOperator,
		},
	) {
		// Eat the operator token
		operatorToken := eatAnyOfToken(
			parser,
			[]lexer.TokenType{
				lexer.TokenAdditiveOperator,
				lexer.TokenLogicalNotOperator,
			},
		)
		operatorValue := parser.source[operatorToken.Start:operatorToken.End]

		// Map the operator string to our operator type
		var operator ast.UnaryOperator
		switch operatorValue {
		case "+":
			operator = ast.OperatorPlus
		case "-":
			operator = ast.OperatorMinus
		case "!":
			operator = ast.OperatorNot
		default:
			panic(fmt.Sprintf("Unknown unary operator %s", operatorValue))
		}

		// Create a unary expression node
		return &ast.UnaryExpression{
			Operator: operator,
			Right:    parseUnaryExpression(parser), // Recursive call for right operand
		}
	}

	// If no operator, then it's a left-hand-side expression
	return parseLeftHandSideExpression(parser)
}
