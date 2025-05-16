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
	visitor.BeginExpr("init")

	// Process identifier
	visitor.WriteSpaceOrNewLine()
	expression.Identifier.Accept(visitor)

	// Add type annotation to S-expression
	visitor.WriteSpaceOrNewLine()
	visitor.BeginExpr("type")
	visitType(visitor, expression.TypeAnnotation)
	visitor.EndExpr()

	// Process initializer if present
	if expression.Initializer != nil {
		visitor.WriteSpaceOrNewLine()
		expression.Initializer.Accept(visitor)
	}

	visitor.EndExpr()
	return nil
}

func visitAssignmentExpression(visitor *SExpressionVisitor, expression *ast.AssignmentExpression) error {
	visitor.BeginExpr("assign")

	// Write the operator
	visitor.WriteSpaceOrNewLine()
	visitor.WriteString(fmt.Sprintf("\"%s\"", expression.Operator.String()))

	// Process left operand
	visitor.WriteSpaceOrNewLine()
	expression.Left.Accept(visitor)

	// Process right operand
	visitor.WriteSpaceOrNewLine()
	expression.Right.Accept(visitor)

	visitor.EndExpr()
	return nil
}

func visitBinaryExpression(visitor *SExpressionVisitor, expression *ast.BinaryExpression) error {
	visitor.BeginExpr("binary")

	visitor.WriteSpaceOrNewLine()
	visitor.WriteString(fmt.Sprintf("\"%s\"", expression.Operator.String()))

	visitor.WriteSpaceOrNewLine()
	expression.Left.Accept(visitor)

	visitor.WriteSpaceOrNewLine()
	expression.Right.Accept(visitor)

	visitor.EndExpr()
	return nil
}

func visitUnaryExpression(visitor *SExpressionVisitor, expression *ast.UnaryExpression) error {
	visitor.BeginExpr("unary")

	visitor.WriteSpaceOrNewLine()
	visitor.WriteString(fmt.Sprintf("\"%s\"", expression.Operator.String()))

	visitor.WriteSpaceOrNewLine()
	expression.Right.Accept(visitor)

	visitor.EndExpr()
	return nil
}

func visitLogicalExpression(visitor *SExpressionVisitor, expression *ast.LogicalExpression) error {
	visitor.BeginExpr("logical")

	visitor.WriteSpaceOrNewLine()
	visitor.WriteString(fmt.Sprintf("\"%s\"", expression.Operator.String()))

	visitor.WriteSpaceOrNewLine()
	expression.Left.Accept(visitor)

	visitor.WriteSpaceOrNewLine()
	expression.Right.Accept(visitor)

	visitor.EndExpr()
	return nil
}

func visitBooleanLiteralExpression(visitor *SExpressionVisitor, expression *ast.BooleanLiteralExpression) error {
	visitor.BeginExpr("boolean")
	visitor.WriteString(fmt.Sprintf(" %t", expression.Value))
	visitor.EndExpr()
	return nil
}

func visitNilLiteralExpression(visitor *SExpressionVisitor) error {
	visitor.BeginExpr("nil")
	visitor.EndExpr()
	return nil
}

func visitNumericLiteralExpression(visitor *SExpressionVisitor, expression *ast.NumericLiteralExpression) error {
	visitor.BeginExpr("number")
	visitor.WriteString(fmt.Sprintf(" %d", expression.Value))
	visitor.EndExpr()
	return nil
}

func visitStringLiteralExpression(visitor *SExpressionVisitor, expression *ast.StringLiteralExpression) error {
	visitor.BeginExpr("string")

	// Replace double quotes with escaped quotes
	escaped := strings.Replace(expression.Value, "\"", "\\\"", -1)
	visitor.WriteString(fmt.Sprintf(" \"%s\"", escaped))

	visitor.EndExpr()
	return nil
}

func visitIdentifierExpression(visitor *SExpressionVisitor, expression *ast.IdentifierExpression) error {
	visitor.BeginExpr("id")

	// Write the identifier name
	visitor.WriteString(fmt.Sprintf(" %s", expression.Name))

	visitor.EndExpr()
	return nil
}

func visitMemberExpression(visitor *SExpressionVisitor, expression *ast.MemberExpression) error {
	visitor.BeginExpr("member")

	// Indicate whether the member access is computed (bracket notation) or not (dot notation)
	visitor.WriteSpaceOrNewLine()
	if expression.Computed {
		visitor.WriteString("\"computed\"")
	} else {
		visitor.WriteString("\"static\"")
	}

	// Process object expression (the left part of the member access)
	visitor.WriteSpaceOrNewLine()
	expression.Object.Accept(visitor)

	// Process property expression (the right part of the member access)
	visitor.WriteSpaceOrNewLine()
	expression.Property.Accept(visitor)

	visitor.EndExpr()
	return nil
}

func visitCallExpression(visitor *SExpressionVisitor, expression *ast.CallExpression) error {
	visitor.BeginExpr("call")

	// Process callee expression
	visitor.WriteSpaceOrNewLine()
	expression.Callee.Accept(visitor)

	if len(expression.Arguments) > 0 {
		// Create args expression
		visitor.WriteSpaceOrNewLine()
		visitor.BeginExpr("args")

		// Process each argument
		for _, arg := range expression.Arguments {
			visitor.WriteSpaceOrNewLine()
			arg.Accept(visitor)
		}

		visitor.EndExpr() // Close args expression
	}

	visitor.EndExpr()
	return nil
}

func visitThisExpression(visitor *SExpressionVisitor) error {
	visitor.BeginExpr("this")
	visitor.EndExpr()
	return nil
}

func visitSuperExpression(visitor *SExpressionVisitor) error {
	visitor.BeginExpr("super")
	visitor.EndExpr()
	return nil
}

func visitNewExpression(visitor *SExpressionVisitor, expression *ast.NewExpression) error {
	visitor.BeginExpr("new")

	// Process callee expression
	visitor.WriteSpaceOrNewLine()
	expression.Callee.Accept(visitor)

	if len(expression.Arguments) > 0 {
		// Create args expression
		visitor.WriteSpaceOrNewLine()
		visitor.BeginExpr("args")

		// Process each argument
		for _, arg := range expression.Arguments {
			visitor.WriteSpaceOrNewLine()
			arg.Accept(visitor)
		}

		visitor.EndExpr() // Close args expression
	}

	visitor.EndExpr()
	return nil
}
