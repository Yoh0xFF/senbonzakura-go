package ast

import "fmt"

// AssignmentOperator represents different assignment operators
type AssignmentOperator int

const (
	OperatorAssign AssignmentOperator = iota
	OperatorAssignAdd
	OperatorAssignSubtract
	OperatorAssignMultiply
	OperatorAssignDivide
)

// String returns the string representation of an AssignmentOperator
func (op AssignmentOperator) String() string {
	switch op {
	case OperatorAssign:
		return "="
	case OperatorAssignAdd:
		return "+="
	case OperatorAssignSubtract:
		return "-="
	case OperatorAssignMultiply:
		return "*="
	case OperatorAssignDivide:
		return "/="
	default:
		return fmt.Sprintf("Unknown assignment operator: %d", op)
	}
}

// BinaryOperator represents different binary operators
type BinaryOperator int

const (
	OperatorAdd BinaryOperator = iota
	OperatorSubtract
	OperatorMultiply
	OperatorDivide
	OperatorEqual
	OperatorNotEqual
	OperatorGreaterThan
	OperatorGreaterThanOrEqualTo
	OperatorLessThan
	OperatorLessThanOrEqualTo
)

// String returns the string representation of a BinaryOperator
func (op BinaryOperator) String() string {
	switch op {
	case OperatorAdd:
		return "+"
	case OperatorSubtract:
		return "-"
	case OperatorMultiply:
		return "*"
	case OperatorDivide:
		return "/"
	case OperatorEqual:
		return "=="
	case OperatorNotEqual:
		return "!="
	case OperatorGreaterThan:
		return ">"
	case OperatorGreaterThanOrEqualTo:
		return ">="
	case OperatorLessThan:
		return "<"
	case OperatorLessThanOrEqualTo:
		return "<="
	default:
		return fmt.Sprintf("Unknown binary operator: %d", op)
	}
}

// UnaryOperator represents different unary operators
type UnaryOperator int

const (
	OperatorPlus UnaryOperator = iota
	OperatorMinus
	OperatorNot
)

// String returns the string representation of a UnaryOperator
func (op UnaryOperator) String() string {
	switch op {
	case OperatorPlus:
		return "+"
	case OperatorMinus:
		return "-"
	case OperatorNot:
		return "!"
	default:
		return fmt.Sprintf("Unknown unary operator: %d", op)
	}
}

// LogicalOperator represents different logical operators
type LogicalOperator int

const (
	OperatorAnd LogicalOperator = iota
	OperatorOr
)

// String returns the string representation of a LogicalOperator
func (op LogicalOperator) String() string {
	switch op {
	case OperatorAnd:
		return "&&"
	case OperatorOr:
		return "||"
	default:
		return fmt.Sprintf("Unknown logical operator: %d", op)
	}
}
