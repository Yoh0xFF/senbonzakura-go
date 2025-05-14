package lexer

import "fmt"

// TokenType represents the different types of tokens in the language
type TokenType int

const (
	// **************************************************
	// Whitespace and comments
	// **************************************************

	Whitespace TokenType = iota
	SingleLineComment
	MultiLineComment

	// **************************************************
	// Symbols
	// **************************************************

	StatementEnd
	OpeningBrace
	ClosingBrace
	OpeningParenthesis
	ClosingParenthesis
	OpeningBracket
	ClosingBracket
	Comma
	Dot
	Colon

	// **************************************************
	// Keywords
	// **************************************************

	LetKeyword
	IfKeyword
	ElseKeyword
	WhileKeyword
	DoKeyword
	ForKeyword
	DefKeyword
	ReturnKeyword
	ClassKeyword
	ExtendsKeyword
	ThisKeyword
	SuperKeyword
	NewKeyword
	TypeKeyword
	NumberTypeKeyword
	StringTypeKeyword
	BooleanTypeKeyword
	VoidTypeKeyword

	// **************************************************
	// Identifier
	// **************************************************

	Identifier

	// **************************************************
	// Equality operators
	// **************************************************

	EqualityOperator

	// **************************************************
	// Assignment operators
	// **************************************************

	SimpleAssignmentOperator
	ComplexAssignmentOperator

	// **************************************************
	// Math operators
	// **************************************************

	AdditiveOperator
	FactorOperator

	// **************************************************
	// Relational operators
	// **************************************************

	RelationalOperator

	// **************************************************
	// Logical operators
	// **************************************************

	LogicalAndOperator
	LogicalOrOperator
	LogicalNotOperator

	// **************************************************
	// Literals
	// **************************************************

	Boolean
	Nil
	Number
	String

	// **************************************************
	// End
	// **************************************************

	End
)

// String returns a string representation of the TokenType
func (t TokenType) String() string {
	switch t {
	case Whitespace:
		return "Whitespace"
	case SingleLineComment:
		return "SingleLineComment"
	case MultiLineComment:
		return "MultiLineComment"
	case StatementEnd:
		return "StatementEnd"
	case OpeningBrace:
		return "OpeningBrace"
	case ClosingBrace:
		return "ClosingBrace"
	case OpeningParenthesis:
		return "OpeningParenthesis"
	case ClosingParenthesis:
		return "ClosingParenthesis"
	case OpeningBracket:
		return "OpeningBracket"
	case ClosingBracket:
		return "ClosingBracket"
	case Comma:
		return "Comma"
	case Dot:
		return "Dot"
	case Colon:
		return "Colon"
	case LetKeyword:
		return "LetKeyword"
	case IfKeyword:
		return "IfKeyword"
	case ElseKeyword:
		return "ElseKeyword"
	case WhileKeyword:
		return "WhileKeyword"
	case DoKeyword:
		return "DoKeyword"
	case ForKeyword:
		return "ForKeyword"
	case DefKeyword:
		return "DefKeyword"
	case ReturnKeyword:
		return "ReturnKeyword"
	case ClassKeyword:
		return "ClassKeyword"
	case ExtendsKeyword:
		return "ExtendsKeyword"
	case ThisKeyword:
		return "ThisKeyword"
	case SuperKeyword:
		return "SuperKeyword"
	case NewKeyword:
		return "NewKeyword"
	case TypeKeyword:
		return "TypeKeyword"
	case NumberTypeKeyword:
		return "NumberTypeKeyword"
	case StringTypeKeyword:
		return "StringTypeKeyword"
	case BooleanTypeKeyword:
		return "BooleanTypeKeyword"
	case VoidTypeKeyword:
		return "VoidTypeKeyword"
	case Identifier:
		return "Identifier"
	case EqualityOperator:
		return "EqualityOperator"
	case SimpleAssignmentOperator:
		return "SimpleAssignmentOperator"
	case ComplexAssignmentOperator:
		return "ComplexAssignmentOperator"
	case AdditiveOperator:
		return "AdditiveOperator"
	case FactorOperator:
		return "FactorOperator"
	case RelationalOperator:
		return "RelationalOperator"
	case LogicalAndOperator:
		return "LogicalAndOperator"
	case LogicalOrOperator:
		return "LogicalOrOperator"
	case LogicalNotOperator:
		return "LogicalNotOperator"
	case Boolean:
		return "Boolean"
	case Nil:
		return "Nil"
	case Number:
		return "Number"
	case String:
		return "String"
	case End:
		return "End"
	default:
		return fmt.Sprintf("TokenType(%d)", t)
	}
}

// Token represents a lexical token in the source code
type Token struct {
	TokenType TokenType
	Start     int // Start position
	End       int // End position
}

// String returns a string representation of the Token
func (t Token) String() string {
	return fmt.Sprintf("Token (%s, %d, %d)", t.TokenType.String(), t.Start, t.End)
}

// Equal compares two tokens for equality
func (t Token) Equal(other Token) bool {
	return t.TokenType == other.TokenType && t.Start == other.Start && t.End == other.End
}
