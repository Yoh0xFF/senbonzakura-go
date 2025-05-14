package ast

type VariableExpression struct {
	nodeType       NodeType
	Identifier     Expression
	TypeAnnotation Type
	Initializer    Expression // can be nil
}

type AssignmentExpression struct {
	nodeType NodeType
	Operator AssignmentOperator
	Left     Expression
	Right    Expression
}

type BinaryExpression struct {
	nodeType NodeType
	Operator BinaryOperator
	Left     Expression
	Right    Expression
}

type UnaryExpression struct {
	nodeType NodeType
	Operator UnaryOperator
	Right    Expression
}

type LogicalExpression struct {
	nodeType NodeType
	Operator LogicalOperator
	Left     Expression
	Right    Expression
}

type BooleanLiteralExpression struct {
	nodeType NodeType
	Value    bool
}

type NilLiteralExpression struct {
	nodeType NodeType
}

type StringLiteralExpression struct {
	nodeType NodeType
	Value    string
}

type NumericLiteralExpression struct {
	nodeType NodeType
	Value    int32
}

type IdentifierExpression struct {
	nodeType NodeType
	Name     string
}

type MemberExpression struct {
	nodeType NodeType
	Computed bool
	Object   Expression
	Property Expression
}

type CallExpression struct {
	nodeType  NodeType
	Callee    Expression
	Arguments []Expression
}

type ThisExpression struct {
	nodeType NodeType
}

type SuperExpression struct {
	nodeType NodeType
}

type NewExpression struct {
	nodeType  NodeType
	Callee    Expression
	Arguments []Expression
}

// Implementation of isExpression interface method
func (e VariableExpression) isExpression()       {}
func (e AssignmentExpression) isExpression()     {}
func (e BinaryExpression) isExpression()         {}
func (e UnaryExpression) isExpression()          {}
func (e LogicalExpression) isExpression()        {}
func (e BooleanLiteralExpression) isExpression() {}
func (e NilLiteralExpression) isExpression()     {}
func (e StringLiteralExpression) isExpression()  {}
func (e NumericLiteralExpression) isExpression() {}
func (e IdentifierExpression) isExpression()     {}
func (e MemberExpression) isExpression()         {}
func (e CallExpression) isExpression()           {}
func (e ThisExpression) isExpression()           {}
func (e SuperExpression) isExpression()          {}
func (e NewExpression) isExpression()            {}

// NodeType Implementation of NodeType interface method
func (e VariableExpression) NodeType() NodeType       { return e.nodeType }
func (e AssignmentExpression) NodeType() NodeType     { return e.nodeType }
func (e BinaryExpression) NodeType() NodeType         { return e.nodeType }
func (e UnaryExpression) NodeType() NodeType          { return e.nodeType }
func (e LogicalExpression) NodeType() NodeType        { return e.nodeType }
func (e BooleanLiteralExpression) NodeType() NodeType { return e.nodeType }
func (e NilLiteralExpression) NodeType() NodeType     { return e.nodeType }
func (e StringLiteralExpression) NodeType() NodeType  { return e.nodeType }
func (e NumericLiteralExpression) NodeType() NodeType { return e.nodeType }
func (e IdentifierExpression) NodeType() NodeType     { return e.nodeType }
func (e MemberExpression) NodeType() NodeType         { return e.nodeType }
func (e CallExpression) NodeType() NodeType           { return e.nodeType }
func (e ThisExpression) NodeType() NodeType           { return e.nodeType }
func (e SuperExpression) NodeType() NodeType          { return e.nodeType }
func (e NewExpression) NodeType() NodeType            { return e.nodeType }
