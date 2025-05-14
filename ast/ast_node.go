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

	NodeProgram
	NodeBlock
	NodeEmpty
	NodeExpressionStmt
	NodeVariableDeclaration
	NodeIf
	NodeWhile
	NodeDoWhile
	NodeFor
	NodeFunctionDeclaration
	NodeReturn
	NodeClassDeclaration

	/**
	********************
	* Expression types *
	********************
	 */

	NodeVariable
	NodeAssignment
	NodeBinary
	NodeUnary
	NodeLogical
	NodeBooleanLiteral
	NodeNilLiteral
	NodeStringLiteral
	NodeNumericLiteral
	NodeIdentifier
	NodeMember
	NodeCall
	NodeThis
	NodeSuper
	NodeNew
)

// String representation for debugging
func (t NodeType) String() string {
	switch t {
	case NodeUnknown:
		return "Unknown"

	// Statements
	case NodeProgram:
		return "Program"
	case NodeBlock:
		return "Block"
	case NodeEmpty:
		return "Empty"
	case NodeExpressionStmt:
		return "ExpressionStatement"
	case NodeVariableDeclaration:
		return "VariableDeclaration"
	case NodeIf:
		return "If"
	case NodeWhile:
		return "While"
	case NodeDoWhile:
		return "DoWhile"
	case NodeFor:
		return "For"
	case NodeFunctionDeclaration:
		return "FunctionDeclaration"
	case NodeReturn:
		return "Return"
	case NodeClassDeclaration:
		return "ClassDeclaration"

	// Expressions
	case NodeVariable:
		return "Variable"
	case NodeAssignment:
		return "Assignment"
	case NodeBinary:
		return "Binary"
	case NodeUnary:
		return "Unary"
	case NodeLogical:
		return "Logical"
	case NodeBooleanLiteral:
		return "BooleanLiteral"
	case NodeNilLiteral:
		return "NilLiteral"
	case NodeStringLiteral:
		return "StringLiteral"
	case NodeNumericLiteral:
		return "NumericLiteral"
	case NodeIdentifier:
		return "Identifier"
	case NodeMember:
		return "Member"
	case NodeCall:
		return "Call"
	case NodeThis:
		return "This"
	case NodeSuper:
		return "Super"
	case NodeNew:
		return "New"
	default:
		return "InvalidNodeType"
	}
}

// IsStatement Helper methods for node categories
func (t NodeType) IsStatement() bool {
	return t >= NodeProgram && t <= NodeClassDeclaration
}

// IsExpression Helper methods for node categories
func (t NodeType) IsExpression() bool {
	return t >= NodeVariable && t <= NodeNew
}

// IsLiteral Helper methods for node categories
func (t NodeType) IsLiteral() bool {
	switch t {
	case NodeBooleanLiteral, NodeNilLiteral, NodeStringLiteral, NodeNumericLiteral:
		return true
	default:
		return false
	}
}
