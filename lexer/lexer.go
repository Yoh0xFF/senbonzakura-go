package lexer

import (
	"fmt"
)

// Lexer lazily pulls tokens from a stream
type Lexer struct {
	source string
	index  int
	rules  []RegexRule
}

// NewLexer creates a new lexer instance
func NewLexer(source string) *Lexer {
	return &Lexer{
		source: source,
		index:  0,
		rules:  GetRegexRules(),
	}
}

// NextToken obtains the next token from the source
func (l *Lexer) NextToken() Token {
	currentIndex := l.index

	// Check if we're at the end of the source
	if currentIndex >= len(l.source) {
		return Token{
			TokenType: TokenEnd,
			Start:     currentIndex,
			End:       currentIndex,
		}
	}

	// Slice the string starting from the current position
	expression := l.source[currentIndex:]

	for _, rule := range l.rules {
		match := rule.Pattern.FindStringIndex(expression)
		if match == nil {
			continue // Try to match other token
		}

		// match[0] is the start index (should be 0 for our patterns)
		// match[1] is the end index
		tokenText := expression[match[0]:match[1]]
		tokenLen := len(tokenText)

		switch rule.TokenType {
		case TokenWhitespace, TokenSingleLineComment, TokenMultiLineComment:
			// Skip whitespace and comments
			l.index = currentIndex + tokenLen
			return l.NextToken()
		default:
			// Return the matched token
			l.index = currentIndex + tokenLen
			return Token{
				TokenType: rule.TokenType,
				Start:     currentIndex,
				End:       currentIndex + tokenLen,
			}
		}
	}

	// If we get here, no token matched
	panic(fmt.Sprintf(
		"Invalid token at index %d, remaining text: '%s'",
		currentIndex,
		l.source[currentIndex:],
	))
}

// Clone creates a copy of the lexer at its current state
func (l *Lexer) Clone() *Lexer {
	return &Lexer{
		source: l.source,
		index:  l.index,
		rules:  l.rules,
	}
}

// Remaining returns the remaining source text
func (l *Lexer) Remaining() string {
	if l.index >= len(l.source) {
		return ""
	}
	return l.source[l.index:]
}

// Position returns the current position in the source
func (l *Lexer) Position() int {
	return l.index
}
