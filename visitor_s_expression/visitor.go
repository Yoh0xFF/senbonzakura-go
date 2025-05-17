package visitor_s_expression

import (
	"strings"

	"github.com/yoh0xff/senbonzakura/ast"
)

// SExpressionConfig defines formatting options for the visitor
type SExpressionConfig struct {
	Pretty     bool // whether to use pretty formatting
	IndentSize int  // size of each indent level
}

// SExpressionVisitor walks the AST and outputs S-expressions
type SExpressionVisitor struct {
	config      SExpressionConfig
	indentLevel int
	buffer      strings.Builder
}

// NewSExpressionVisitor creates a new visitor with default configuration
func NewSExpressionVisitor() *SExpressionVisitor {
	return NewSExpressionVisitorWithConfig(SExpressionConfig{
		Pretty:     false,
		IndentSize: 2,
	})
}

// NewSExpressionVisitorWithConfig creates a new visitor with the given configuration
func NewSExpressionVisitorWithConfig(config SExpressionConfig) *SExpressionVisitor {
	return &SExpressionVisitor{
		config:      config,
		indentLevel: 0,
		buffer:      strings.Builder{},
	}
}

// VisitStatement implements the ast.Visitor interface
func (v *SExpressionVisitor) VisitStatement(statement ast.Statement) any {
	return visitStatement(v, statement)
}

// VisitExpression implements the ast.Visitor interface
func (v *SExpressionVisitor) VisitExpression(expression ast.Expression) any {
	return visitExpression(v, expression)
}
