package parser

import (
	"github.com/yoh0xff/senbonzakura/ast"
	"github.com/yoh0xff/senbonzakura/lexer"
)

// parseWhileStatement parses while loop statements
//
// WhileStatement
//
//	: while '(' Expression ')' Statement ';'
//	;
func parseWhileStatement(parser *Parser) ast.Statement {
	eatToken(parser, lexer.TokenWhileKeyword)

	eatToken(parser, lexer.TokenOpeningParenthesis)
	condition := ParseRootExpression(parser)
	eatToken(parser, lexer.TokenClosingParenthesis)

	bodyStmt := parseStatement(parser)

	// Convert the body to a block statement if it isn't already one
	body, ok := bodyStmt.(*ast.BlockStatement)
	if !ok {
		body = &ast.BlockStatement{
			Body: []ast.Statement{bodyStmt},
		}
	}

	return &ast.WhileStatement{
		Condition: condition,
		Body:      body,
	}
}

// parseDoWhileStatement parses do-while loop statements
//
// DoWhileStatement
//
//	: do Statement while '(' Expression ')' ';'
//	;
func parseDoWhileStatement(parser *Parser) ast.Statement {
	eatToken(parser, lexer.TokenDoKeyword)

	bodyStmt := parseStatement(parser)

	// Convert the body to a block statement if it isn't already one
	body, ok := bodyStmt.(*ast.BlockStatement)
	if !ok {
		body = &ast.BlockStatement{
			Body: []ast.Statement{bodyStmt},
		}
	}

	eatToken(parser, lexer.TokenWhileKeyword)

	eatToken(parser, lexer.TokenOpeningParenthesis)
	condition := ParseRootExpression(parser)
	eatToken(parser, lexer.TokenClosingParenthesis)

	eatToken(parser, lexer.TokenStatementEnd)

	return &ast.DoWhileStatement{
		Body:      body,
		Condition: condition,
	}
}

// parseForStatement parses for loop statements
//
// ForStatement
//
//	: for '(' [InitExpression] ';' [Expression] ';' [Expression] ')' Statement
//	;
func parseForStatement(parser *Parser) ast.Statement {
	eatToken(parser, lexer.TokenForKeyword)
	eatToken(parser, lexer.TokenOpeningParenthesis)

	var initializer ast.Statement
	if !isNextTokenOfType(parser, lexer.TokenStatementEnd) {
		initializer = parseForStatementInitStatement(parser)
	}
	eatToken(parser, lexer.TokenStatementEnd)

	var condition ast.Expression
	if !isNextTokenOfType(parser, lexer.TokenStatementEnd) {
		condition = ParseRootExpression(parser)
	}
	eatToken(parser, lexer.TokenStatementEnd)

	var increment ast.Expression
	if !isNextTokenOfType(parser, lexer.TokenClosingParenthesis) {
		increment = ParseRootExpression(parser)
	}
	eatToken(parser, lexer.TokenClosingParenthesis)

	bodyStmt := parseStatement(parser)

	// Convert the body to a block statement if it isn't already one
	body, ok := bodyStmt.(*ast.BlockStatement)
	if !ok {
		body = &ast.BlockStatement{
			Body: []ast.Statement{bodyStmt},
		}
	}

	return &ast.ForStatement{
		Initializer: initializer,
		Condition:   condition,
		Increment:   increment,
		Body:        body,
	}
}

// parseForStatementInitStatement parses initializer statements in for loops
//
// ForStatementInit
//
//	: VariableDeclarationStatement
//	| Expression
//	;
func parseForStatementInitStatement(parser *Parser) ast.Statement {
	if isNextTokenOfType(parser, lexer.TokenLetKeyword) {
		return parseVariableDeclarationStatement(parser, false)
	}
	return parseExpressionStatement(parser, false)
}
