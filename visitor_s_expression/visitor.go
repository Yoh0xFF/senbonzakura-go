package visitor_s_expression

import "github.com/yoh0xff/senbonzakura/ast"

type SExpressionVisitor struct{}

func (v *SExpressionVisitor) VisitStatement(statement ast.Statement) any {
	return visitStatement(v, statement)
}

func (v *SExpressionVisitor) VisitExpression(expression ast.Expression) any {
	return visitExpression(v, expression)
}
