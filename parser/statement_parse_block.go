package parser

import (
	"github.com/yoh0xff/senbonzakura/ast"
	"github.com/yoh0xff/senbonzakura/lexer"
)

// parseBlockStatement parses block statements
//
// BlockStatement
//
//	: '{' OptStatementList '}'
//	;
func parseBlockStatement(parser *Parser) ast.Statement {
	eatToken(parser, lexer.TokenOpeningBrace)

	var block []ast.Statement
	if !isNextTokenOfType(parser, lexer.TokenClosingBrace) {
		block = parseStatementList(parser, lexer.TokenClosingBrace)
	} else {
		block = []ast.Statement{}
	}

	eatToken(parser, lexer.TokenClosingBrace)

	return &ast.BlockStatement{
		Body: block,
	}
}

// parseStatementList parses a list of statements
//
// StatementList
//
//	: Statement
//	| StatementList Statement
//	;
func parseStatementList(parser *Parser, stopTokenType lexer.TokenType) []ast.Statement {
	var statementList []ast.Statement

	for !isNextTokenOfType(parser, lexer.TokenEnd) &&
		(stopTokenType == -1 || !isNextTokenOfType(parser, stopTokenType)) {
		statement := parseStatement(parser)
		statementList = append(statementList, statement)
	}

	return statementList
}

// parseStatement parses a statement
//
// Statement
//
//	: ExpressionStatement
//	| BlockStatement
//	| EmptyStatement
//	| VariableStatement
//	| ConditionalStatement
//	| IterationStatement
//	| FunctionDeclarationStatement
//	| ReturnStatement
//	| ClassDeclaration
//	;
func parseStatement(parser *Parser) ast.Statement {
	switch parser.lookahead.TokenType {
	case lexer.TokenStatementEnd:
		return parseEmptyStatement(parser)
	case lexer.TokenOpeningBrace:
		return parseBlockStatement(parser)
	case lexer.TokenLetKeyword:
		return parseVariableDeclarationStatement(parser, true)
	case lexer.TokenIfKeyword:
		return parseIfStatement(parser)
	case lexer.TokenWhileKeyword:
		return parseWhileStatement(parser)
	case lexer.TokenDoKeyword:
		return parseDoWhileStatement(parser)
	case lexer.TokenForKeyword:
		return parseForStatement(parser)
	case lexer.TokenDefKeyword:
		return parseFunctionDeclarationStatement(parser)
	case lexer.TokenReturnKeyword:
		return parseReturnStatement(parser)
	case lexer.TokenClassKeyword:
		return parseClassDeclarationStatement(parser)
	default:
		return parseExpressionStatement(parser, true)
	}
}
