package ast

type ProgramStatement struct {
	nodeType NodeType
	Body     Statement
}

type BlockStatement struct {
	nodeType NodeType
	Body     Statement
}

type EmptyStatement struct {
	nodeType NodeType
}

type ExpressionStatement struct {
	nodeType   NodeType
	Expression Expression
}

type VariableDeclarationStatement struct {
	nodeType  NodeType
	Variables []Expression
}

type IfStatement struct {
	nodeType    NodeType
	Condition   Expression
	Consequent  Statement
	Alternative Statement // can be nil
}

type WhileStatement struct {
	nodeType  NodeType
	Condition Expression
	Body      Statement
}

type DoWhileStatement struct {
	nodeType  NodeType
	Body      Statement
	Condition Expression
}

type ForStatement struct {
	nodeType    NodeType
	Initializer Statement  // can be nil
	Condition   Expression // can be nil
	Increment   Expression // can be nil
	Body        Statement
}

type FunctionDeclarationStatement struct {
	nodeType   NodeType
	Name       Expression
	Parameters []Parameter
	ReturnType Type
	Body       Statement
}

type ReturnStatement struct {
	nodeType NodeType
	Argument Expression // can be nil
}

type ClassDeclarationStatement struct {
	nodeType   NodeType
	Name       Expression
	SuperClass Expression // can be nil
	Body       Statement
}

// Implementation of isStatement interface method
func (s ProgramStatement) isStatement()             {}
func (s BlockStatement) isStatement()               {}
func (s EmptyStatement) isStatement()               {}
func (s ExpressionStatement) isStatement()          {}
func (s VariableDeclarationStatement) isStatement() {}
func (s IfStatement) isStatement()                  {}
func (s WhileStatement) isStatement()               {}
func (s DoWhileStatement) isStatement()             {}
func (s ForStatement) isStatement()                 {}
func (s FunctionDeclarationStatement) isStatement() {}
func (s ReturnStatement) isStatement()              {}
func (s ClassDeclarationStatement) isStatement()    {}

// NodeType Implementation of NodeType interface method
func (s ProgramStatement) NodeType() NodeType             { return s.nodeType }
func (s BlockStatement) NodeType() NodeType               { return s.nodeType }
func (s EmptyStatement) NodeType() NodeType               { return s.nodeType }
func (s ExpressionStatement) NodeType() NodeType          { return s.nodeType }
func (s VariableDeclarationStatement) NodeType() NodeType { return s.nodeType }
func (s IfStatement) NodeType() NodeType                  { return s.nodeType }
func (s WhileStatement) NodeType() NodeType               { return s.nodeType }
func (s DoWhileStatement) NodeType() NodeType             { return s.nodeType }
func (s ForStatement) NodeType() NodeType                 { return s.nodeType }
func (s FunctionDeclarationStatement) NodeType() NodeType { return s.nodeType }
func (s ReturnStatement) NodeType() NodeType              { return s.nodeType }
func (s ClassDeclarationStatement) NodeType() NodeType    { return s.nodeType }
