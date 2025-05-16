package visitor_s_expression

import (
	"github.com/yoh0xff/senbonzakura/ast"
	"strings"
)

// SExpressionConfig defines formatting options for the visitor
type SExpressionConfig struct {
	Pretty     bool // whether to use pretty formatting
	IndentSize int  // size of each indent level
}

// DefaultConfig returns the default configuration
func DefaultConfig() SExpressionConfig {
	return SExpressionConfig{
		Pretty:     false,
		IndentSize: 2,
	}
}

// SExpressionVisitor walks the AST and outputs S-expressions
type SExpressionVisitor struct {
	config      SExpressionConfig
	indentLevel int
	buffer      strings.Builder
}

// NewSExpressionVisitor creates a new visitor with default configuration
func NewSExpressionVisitor() *SExpressionVisitor {
	return NewSExpressionVisitorWithConfig(DefaultConfig())
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

// writeIndent writes the appropriate indentation based on the current indent level
func (v *SExpressionVisitor) writeIndent() {
	if v.config.Pretty && v.indentLevel > 0 {
		indent := strings.Repeat(" ", v.indentLevel*v.config.IndentSize)
		v.buffer.WriteString(indent)
	}
}

// beginExpression starts a new S-expression with the given tag
func (v *SExpressionVisitor) beginExpression(tag string) {
	v.writeIndent()
	v.buffer.WriteString("(")
	v.buffer.WriteString(tag)

	if v.config.Pretty {
		v.indentLevel++
	}
}

// endExpression closes the current S-expression
func (v *SExpressionVisitor) endExpression() {
	if v.config.Pretty {
		v.indentLevel--
	}

	v.buffer.WriteString(")")
}

// writeSpaceOrNewLine writes a space or a newline based on formatting rules
func (v *SExpressionVisitor) writeSpaceOrNewLine() {
	if v.config.Pretty {
		v.buffer.WriteString("\n")
		v.writeIndent()
	} else {
		v.buffer.WriteString(" ")
	}
}

// writeString writes a string to the output
func (v *SExpressionVisitor) writeString(s string) {
	v.buffer.WriteString(s)
}

// String returns the final S-expression string
func (v *SExpressionVisitor) String() string {
	return v.buffer.String()
}
