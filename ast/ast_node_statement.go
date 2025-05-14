package ast

type ProgramStatement struct {
	Body Statement
}

type BlockStatement struct {
	Body Statement
}

type EmptyStatement struct {
}

type ExpressionStatement struct {
	Expression Expression
}

type VariableDeclarationStatement struct {
	Variables []Expression
}

type IfStatement struct {
	Condition   Expression
	Consequent  Statement
	Alternative Statement // can be nil
}

type WhileStatement struct {
	Condition Expression
	Body      Statement
}

type DoWhileStatement struct {
	Body      Statement
	Condition Expression
}

type ForStatement struct {
	Initializer Statement  // can be nil
	Condition   Expression // can be nil
	Increment   Expression // can be nil
	Body        Statement
}

type FunctionDeclarationStatement struct {
	Name       Expression
	Parameters []Parameter
	ReturnType Type
	Body       Statement
}

type ReturnStatement struct {
	Argument Expression // can be nil
}

type ClassDeclarationStatement struct {
	Name       Expression
	SuperClass Expression // can be nil
	Body       Statement
}

// Implementation of isStatement interface method
func (s *ProgramStatement) isStatement()             {}
func (s *BlockStatement) isStatement()               {}
func (s *EmptyStatement) isStatement()               {}
func (s *ExpressionStatement) isStatement()          {}
func (s *VariableDeclarationStatement) isStatement() {}
func (s *IfStatement) isStatement()                  {}
func (s *WhileStatement) isStatement()               {}
func (s *DoWhileStatement) isStatement()             {}
func (s *ForStatement) isStatement()                 {}
func (s *FunctionDeclarationStatement) isStatement() {}
func (s *ReturnStatement) isStatement()              {}
func (s *ClassDeclarationStatement) isStatement()    {}

// NodeType Implementation of NodeType interface method
func (s *ProgramStatement) NodeType() NodeType             { return NodeProgramStatement }
func (s *BlockStatement) NodeType() NodeType               { return NodeBlockStatement }
func (s *EmptyStatement) NodeType() NodeType               { return NodeEmptyStatement }
func (s *ExpressionStatement) NodeType() NodeType          { return NodeExpressionStatement }
func (s *VariableDeclarationStatement) NodeType() NodeType { return NodeVariableDeclarationStatement }
func (s *IfStatement) NodeType() NodeType                  { return NodeIfStatement }
func (s *WhileStatement) NodeType() NodeType               { return NodeWhileStatement }
func (s *DoWhileStatement) NodeType() NodeType             { return NodeDoWhileStatement }
func (s *ForStatement) NodeType() NodeType                 { return NodeForStatement }
func (s *FunctionDeclarationStatement) NodeType() NodeType { return NodeFunctionDeclarationStatement }
func (s *ReturnStatement) NodeType() NodeType              { return NodeReturnStatement }
func (s *ClassDeclarationStatement) NodeType() NodeType    { return NodeClassDeclarationStatement }
