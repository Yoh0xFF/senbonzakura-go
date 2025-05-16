package ast

type VariableExpression struct {
	Identifier     *IdentifierExpression
	TypeAnnotation Type
	Initializer    Expression // can be nil
}

type AssignmentExpression struct {
	Operator AssignmentOperator
	Left     Expression
	Right    Expression
}

type BinaryExpression struct {
	Operator BinaryOperator
	Left     Expression
	Right    Expression
}

type UnaryExpression struct {
	Operator UnaryOperator
	Right    Expression
}

type LogicalExpression struct {
	Operator LogicalOperator
	Left     Expression
	Right    Expression
}

type BooleanLiteralExpression struct {
	Value bool
}

type NilLiteralExpression struct {
}

type StringLiteralExpression struct {
	Value string
}

type NumericLiteralExpression struct {
	Value int32
}

type IdentifierExpression struct {
	Name string
}

type MemberExpression struct {
	Computed bool
	Object   Expression
	Property Expression
}

type CallExpression struct {
	Callee    Expression
	Arguments []Expression
}

type ThisExpression struct {
}

type SuperExpression struct {
}

type NewExpression struct {
	Callee    Expression
	Arguments []Expression
}

// Implementation of isExpression interface method
func (e *VariableExpression) isExpression()       {}
func (e *AssignmentExpression) isExpression()     {}
func (e *BinaryExpression) isExpression()         {}
func (e *UnaryExpression) isExpression()          {}
func (e *LogicalExpression) isExpression()        {}
func (e *BooleanLiteralExpression) isExpression() {}
func (e *NilLiteralExpression) isExpression()     {}
func (e *StringLiteralExpression) isExpression()  {}
func (e *NumericLiteralExpression) isExpression() {}
func (e *IdentifierExpression) isExpression()     {}
func (e *MemberExpression) isExpression()         {}
func (e *CallExpression) isExpression()           {}
func (e *ThisExpression) isExpression()           {}
func (e *SuperExpression) isExpression()          {}
func (e *NewExpression) isExpression()            {}

// NodeType Implementation of NodeType interface method
func (e *VariableExpression) NodeType() NodeType       { return NodeVariableExpression }
func (e *AssignmentExpression) NodeType() NodeType     { return NodeAssignmentExpression }
func (e *BinaryExpression) NodeType() NodeType         { return NodeBinaryExpression }
func (e *UnaryExpression) NodeType() NodeType          { return NodeUnaryExpression }
func (e *LogicalExpression) NodeType() NodeType        { return NodeLogicalExpression }
func (e *BooleanLiteralExpression) NodeType() NodeType { return NodeBooleanLiteralExpression }
func (e *NilLiteralExpression) NodeType() NodeType     { return NodeNilLiteralExpression }
func (e *StringLiteralExpression) NodeType() NodeType  { return NodeStringLiteralExpression }
func (e *NumericLiteralExpression) NodeType() NodeType { return NodeNumericLiteralExpression }
func (e *IdentifierExpression) NodeType() NodeType     { return NodeIdentifierExpression }
func (e *MemberExpression) NodeType() NodeType         { return NodeMemberExpression }
func (e *CallExpression) NodeType() NodeType           { return NodeCallExpression }
func (e *ThisExpression) NodeType() NodeType           { return NodeThisExpression }
func (e *SuperExpression) NodeType() NodeType          { return NodeSuperExpression }
func (e *NewExpression) NodeType() NodeType            { return NodeNewExpression }
