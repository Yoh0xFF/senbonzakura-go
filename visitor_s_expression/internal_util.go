package visitor_s_expression

import "strings"

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
