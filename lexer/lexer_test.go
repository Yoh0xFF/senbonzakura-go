package lexer

import (
	"strings"
	"testing"
)

// Helper function to find substring index
func findIndex(source, substr string) int {
	idx := strings.Index(source, substr)
	if idx == -1 {
		return 0
	}
	return idx
}

// Test comments
func TestSingleLineComment(t *testing.T) {
	source := `
        // This is single line comment
        17
    `

	lexer := NewLexer(source)
	nextToken := lexer.NextToken()

	expected := Token{
		TokenType: Number,
		Start:     findIndex(source, "1"),
		End:       findIndex(source, "7") + 1,
	}

	if !nextToken.Equal(expected) {
		t.Errorf("Expected token %v, got %v", expected, nextToken)
	}
}

func TestMultiLineComment(t *testing.T) {
	source := `
        /*
        This is multi line comment,
        and we should skip it
        */
        1719
    `

	lexer := NewLexer(source)
	nextToken := lexer.NextToken()

	expected := Token{
		TokenType: Number,
		Start:     findIndex(source, "1"),
		End:       findIndex(source, "9") + 1,
	}

	if !nextToken.Equal(expected) {
		t.Errorf("Expected token %v, got %v", expected, nextToken)
	}
}

func TestFormattedMultiLineComment(t *testing.T) {
	source := `
        /**
        * This is multi line comment,
        * and we should skip it
        */
        1719
    `

	lexer := NewLexer(source)
	nextToken := lexer.NextToken()

	expected := Token{
		TokenType: Number,
		Start:     findIndex(source, "1"),
		End:       findIndex(source, "9") + 1,
	}

	if !nextToken.Equal(expected) {
		t.Errorf("Expected token %v, got %v", expected, nextToken)
	}
}

// Test numbers
func TestNumberToken(t *testing.T) {
	source := "12"
	lexer := NewLexer(source)
	nextToken := lexer.NextToken()

	expected := Token{
		TokenType: Number,
		Start:     findIndex(source, "1"),
		End:       findIndex(source, "2") + 1,
	}

	if !nextToken.Equal(expected) {
		t.Errorf("Expected token %v, got %v", expected, nextToken)
	}
}

func TestSkipWhitespace(t *testing.T) {
	source := "    12"
	lexer := NewLexer(source)
	nextToken := lexer.NextToken()

	expected := Token{
		TokenType: Number,
		Start:     findIndex(source, "1"),
		End:       findIndex(source, "2") + 1,
	}

	if !nextToken.Equal(expected) {
		t.Errorf("Expected token %v, got %v", expected, nextToken)
	}
}

// Test strings
func TestStringTokens(t *testing.T) {
	source := `  "Hello" 'world'  `
	tokenA := `"Hello"`
	tokenB := `'world'`
	lexer := NewLexer(source)

	nextTokenA := lexer.NextToken()
	nextTokenB := lexer.NextToken()

	expectedA := Token{
		TokenType: String,
		Start:     strings.Index(source, tokenA),
		End:       strings.Index(source, tokenA) + len(tokenA),
	}

	expectedB := Token{
		TokenType: String,
		Start:     strings.Index(source, tokenB),
		End:       strings.Index(source, tokenB) + len(tokenB),
	}

	if !nextTokenA.Equal(expectedA) {
		t.Errorf("Expected token %v, got %v", expectedA, nextTokenA)
	}

	if !nextTokenB.Equal(expectedB) {
		t.Errorf("Expected token %v, got %v", expectedB, nextTokenB)
	}
}

func TestStringTokenWithWhitespace(t *testing.T) {
	source := `  " Hello "  `
	token := `" Hello "`
	lexer := NewLexer(source)

	nextToken := lexer.NextToken()

	expected := Token{
		TokenType: String,
		Start:     strings.Index(source, token),
		End:       strings.Index(source, token) + len(token),
	}

	if !nextToken.Equal(expected) {
		t.Errorf("Expected token %v, got %v", expected, nextToken)
	}
}

// Test helper functions for comprehensive testing
func TestLexerEndOfInput(t *testing.T) {
	source := "42"
	lexer := NewLexer(source)

	// Get the number token
	_ = lexer.NextToken()

	// Next should be End token
	endToken := lexer.NextToken()

	if endToken.TokenType != End {
		t.Errorf("Expected End token, got %v", endToken.TokenType)
	}
}

func TestLexerMultipleTokens(t *testing.T) {
	source := `let x = 42;`
	lexer := NewLexer(source)

	expectedTokens := []TokenType{
		LetKeyword,
		Identifier,
		SimpleAssignmentOperator,
		Number,
		StatementEnd,
		End,
	}

	for i, expectedType := range expectedTokens {
		token := lexer.NextToken()
		if token.TokenType != expectedType {
			t.Errorf("Token %d: expected %v, got %v", i, expectedType, token.TokenType)
		}
	}
}

func TestLexerInvalidToken(t *testing.T) {
	source := "@invalid"
	lexer := NewLexer(source)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for invalid token, but got none")
		}
	}()

	lexer.NextToken()
}

// Benchmark tests
func BenchmarkLexerSimple(b *testing.B) {
	source := "let x = 42;"

	for i := 0; i < b.N; i++ {
		lexer := NewLexer(source)
		for token := lexer.NextToken(); token.TokenType != End; token = lexer.NextToken() {
			// Process all tokens
		}
	}
}

func BenchmarkLexerComplex(b *testing.B) {
	source := `
		// Complex test
		class Person {
			def constructor(name: string, age: number) {
				this.name = name;
				this.age = age;
			}
			
			def getName(): string {
				return this.name;
			}
		}
		
		let person = new Person("John", 30);
	`

	for i := 0; i < b.N; i++ {
		lexer := NewLexer(source)
		for token := lexer.NextToken(); token.TokenType != End; token = lexer.NextToken() {
			// Process all tokens
		}
	}
}
