package visitor_s_expression

import (
	"fmt"
	"github.com/yoh0xff/senbonzakura/ast"
	"strings"
)

func visitExpression(visitor *SExpressionVisitor, expression ast.Expression) any {
	switch expression.NodeType() {
	case ast.NodeVariableExpression:
		return visitVariableExpression(visitor, expression.(*ast.VariableExpression))
	case ast.NodeAssignmentExpression:
		return visitAssignmentExpression(visitor, expression.(*ast.AssignmentExpression))
	case ast.NodeBinaryExpression:
		return visitBinaryExpression(visitor, expression.(*ast.BinaryExpression))
	case ast.NodeUnaryExpression:
		return visitUnaryExpression(visitor, expression.(*ast.UnaryExpression))
	case ast.NodeLogicalExpression:
		return visitLogicalExpression(visitor, expression.(*ast.LogicalExpression))
	case ast.NodeBooleanLiteralExpression:
		return visitBooleanLiteralExpression(visitor, expression.(*ast.BooleanLiteralExpression))
	case ast.NodeNilLiteralExpression:
		return visitNilLiteralExpression(visitor)
	case ast.NodeNumericLiteralExpression:
		return visitNumericLiteralExpression(visitor, expression.(*ast.NumericLiteralExpression))
	case ast.NodeStringLiteralExpression:
		return visitStringLiteralExpression(visitor, expression.(*ast.StringLiteralExpression))
	case ast.NodeIdentifierExpression:
		return visitIdentifierExpression(visitor, expression.(*ast.IdentifierExpression))
	case ast.NodeMemberExpression:
		return visitMemberExpression(visitor, expression.(*ast.MemberExpression))
	case ast.NodeCallExpression:
		return visitCallExpression(visitor, expression.(*ast.CallExpression))
	case ast.NodeThisExpression:
		return visitThisExpression(visitor)
	case ast.NodeSuperExpression:
		return visitSuperExpression(visitor)
	case ast.NodeNewExpression:
		return visitNewExpression(visitor, expression.(*ast.NewExpression))
	default:
		return fmt.Errorf("unknown expression type: %T", expression)
	}
}

func visitVariableExpression(visitor *SExpressionVisitor, expression *ast.VariableExpression) error {
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
	return nil
}

func visitAssignmentExpression(visitor *SExpressionVisitor, expression *ast.AssignmentExpression) error {
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
	return nil
}

func visitBinaryExpression(visitor *SExpressionVisitor, expression *ast.BinaryExpression) error {
	visitor.beginExpression("binary")

	visitor.writeSpaceOrNewLine()
	visitor.writeString(fmt.Sprintf("\"%s\"", expression.Operator.String()))

	visitor.writeSpaceOrNewLine()
	expression.Left.Accept(visitor)

	visitor.writeSpaceOrNewLine()
	expression.Right.Accept(visitor)

	visitor.endExpression()
	return nil
}

func visitUnaryExpression(visitor *SExpressionVisitor, expression *ast.UnaryExpression) error {
	visitor.beginExpression("unary")

	visitor.writeSpaceOrNewLine()
	visitor.writeString(fmt.Sprintf("\"%s\"", expression.Operator.String()))

	visitor.writeSpaceOrNewLine()
	expression.Right.Accept(visitor)

	visitor.endExpression()
	return nil
}

func visitLogicalExpression(visitor *SExpressionVisitor, expression *ast.LogicalExpression) error {
	visitor.beginExpression("logical")

	visitor.writeSpaceOrNewLine()
	visitor.writeString(fmt.Sprintf("\"%s\"", expression.Operator.String()))

	visitor.writeSpaceOrNewLine()
	expression.Left.Accept(visitor)

	visitor.writeSpaceOrNewLine()
	expression.Right.Accept(visitor)

	visitor.endExpression()
	return nil
}

func visitBooleanLiteralExpression(visitor *SExpressionVisitor, expression *ast.BooleanLiteralExpression) error {
	visitor.beginExpression("boolean")
	visitor.writeString(fmt.Sprintf(" %t", expression.Value))
	visitor.endExpression()
	return nil
}

func visitNilLiteralExpression(visitor *SExpressionVisitor) error {
	visitor.beginExpression("nil")
	visitor.endExpression()
	return nil
}

func visitNumericLiteralExpression(visitor *SExpressionVisitor, expression *ast.NumericLiteralExpression) error {
	visitor.beginExpression("number")
	visitor.writeString(fmt.Sprintf(" %d", expression.Value))
	visitor.endExpression()
	return nil
}

func visitStringLiteralExpression(visitor *SExpressionVisitor, expression *ast.StringLiteralExpression) error {
	visitor.beginExpression("string")

	// Replace double quotes with escaped quotes
	escaped := strings.Replace(expression.Value, "\"", "\\\"", -1)
	visitor.writeString(fmt.Sprintf(" \"%s\"", escaped))

	visitor.endExpression()
	return nil
}

func visitIdentifierExpression(visitor *SExpressionVisitor, expression *ast.IdentifierExpression) error {
	visitor.beginExpression("id")

	// Write the identifier name
	visitor.writeString(fmt.Sprintf(" %s", expression.Name))

	visitor.endExpression()
	return nil
}

func visitMemberExpression(visitor *SExpressionVisitor, expression *ast.MemberExpression) error {
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
	return nil
}

func visitCallExpression(visitor *SExpressionVisitor, expression *ast.CallExpression) error {
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
	return nil
}

func visitThisExpression(visitor *SExpressionVisitor) error {
	visitor.beginExpression("this")
	visitor.endExpression()
	return nil
}

func visitSuperExpression(visitor *SExpressionVisitor) error {
	visitor.beginExpression("super")
	visitor.endExpression()
	return nil
}

func visitNewExpression(visitor *SExpressionVisitor, expression *ast.NewExpression) error {
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
	return nil
}
