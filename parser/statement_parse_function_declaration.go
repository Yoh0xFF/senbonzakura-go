package parser

import (
	"github.com/yoh0xff/senbonzakura/ast"
	"github.com/yoh0xff/senbonzakura/lexer"
)

// parseFunctionDeclarationStatement parses function declarations
//
// FunctionDeclaration
//
//	: def '(' [FormalParameterList] ')' [':' Type] BlockStatement
//	;
func parseFunctionDeclarationStatement(parser *Parser) ast.Statement {
	eatToken(parser, lexer.TokenDefKeyword)
	name := parseIdentifierExpression(parser).(*ast.IdentifierExpression)

	eatToken(parser, lexer.TokenOpeningParenthesis)
	var parameters []ast.Parameter
	if !isNextTokenOfType(parser, lexer.TokenClosingParenthesis) {
		parameters = parseFormalParameterListExpression(parser)
	} else {
		parameters = []ast.Parameter{}
	}
	eatToken(parser, lexer.TokenClosingParenthesis)

	// Parse return type
	var returnType ast.Type
	if isNextTokenOfType(parser, lexer.TokenColon) {
		eatToken(parser, lexer.TokenColon)
		returnType = parseType(parser)
	} else {
		returnType = &ast.VoidType{}
	}

	body := parseBlockStatement(parser).(*ast.BlockStatement)

	return &ast.FunctionDeclarationStatement{
		Name:       name,
		Parameters: parameters,
		ReturnType: returnType,
		Body:       body,
	}
}

// parseFormalParameterListExpression parses function parameter lists
//
// FormalParameterList
//
//	: IdentifierExpression ':' Type
//	| FormalParameterList ',' IdentifierExpression ':' Type
//	;
func parseFormalParameterListExpression(parser *Parser) []ast.Parameter {
	parameters := []ast.Parameter{}

	// Parse first parameter
	paramName := parseIdentifierExpression(parser)
	eatToken(parser, lexer.TokenColon)
	paramType := parseType(parser)
	parameters = append(parameters, ast.Parameter{
		Name: paramName,
		Type: paramType,
	})

	// Parse additional parameters if any
	for isNextTokenOfType(parser, lexer.TokenComma) {
		eatToken(parser, lexer.TokenComma)
		paramName := parseIdentifierExpression(parser)
		eatToken(parser, lexer.TokenColon)
		paramType := parseType(parser)
		parameters = append(parameters, ast.Parameter{
			Name: paramName,
			Type: paramType,
		})
	}

	return parameters
}

// parseReturnStatement parses return statements
//
// ReturnStatement
//
//	: return [Expression] ';'
//	;
func parseReturnStatement(parser *Parser) ast.Statement {
	eatToken(parser, lexer.TokenReturnKeyword)
	var argument ast.Expression
	if !isNextTokenOfType(parser, lexer.TokenStatementEnd) {
		argument = ParseRootExpression(parser)
	}
	eatToken(parser, lexer.TokenStatementEnd)

	return &ast.ReturnStatement{
		Argument: argument,
	}
}
