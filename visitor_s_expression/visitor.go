package visitor_s_expression

import (
	"github.com/yoh0xff/senbonzakura/ast"
	"strings"
)

type SExpressionVisitor struct {
	indentLevel int
	buffer      strings.Builder
}

func NewSExpressionVisitor() *SExpressionVisitor {
	return &SExpressionVisitor{
		indentLevel: 0,
		buffer:      strings.Builder{},
	}
}

func (v *SExpressionVisitor) VisitStatement(statement ast.Statement) any {
	return visitStatement(v, statement)
}

func (v *SExpressionVisitor) VisitExpression(expression ast.Expression) any {
	return visitExpression(v, expression)
}

// beginExpression starts a new S-expression with the given tag
func (v *SExpressionVisitor) beginExpression(tag string) {
	v.buffer.WriteString("(")
	v.buffer.WriteString(tag)
	v.indentLevel++
}

// endExpression closes the current S-expression
func (v *SExpressionVisitor) endExpression() {
	v.indentLevel--
	v.buffer.WriteString(")")
}

// writeSpaceOrNewLine writes a space or a newline based on formatting rules
func (v *SExpressionVisitor) writeSpaceOrNewLine() {
	v.buffer.WriteString(" ")
}

// writeString writes a string to the output
func (v *SExpressionVisitor) writeString(s string) {
	v.buffer.WriteString(s)
}

// String returns the final S-expression string
func (v *SExpressionVisitor) String() string {
	return v.buffer.String()
}
