package ast

type NodeType int

// Statement node types
const (
	NodeUnknown NodeType = iota

	/**
	*******************
	* Statement types *
	*******************
	 */

	NodeProgramStatement
	NodeBlockStatement
	NodeEmptyStatement
	NodeExpressionStatement
	NodeVariableDeclarationStatement
	NodeIfStatement
	NodeWhileStatement
	NodeDoWhileStatement
	NodeForStatement
	NodeFunctionDeclarationStatement
	NodeReturnStatement
	NodeClassDeclarationStatement

	/**
	********************
	* Expression types *
	********************
	 */

	NodeVariableExpression
	NodeAssignmentExpression
	NodeBinaryExpression
	NodeUnaryExpression
	NodeLogicalExpression
	NodeBooleanLiteralExpression
	NodeNilLiteralExpression
	NodeStringLiteralExpression
	NodeNumericLiteralExpression
	NodeIdentifierExpression
	NodeMemberExpression
	NodeCallExpression
	NodeThisExpression
	NodeSuperExpression
	NodeNewExpression
)

// String representation for debugging
func (t NodeType) String() string {
	switch t {
	case NodeUnknown:
		return "Unknown"

	// Statements
	case NodeProgramStatement:
		return "ProgramStatement"
	case NodeBlockStatement:
		return "BlockStatement"
	case NodeEmptyStatement:
		return "EmptyStatement"
	case NodeExpressionStatement:
		return "ExpressionStatement"
	case NodeVariableDeclarationStatement:
		return "VariableDeclarationStatement"
	case NodeIfStatement:
		return "IfStatement"
	case NodeWhileStatement:
		return "WhileStatement"
	case NodeDoWhileStatement:
		return "DoWhileStatement"
	case NodeForStatement:
		return "ForStatement"
	case NodeFunctionDeclarationStatement:
		return "FunctionDeclarationStatement"
	case NodeReturnStatement:
		return "ReturnStatement"
	case NodeClassDeclarationStatement:
		return "ClassDeclarationStatement"

	// Expressions
	case NodeVariableExpression:
		return "VariableExpression"
	case NodeAssignmentExpression:
		return "AssignmentExpression"
	case NodeBinaryExpression:
		return "BinaryExpression"
	case NodeUnaryExpression:
		return "UnaryExpression"
	case NodeLogicalExpression:
		return "LogicalExpression"
	case NodeBooleanLiteralExpression:
		return "BooleanLiteralExpression"
	case NodeNilLiteralExpression:
		return "NilLiteralExpression"
	case NodeStringLiteralExpression:
		return "StringLiteralExpression"
	case NodeNumericLiteralExpression:
		return "NumericLiteralExpression"
	case NodeIdentifierExpression:
		return "IdentifierExpression"
	case NodeMemberExpression:
		return "MemberExpression"
	case NodeCallExpression:
		return "CallExpression"
	case NodeThisExpression:
		return "ThisExpression"
	case NodeSuperExpression:
		return "SuperExpression"
	case NodeNewExpression:
		return "NewExpression"
	default:
		return "InvalidNodeType"
	}
}

// IsStatement Helper methods for node categories
func (t NodeType) IsStatement() bool {
	return t >= NodeProgramStatement && t <= NodeClassDeclarationStatement
}

// IsExpression Helper methods for node categories
func (t NodeType) IsExpression() bool {
	return t >= NodeVariableExpression && t <= NodeNewExpression
}

// IsLiteral Helper methods for node categories
func (t NodeType) IsLiteral() bool {
	switch t {
	case NodeBooleanLiteralExpression, NodeNilLiteralExpression, NodeStringLiteralExpression, NodeNumericLiteralExpression:
		return true
	default:
		return false
	}
}
