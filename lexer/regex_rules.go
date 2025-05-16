package lexer

import (
	"log"
	"regexp"
	"sync"
)

// RegexRule represents a regex pattern and its associated token type
type RegexRule struct {
	Pattern   *regexp.Regexp
	TokenType TokenType
}

var (
	regexRules []RegexRule
	regexOnce  sync.Once
)

// GetRegexRules returns the initialized regex rules (lazy initialization)
func GetRegexRules() []RegexRule {
	regexOnce.Do(initRegexRules)
	return regexRules
}

func initRegexRules() {
	// Compile regex patterns
	patterns := []struct {
		pattern   string
		tokenType TokenType
		desc      string
	}{
		// TokenWhitespace
		{`^\s+`, TokenWhitespace, "whitespace"},

		// Comments
		{`^//.*`, TokenSingleLineComment, "single line comments"},
		{`^/\*[\s\S]*?\*/`, TokenMultiLineComment, "multi line comments"},

		// Symbols, delimiters
		{`^;`, TokenStatementEnd, "statement end (;) symbol"},
		{`^\{`, TokenOpeningBrace, "opening brace ({) symbol"},
		{`^}`, TokenClosingBrace, "closing brace (}) symbol"},
		{`^\(`, TokenOpeningParenthesis, "opening parenthesis (() symbol"},
		{`^\)`, TokenClosingParenthesis, "closing parenthesis ()) symbol"},
		{`^\[`, TokenOpeningBracket, "opening bracket ([) symbol"},
		{`^]`, TokenClosingBracket, "closing bracket (]) symbol"},
		{`^,`, TokenComma, "comma (,) symbol"},
		{`^\.`, TokenDot, "dot (.) symbol"},
		{`^:`, TokenColon, "colon (:) symbol"},

		// Keywords
		{`^\btrue\b`, TokenBoolean, "the 'true' keyword"},
		{`^\bfalse\b`, TokenBoolean, "the 'false' keyword"},
		{`^\bnil\b`, TokenNil, "the 'nil' keyword"},
		{`^\blet\b`, TokenLetKeyword, "the 'let' keyword"},
		{`^\bif\b`, TokenIfKeyword, "the 'if' keyword"},
		{`^\belse\b`, TokenElseKeyword, "the 'else' keyword"},
		{`^\bwhile\b`, TokenWhileKeyword, "the 'while' keyword"},
		{`^\bdo\b`, TokenDoKeyword, "the 'do' keyword"},
		{`^\bfor\b`, TokenForKeyword, "the 'for' keyword"},
		{`^\bdef\b`, TokenDefKeyword, "the 'def' keyword"},
		{`^\breturn\b`, TokenReturnKeyword, "the 'return' keyword"},
		{`^\bclass\b`, TokenClassKeyword, "the 'class' keyword"},
		{`^\bextends\b`, TokenExtendsKeyword, "the 'extends' keyword"},
		{`^\bthis\b`, TokenThisKeyword, "the 'this' keyword"},
		{`^\bsuper\b`, TokenSuperKeyword, "the 'super' keyword"},
		{`^\bnew\b`, TokenNewKeyword, "the 'new' keyword"},
		{`^\btype\b`, TokenTypeKeyword, "type keyword"},
		{`^\bnumber\b`, TokenNumberTypeKeyword, "number type"},
		{`^\bstring\b`, TokenStringTypeKeyword, "string type"},
		{`^\bboolean\b`, TokenBooleanTypeKeyword, "boolean type"},
		{`^\bvoid\b`, TokenVoidTypeKeyword, "void type"},

		// Equality Operator
		{`^[=!]=`, TokenEqualityOperator, "equality operator"},

		// Assignment operators
		{`^=`, TokenSimpleAssignmentOperator, "single assignment operator"},
		{`^[*/+-]=`, TokenComplexAssignmentOperator, "complex assignment operator"},

		// Math operators
		{`^[+\-]`, TokenAdditiveOperator, "additive operators (+, -)"},
		{`^[*/]`, TokenFactorOperator, "factor operators (*, /"},

		// Relational operators
		{`^[><]=?`, TokenRelationalOperator, "relational operators (>, >=, <, <=)"},

		// Logical operators
		{`^&&`, TokenLogicalAndOperator, "logical and operator"},
		{`^\|\|`, TokenLogicalOrOperator, "logical or operator"},
		{`^!`, TokenLogicalNotOperator, "logical not operator"},

		// Numbers
		{`^\d+`, TokenNumber, "number literal"},

		// Strings
		{`^"[^"]*"`, TokenString, "double quote string literal"},
		{`^'[^']*'`, TokenString, "single quote string literal"},

		// Identifiers
		{`^\w+`, TokenIdentifier, "identifiers"},
	}

	// Compile all patterns and create rules
	regexRules = make([]RegexRule, 0, len(patterns))

	for _, p := range patterns {
		compiled, err := regexp.Compile(p.pattern)
		if err != nil {
			log.Fatalf("Failed to compile regex for %s: %v", p.desc, err)
		}

		regexRules = append(regexRules, RegexRule{
			Pattern:   compiled,
			TokenType: p.tokenType,
		})
	}
}
