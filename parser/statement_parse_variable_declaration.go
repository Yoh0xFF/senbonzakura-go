package parser

import (
	"github.com/yoh0xff/senbonzakura/ast"
	"github.com/yoh0xff/senbonzakura/lexer"
)

// parseVariableDeclarationStatement parses variable declaration statements
//
// VariableDeclarationStatement
//
//	: 'let' VariableList ';'
//
// VariableList
//
//	: VariableExpression
//	| VariableList ',' VariableExpression
//	;
func parseVariableDeclarationStatement(parser *Parser, consumeStatementEnd bool) ast.Statement {
	var variables []*ast.VariableExpression

	eatToken(parser, lexer.TokenLetKeyword)
	for {
		// Parse a variable expression and append it to our list
		varExpr := parseVariableExpression(parser).(*ast.VariableExpression)
		variables = append(variables, varExpr)

		// If we don't see a comma, break the loop
		if !isNextTokenOfType(parser, lexer.TokenComma) {
			break
		}

		// Eat the comma and continue with the next variable
		eatToken(parser, lexer.TokenComma)
	}

	if consumeStatementEnd {
		eatToken(parser, lexer.TokenStatementEnd)
	}

	return &ast.VariableDeclarationStatement{
		Variables: variables,
	}
}

// parseVariableExpression parses variable expressions
//
// VariableInitializationExpression
//
//	: Identifier ':' Type ['=' Initializer]
//	;
func parseVariableExpression(parser *Parser) ast.Expression {
	identifier := parseIdentifierExpression(parser).(*ast.IdentifierExpression)

	// Require type annotation
	eatToken(parser, lexer.TokenColon)
	typeAnnotation := parseType(parser)

	// Check for an initializer
	var initializer ast.Expression
	if !isNextTokenAnyOfType(parser, []lexer.TokenType{lexer.TokenStatementEnd, lexer.TokenComma}) {
		eatToken(parser, lexer.TokenSimpleAssignmentOperator)
		initializer = parseAssignmentExpression(parser)
	}

	return &ast.VariableExpression{
		Identifier:     identifier,
		TypeAnnotation: typeAnnotation,
		Initializer:    initializer, // Will be nil if no initializer
	}
}
