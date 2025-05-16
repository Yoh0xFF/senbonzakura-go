package lexer

import "fmt"

// TokenType represents the different types of tokens in the language
type TokenType int

const (
	// TokenWhitespace and comments

	TokenWhitespace TokenType = iota
	TokenSingleLineComment
	TokenMultiLineComment

	// Symbols

	TokenStatementEnd
	TokenOpeningBrace
	TokenClosingBrace
	TokenOpeningParenthesis
	TokenClosingParenthesis
	TokenOpeningBracket
	TokenClosingBracket
	TokenComma
	TokenDot
	TokenColon

	// Keywords

	TokenLetKeyword
	TokenIfKeyword
	TokenElseKeyword
	TokenWhileKeyword
	TokenDoKeyword
	TokenForKeyword
	TokenDefKeyword
	TokenReturnKeyword
	TokenClassKeyword
	TokenExtendsKeyword
	TokenThisKeyword
	TokenSuperKeyword
	TokenNewKeyword
	TokenTypeKeyword
	TokenNumberTypeKeyword
	TokenStringTypeKeyword
	TokenBooleanTypeKeyword
	TokenVoidTypeKeyword

	// Identifier

	TokenIdentifier

	// Equality operators

	TokenEqualityOperator

	// Assignment operators

	TokenSimpleAssignmentOperator
	TokenComplexAssignmentOperator

	// Math operators

	TokenAdditiveOperator
	TokenFactorOperator

	// Relational operators

	TokenRelationalOperator

	// Logical operators

	TokenLogicalAndOperator
	TokenLogicalOrOperator
	TokenLogicalNotOperator

	// Literals

	TokenBoolean
	TokenNil
	TokenNumber
	TokenString

	// TokenEnd

	TokenEnd
)

// TokenString returns a string representation of the TokenType
func (t TokenType) String() string {
	switch t {
	case TokenWhitespace:
		return "TokenWhitespace"
	case TokenSingleLineComment:
		return "TokenSingleLineComment"
	case TokenMultiLineComment:
		return "TokenMultiLineComment"
	case TokenStatementEnd:
		return "TokenStatementEnd"
	case TokenOpeningBrace:
		return "TokenOpeningBrace"
	case TokenClosingBrace:
		return "TokenClosingBrace"
	case TokenOpeningParenthesis:
		return "TokenOpeningParenthesis"
	case TokenClosingParenthesis:
		return "TokenClosingParenthesis"
	case TokenOpeningBracket:
		return "TokenOpeningBracket"
	case TokenClosingBracket:
		return "TokenClosingBracket"
	case TokenComma:
		return "TokenComma"
	case TokenDot:
		return "TokenDot"
	case TokenColon:
		return "TokenColon"
	case TokenLetKeyword:
		return "TokenLetKeyword"
	case TokenIfKeyword:
		return "TokenIfKeyword"
	case TokenElseKeyword:
		return "TokenElseKeyword"
	case TokenWhileKeyword:
		return "TokenWhileKeyword"
	case TokenDoKeyword:
		return "TokenDoKeyword"
	case TokenForKeyword:
		return "TokenForKeyword"
	case TokenDefKeyword:
		return "TokenDefKeyword"
	case TokenReturnKeyword:
		return "TokenReturnKeyword"
	case TokenClassKeyword:
		return "TokenClassKeyword"
	case TokenExtendsKeyword:
		return "TokenExtendsKeyword"
	case TokenThisKeyword:
		return "TokenThisKeyword"
	case TokenSuperKeyword:
		return "TokenSuperKeyword"
	case TokenNewKeyword:
		return "TokenNewKeyword"
	case TokenTypeKeyword:
		return "TokenTypeKeyword"
	case TokenNumberTypeKeyword:
		return "TokenNumberTypeKeyword"
	case TokenStringTypeKeyword:
		return "TokenStringTypeKeyword"
	case TokenBooleanTypeKeyword:
		return "TokenBooleanTypeKeyword"
	case TokenVoidTypeKeyword:
		return "TokenVoidTypeKeyword"
	case TokenIdentifier:
		return "TokenIdentifier"
	case TokenEqualityOperator:
		return "TokenEqualityOperator"
	case TokenSimpleAssignmentOperator:
		return "TokenSimpleAssignmentOperator"
	case TokenComplexAssignmentOperator:
		return "TokenComplexAssignmentOperator"
	case TokenAdditiveOperator:
		return "TokenAdditiveOperator"
	case TokenFactorOperator:
		return "TokenFactorOperator"
	case TokenRelationalOperator:
		return "TokenRelationalOperator"
	case TokenLogicalAndOperator:
		return "TokenLogicalAndOperator"
	case TokenLogicalOrOperator:
		return "TokenLogicalOrOperator"
	case TokenLogicalNotOperator:
		return "TokenLogicalNotOperator"
	case TokenBoolean:
		return "TokenBoolean"
	case TokenNil:
		return "TokenNil"
	case TokenNumber:
		return "TokenNumber"
	case TokenString:
		return "TokenString"
	case TokenEnd:
		return "TokenEnd"
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

// TokenString returns a string representation of the Token
func (t Token) String() string {
	return fmt.Sprintf("Token (%s, %d, %d)", t.TokenType.String(), t.Start, t.End)
}

// Equal compares two tokens for equality
func (t Token) Equal(other Token) bool {
	return t.TokenType == other.TokenType && t.Start == other.Start && t.End == other.End
}
