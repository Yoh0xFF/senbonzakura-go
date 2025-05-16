package parser

import (
	"github.com/yoh0xff/senbonzakura/ast"
	"github.com/yoh0xff/senbonzakura/lexer"
)

// parseClassDeclarationStatement parses class declarations
//
// ClassDeclaration
//
//	: class IdentifierExpression [ClassExtendsExpression] BlockStatement
//	;
func parseClassDeclarationStatement(parser *Parser) ast.Statement {
	eatToken(parser, lexer.TokenClassKeyword)

	// Parse the class name (identifier)
	name := parseIdentifierExpression(parser).(*ast.IdentifierExpression)

	// Check for extends clause
	var superClass *ast.IdentifierExpression
	if isNextTokenOfType(parser, lexer.TokenExtendsKeyword) {
		superClass = parseClassExtendsExpression(parser).(*ast.IdentifierExpression)
	}

	// Parse the class body
	body := parseBlockStatement(parser).(*ast.BlockStatement)

	return &ast.ClassDeclarationStatement{
		Name:       name,
		SuperClass: superClass,
		Body:       body,
	}
}

// parseClassExtendsExpression parses class extension clauses
//
// ClassExtendsExpression
//
//	: extends IdentifierExpression
//	;
func parseClassExtendsExpression(parser *Parser) ast.Expression {
	eatToken(parser, lexer.TokenExtendsKeyword)
	return parseIdentifierExpression(parser)
}
