package ast

type Visitor interface {
	// VisitStatement process statement node
	VisitStatement(statement Statement)

	// VisitExpression process expression node
	VisitExpression(expression Expression)
}

type StatementDispatcher interface {
	Accept(visitor Visitor)
}

type ExpressionDispatcher interface {
	Accept(visitor Visitor)
}
