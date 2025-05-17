package visitor_s_expression

import (
	"fmt"
	"strings"

	"github.com/yoh0xff/senbonzakura/ast"
)

func visitExpression(visitor *SExpressionVisitor, expression ast.Expression) {
	switch expression.NodeType() {
	case ast.NodeVariableExpression:
		visitVariableExpression(visitor, expression.(*ast.VariableExpression))
	case ast.NodeAssignmentExpression:
		visitAssignmentExpression(visitor, expression.(*ast.AssignmentExpression))
	case ast.NodeBinaryExpression:
		visitBinaryExpression(visitor, expression.(*ast.BinaryExpression))
	case ast.NodeUnaryExpression:
		visitUnaryExpression(visitor, expression.(*ast.UnaryExpression))
	case ast.NodeLogicalExpression:
		visitLogicalExpression(visitor, expression.(*ast.LogicalExpression))
	case ast.NodeBooleanLiteralExpression:
		visitBooleanLiteralExpression(visitor, expression.(*ast.BooleanLiteralExpression))
	case ast.NodeNilLiteralExpression:
		visitNilLiteralExpression(visitor)
	case ast.NodeNumericLiteralExpression:
		visitNumericLiteralExpression(visitor, expression.(*ast.NumericLiteralExpression))
	case ast.NodeStringLiteralExpression:
		visitStringLiteralExpression(visitor, expression.(*ast.StringLiteralExpression))
	case ast.NodeIdentifierExpression:
		visitIdentifierExpression(visitor, expression.(*ast.IdentifierExpression))
	case ast.NodeMemberExpression:
		visitMemberExpression(visitor, expression.(*ast.MemberExpression))
	case ast.NodeCallExpression:
		visitCallExpression(visitor, expression.(*ast.CallExpression))
	case ast.NodeThisExpression:
		visitThisExpression(visitor)
	case ast.NodeSuperExpression:
		visitSuperExpression(visitor)
	case ast.NodeNewExpression:
		visitNewExpression(visitor, expression.(*ast.NewExpression))
	default:
		panic(fmt.Errorf("unknown expression type: %T", expression))
	}
}

func visitVariableExpression(visitor *SExpressionVisitor, expression *ast.VariableExpression) {
	visitor.beginExpression("init")

	// Process identifier
	visitor.writeSpaceOrNewLine()
	expression.Identifier.Accept(visitor)

	// Add type annotation to S-expression
	visitor.writeSpaceOrNewLine()
	visitor.beginExpression("type")
	visitType(visitor, expression.TypeAnnotation)
	visitor.endExpression()

	// Process initializer if present
	if expression.Initializer != nil {
		visitor.writeSpaceOrNewLine()
		expression.Initializer.Accept(visitor)
	}

	visitor.endExpression()
}

func visitAssignmentExpression(visitor *SExpressionVisitor, expression *ast.AssignmentExpression) {
	visitor.beginExpression("assign")

	// Write the operator
	visitor.writeSpaceOrNewLine()
	visitor.writeString(fmt.Sprintf("\"%s\"", expression.Operator.String()))

	// Process left operand
	visitor.writeSpaceOrNewLine()
	expression.Left.Accept(visitor)

	// Process right operand
	visitor.writeSpaceOrNewLine()
	expression.Right.Accept(visitor)

	visitor.endExpression()
}

func visitBinaryExpression(visitor *SExpressionVisitor, expression *ast.BinaryExpression) {
	visitor.beginExpression("binary")

	visitor.writeSpaceOrNewLine()
	visitor.writeString(fmt.Sprintf("\"%s\"", expression.Operator.String()))

	visitor.writeSpaceOrNewLine()
	expression.Left.Accept(visitor)

	visitor.writeSpaceOrNewLine()
	expression.Right.Accept(visitor)

	visitor.endExpression()
}

func visitUnaryExpression(visitor *SExpressionVisitor, expression *ast.UnaryExpression) {
	visitor.beginExpression("unary")

	visitor.writeSpaceOrNewLine()
	visitor.writeString(fmt.Sprintf("\"%s\"", expression.Operator.String()))

	visitor.writeSpaceOrNewLine()
	expression.Right.Accept(visitor)

	visitor.endExpression()
}

func visitLogicalExpression(visitor *SExpressionVisitor, expression *ast.LogicalExpression) {
	visitor.beginExpression("logical")

	visitor.writeSpaceOrNewLine()
	visitor.writeString(fmt.Sprintf("\"%s\"", expression.Operator.String()))

	visitor.writeSpaceOrNewLine()
	expression.Left.Accept(visitor)

	visitor.writeSpaceOrNewLine()
	expression.Right.Accept(visitor)

	visitor.endExpression()
}

func visitBooleanLiteralExpression(visitor *SExpressionVisitor, expression *ast.BooleanLiteralExpression) {
	visitor.beginExpression("boolean")
	visitor.writeString(fmt.Sprintf(" %t", expression.Value))
	visitor.endExpression()
}

func visitNilLiteralExpression(visitor *SExpressionVisitor) {
	visitor.beginExpression("nil")
	visitor.endExpression()
}

func visitNumericLiteralExpression(visitor *SExpressionVisitor, expression *ast.NumericLiteralExpression) {
	visitor.beginExpression("number")
	visitor.writeString(fmt.Sprintf(" %d", expression.Value))
	visitor.endExpression()
}

func visitStringLiteralExpression(visitor *SExpressionVisitor, expression *ast.StringLiteralExpression) {
	visitor.beginExpression("string")

	// Replace double quotes with escaped quotes
	escaped := strings.Replace(expression.Value, "\"", "\\\"", -1)
	visitor.writeString(fmt.Sprintf(" \"%s\"", escaped))

	visitor.endExpression()
}

func visitIdentifierExpression(visitor *SExpressionVisitor, expression *ast.IdentifierExpression) {
	visitor.beginExpression("id")

	// Write the identifier name
	visitor.writeString(fmt.Sprintf(" %s", expression.Name))

	visitor.endExpression()
}

func visitMemberExpression(visitor *SExpressionVisitor, expression *ast.MemberExpression) {
	visitor.beginExpression("member")

	// Indicate whether the member access is computed (bracket notation) or not (dot notation)
	visitor.writeSpaceOrNewLine()
	if expression.Computed {
		visitor.writeString("\"computed\"")
	} else {
		visitor.writeString("\"static\"")
	}

	// Process object expression (the left part of the member access)
	visitor.writeSpaceOrNewLine()
	expression.Object.Accept(visitor)

	// Process property expression (the right part of the member access)
	visitor.writeSpaceOrNewLine()
	expression.Property.Accept(visitor)

	visitor.endExpression()
}

func visitCallExpression(visitor *SExpressionVisitor, expression *ast.CallExpression) {
	visitor.beginExpression("call")

	// Process callee expression
	visitor.writeSpaceOrNewLine()
	expression.Callee.Accept(visitor)

	if len(expression.Arguments) > 0 {
		// Create args expression
		visitor.writeSpaceOrNewLine()
		visitor.beginExpression("args")

		// Process each argument
		for _, arg := range expression.Arguments {
			visitor.writeSpaceOrNewLine()
			arg.Accept(visitor)
		}

		visitor.endExpression() // Close args expression
	}

	visitor.endExpression()
}

func visitThisExpression(visitor *SExpressionVisitor) {
	visitor.beginExpression("this")
	visitor.endExpression()
}

func visitSuperExpression(visitor *SExpressionVisitor) {
	visitor.beginExpression("super")
	visitor.endExpression()
}

func visitNewExpression(visitor *SExpressionVisitor, expression *ast.NewExpression) {
	visitor.beginExpression("new")

	// Process callee expression
	visitor.writeSpaceOrNewLine()
	expression.Callee.Accept(visitor)

	if len(expression.Arguments) > 0 {
		// Create args expression
		visitor.writeSpaceOrNewLine()
		visitor.beginExpression("args")

		// Process each argument
		for _, arg := range expression.Arguments {
			visitor.writeSpaceOrNewLine()
			arg.Accept(visitor)
		}

		visitor.endExpression() // Close args expression
	}

	visitor.endExpression()
}
