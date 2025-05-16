package parser

import (
	"fmt"
	"github.com/yoh0xff/senbonzakura/ast"
	"github.com/yoh0xff/senbonzakura/lexer"
)

// parseType parses type annotations
func parseType(parser *Parser) ast.Type {
	switch parser.lookahead.TokenType {
	case lexer.TokenNumberTypeKeyword:
		eatToken(parser, lexer.TokenNumberTypeKeyword)
		return &ast.PrimitiveType{Kind: ast.NumberType}

	case lexer.TokenStringTypeKeyword:
		eatToken(parser, lexer.TokenStringTypeKeyword)
		return &ast.PrimitiveType{Kind: ast.StringType}

	case lexer.TokenBooleanTypeKeyword:
		eatToken(parser, lexer.TokenBooleanTypeKeyword)
		return &ast.PrimitiveType{Kind: ast.BooleanType}

	case lexer.TokenVoidTypeKeyword:
		eatToken(parser, lexer.TokenVoidTypeKeyword)
		return &ast.VoidType{}

	case lexer.TokenIdentifier:
		// Handle class types or custom types
		identifierToken := eatToken(parser, lexer.TokenIdentifier)
		typeName := parser.source[identifierToken.Start:identifierToken.End]

		// Check for generic type parameters
		if isNextTokenOfType(parser, lexer.TokenOpeningBracket) {
			eatToken(parser, lexer.TokenOpeningBracket)
			typeArgs := []ast.Type{}

			for {
				typeArgs = append(typeArgs, parseType(parser))

				if !isNextTokenOfType(parser, lexer.TokenComma) {
					break
				}

				eatToken(parser, lexer.TokenComma)
			}

			eatToken(parser, lexer.TokenClosingBracket)

			return &ast.GenericType{
				Base:     typeName,
				TypeArgs: typeArgs,
			}
		} else {
			// For class types
			return &ast.ClassType{
				Name:       typeName,
				SuperClass: nil, // No super class by default
			}
		}

	case lexer.TokenOpeningBracket:
		// Handle array types
		eatToken(parser, lexer.TokenOpeningBracket)
		elementType := parseType(parser)
		eatToken(parser, lexer.TokenClosingBracket)

		return &ast.ArrayType{
			ElementType: elementType,
		}

	default:
		panic(fmt.Sprintf(
			"Expected type annotation, found: %s",
			parser.lookahead.TokenType.String(),
		))
	}
}
