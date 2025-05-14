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
		// Whitespace
		{`^\s+`, Whitespace, "whitespace"},

		// Comments
		{`^//.*`, SingleLineComment, "single line comments"},
		{`^/\*[\s\S]*?\*/`, MultiLineComment, "multi line comments"},

		// Symbols, delimiters
		{`^;`, StatementEnd, "statement end (;) symbol"},
		{`^\{`, OpeningBrace, "opening brace ({) symbol"},
		{`^}`, ClosingBrace, "closing brace (}) symbol"},
		{`^\(`, OpeningParenthesis, "opening parenthesis (() symbol"},
		{`^\)`, ClosingParenthesis, "closing parenthesis ()) symbol"},
		{`^\[`, OpeningBracket, "opening bracket ([) symbol"},
		{`^]`, ClosingBracket, "closing bracket (]) symbol"},
		{`^,`, Comma, "comma (,) symbol"},
		{`^\.`, Dot, "dot (.) symbol"},
		{`^:`, Colon, "colon (:) symbol"},

		// Keywords
		{`^\btrue\b`, Boolean, "the 'true' keyword"},
		{`^\bfalse\b`, Boolean, "the 'false' keyword"},
		{`^\bnil\b`, Nil, "the 'nil' keyword"},
		{`^\blet\b`, LetKeyword, "the 'let' keyword"},
		{`^\bif\b`, IfKeyword, "the 'if' keyword"},
		{`^\belse\b`, ElseKeyword, "the 'else' keyword"},
		{`^\bwhile\b`, WhileKeyword, "the 'while' keyword"},
		{`^\bdo\b`, DoKeyword, "the 'do' keyword"},
		{`^\bfor\b`, ForKeyword, "the 'for' keyword"},
		{`^\bdef\b`, DefKeyword, "the 'def' keyword"},
		{`^\breturn\b`, ReturnKeyword, "the 'return' keyword"},
		{`^\bclass\b`, ClassKeyword, "the 'class' keyword"},
		{`^\bextends\b`, ExtendsKeyword, "the 'extends' keyword"},
		{`^\bthis\b`, ThisKeyword, "the 'this' keyword"},
		{`^\bsuper\b`, SuperKeyword, "the 'super' keyword"},
		{`^\bnew\b`, NewKeyword, "the 'new' keyword"},
		{`^\btype\b`, TypeKeyword, "type keyword"},
		{`^\bnumber\b`, NumberTypeKeyword, "number type"},
		{`^\bstring\b`, StringTypeKeyword, "string type"},
		{`^\bboolean\b`, BooleanTypeKeyword, "boolean type"},
		{`^\bvoid\b`, VoidTypeKeyword, "void type"},

		// Equality Operator
		{`^[=!]=`, EqualityOperator, "equality operator"},

		// Assignment operators
		{`^=`, SimpleAssignmentOperator, "single assignment operator"},
		{`^[*/+-]=`, ComplexAssignmentOperator, "complex assignment operator"},

		// Math operators
		{`^[+\-]`, AdditiveOperator, "additive operators (+, -)"},
		{`^[*/]`, FactorOperator, "factor operators (*, /"},

		// Relational operators
		{`^[><]=?`, RelationalOperator, "relational operators (>, >=, <, <=)"},

		// Logical operators
		{`^&&`, LogicalAndOperator, "logical and operator"},
		{`^\|\|`, LogicalOrOperator, "logical or operator"},
		{`^!`, LogicalNotOperator, "logical not operator"},

		// Numbers
		{`^\d+`, Number, "number literal"},

		// Strings
		{`^"[^"]*"`, String, "double quote string literal"},
		{`^'[^']*'`, String, "single quote string literal"},

		// Identifiers
		{`^\w+`, Identifier, "identifiers"},
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
