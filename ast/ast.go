package ast

type Node interface {
	NodeType() NodeType
}

// Statement represents different statement types in the AST
type Statement interface {
	Node
	StatementDispatcher
	isStatement()
}

// Expression represents different expression types in the AST
type Expression interface {
	Node
	ExpressionDispatcher
	isExpression()
}

// Parameter represents a function parameter with name and type
type Parameter struct {
	Name Expression
	Type Type
}
