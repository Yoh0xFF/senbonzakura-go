package ast

type Visitor interface {
	// VisitStatement process statement node
	VisitStatement(statement Statement) any

	// VisitExpression process expression node
	VisitExpression(expression Expression) any
}

type StatementDispatcher interface {
	Accept(visitor Visitor) any
}

type ExpressionDispatcher interface {
	Accept(visitor Visitor) any
}
