package ast

import "fmt"

// Type represents different type annotations in the AST
type Type interface {
	String() string
	isType()
}

// PrimitiveTypeKind represents primitive types
type PrimitiveTypeKind int

const (
	NumberType PrimitiveTypeKind = iota
	BooleanType
	StringType
)

// PrimitiveType represents a primitive type annotation
type PrimitiveType struct {
	Kind PrimitiveTypeKind
}

// ArrayType represents an array type annotation
type ArrayType struct {
	ElementType Type
}

// FunctionType represents a function type annotation
type FunctionType struct {
	Params     []Type
	ReturnType Type
}

// ClassType represents a class type annotation
type ClassType struct {
	Name       string
	SuperClass *string
}

// GenericType represents a generic type annotation
type GenericType struct {
	Base     string
	TypeArgs []Type
}

// VoidType represents a void type annotation
type VoidType struct{}

// Implementation of Type interface for all types
func (t PrimitiveType) isType() {}
func (t ArrayType) isType()     {}
func (t FunctionType) isType()  {}
func (t ClassType) isType()     {}
func (t GenericType) isType()   {}
func (t VoidType) isType()      {}

// String implementations
func (t PrimitiveType) String() string {
	switch t.Kind {
	case NumberType:
		return "Number"
	case BooleanType:
		return "Boolean"
	case StringType:
		return "String"
	default:
		return fmt.Sprintf("Unknown primitive type: %d", t.Kind)
	}
}

func (t ArrayType) String() string {
	return fmt.Sprintf("Array<%s>", t.ElementType.String())
}

func (t FunctionType) String() string {
	paramStrings := make([]string, len(t.Params))
	for i, param := range t.Params {
		paramStrings[i] = param.String()
	}
	return fmt.Sprintf("Function(%v) -> %s", paramStrings, t.ReturnType.String())
}

func (t ClassType) String() string {
	if t.SuperClass != nil {
		return fmt.Sprintf("Class<%s extends %s>", t.Name, *t.SuperClass)
	}
	return fmt.Sprintf("Class<%s>", t.Name)
}

func (t GenericType) String() string {
	argStrings := make([]string, len(t.TypeArgs))
	for i, arg := range t.TypeArgs {
		argStrings[i] = arg.String()
	}
	return fmt.Sprintf("%s<%v>", t.Base, argStrings)
}

func (t VoidType) String() string {
	return "void"
}
