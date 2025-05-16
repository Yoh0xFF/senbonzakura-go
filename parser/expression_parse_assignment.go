package parser

import (
	"fmt"
	"github.com/yoh0xff/senbonzakura/ast"
	"github.com/yoh0xff/senbonzakura/lexer"
)

// parseAssignmentExpression parses an assignment expression
//
// AssignmentExpression
//
//	: LogicalOrExpression
//	| LeftHandSideExpression ASSIGNMENT_OPERATOR AssignmentExpression
//	;
func parseAssignmentExpression(parser *Parser) ast.Expression {
	left := parseLogicalOrExpression(parser)

	if !isNextTokenAssignmentOperator(parser) {
		return left
	}

	assignmentOperatorToken := eatAnyOfToken(
		parser,
		[]lexer.TokenType{
			lexer.TokenSimpleAssignmentOperator,
			lexer.TokenComplexAssignmentOperator,
		},
	)

	assignmentOperatorValue := parser.source[assignmentOperatorToken.Start:assignmentOperatorToken.End]

	var assignmentOperator ast.AssignmentOperator
	switch assignmentOperatorValue {
	case "=":
		assignmentOperator = ast.OperatorAssign
	case "+=":
		assignmentOperator = ast.OperatorAssignAdd
	case "-=":
		assignmentOperator = ast.OperatorAssignSubtract
	case "*=":
		assignmentOperator = ast.OperatorAssignMultiply
	case "/=":
		assignmentOperator = ast.OperatorAssignDivide
	default:
		panic(fmt.Sprintf("Unknown assignment operator %s", assignmentOperatorValue))
	}

	if !isNextTokenValidAssignmentTarget(left) {
		panic("Invalid left-hand side in the assignment expression")
	}

	right := parseAssignmentExpression(parser)

	return &ast.AssignmentExpression{
		Operator: assignmentOperator,
		Left:     left,
		Right:    right,
	}
}
