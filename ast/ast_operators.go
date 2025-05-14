package ast

import "fmt"

// AssignmentOperator represents different assignment operators
type AssignmentOperator int

const (
	Assign AssignmentOperator = iota
	AssignAdd
	AssignSubtract
	AssignMultiply
	AssignDivide
)

// String returns the string representation of an AssignmentOperator
func (op AssignmentOperator) String() string {
	switch op {
	case Assign:
		return "="
	case AssignAdd:
		return "+="
	case AssignSubtract:
		return "-="
	case AssignMultiply:
		return "*="
	case AssignDivide:
		return "/="
	default:
		return fmt.Sprintf("Unknown assignment operator: %d", op)
	}
}

// BinaryOperator represents different binary operators
type BinaryOperator int

const (
	Add BinaryOperator = iota
	Subtract
	Multiply
	Divide
	Equal
	NotEqual
	GreaterThan
	GreaterThanOrEqualTo
	LessThan
	LessThanOrEqualTo
)

// String returns the string representation of a BinaryOperator
func (op BinaryOperator) String() string {
	switch op {
	case Add:
		return "+"
	case Subtract:
		return "-"
	case Multiply:
		return "*"
	case Divide:
		return "/"
	case Equal:
		return "=="
	case NotEqual:
		return "!="
	case GreaterThan:
		return ">"
	case GreaterThanOrEqualTo:
		return ">="
	case LessThan:
		return "<"
	case LessThanOrEqualTo:
		return "<="
	default:
		return fmt.Sprintf("Unknown binary operator: %d", op)
	}
}

// UnaryOperator represents different unary operators
type UnaryOperator int

const (
	Plus UnaryOperator = iota
	Minus
	Not
)

// String returns the string representation of a UnaryOperator
func (op UnaryOperator) String() string {
	switch op {
	case Plus:
		return "+"
	case Minus:
		return "-"
	case Not:
		return "!"
	default:
		return fmt.Sprintf("Unknown unary operator: %d", op)
	}
}

// LogicalOperator represents different logical operators
type LogicalOperator int

const (
	And LogicalOperator = iota
	Or
)

// String returns the string representation of a LogicalOperator
func (op LogicalOperator) String() string {
	switch op {
	case And:
		return "&&"
	case Or:
		return "||"
	default:
		return fmt.Sprintf("Unknown logical operator: %d", op)
	}
}
