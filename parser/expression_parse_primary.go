package parser

import (
	"github.com/yoh0xff/senbonzakura/ast"
	"github.com/yoh0xff/senbonzakura/lexer"
)

// parsePrimaryExpression parses primary expressions
//
// PrimaryExpression
//
//	: LiteralExpression
//	| GroupExpression
//	| IdentifierExpression
//	| ThisExpression
//	;
func parsePrimaryExpression(parser *Parser) ast.Expression {
	if isNextTokenLiteral(parser) {
		return parseLiteralExpression(parser)
	}

	switch parser.lookahead.TokenType {
	case lexer.TokenOpeningParenthesis:
		return parseGroupExpression(parser)
	case lexer.TokenIdentifier:
		return parseIdentifierExpression(parser)
	case lexer.TokenThisKeyword:
		return parseThisExpression(parser)
	case lexer.TokenSuperKeyword:
		return parseSuperExpression(parser)
	case lexer.TokenNewKeyword:
		return parseNewExpression(parser)
	default:
		return parseLeftHandSideExpression(parser)
	}
}

// parseGroupExpression parses parenthesized expressions
//
// GroupExpression
//
//	: '(' Expression ')'
//	;
func parseGroupExpression(parser *Parser) ast.Expression {
	eatToken(parser, lexer.TokenOpeningParenthesis)
	expression := ParseRootExpression(parser)
	eatToken(parser, lexer.TokenClosingParenthesis)

	return expression
}

// parseIdentifierExpression parses identifiers
//
// IdentifierExpression
//
//	: IDENTIFIER
//	;
func parseIdentifierExpression(parser *Parser) ast.Expression {
	identifierToken := eatToken(parser, lexer.TokenIdentifier)
	identifierValue := parser.source[identifierToken.Start:identifierToken.End]

	return &ast.IdentifierExpression{
		Name: identifierValue,
	}
}

// parseThisExpression parses 'this' expressions
//
// ThisExpression
//
//	: this
//	;
func parseThisExpression(parser *Parser) ast.Expression {
	eatToken(parser, lexer.TokenThisKeyword)
	return &ast.ThisExpression{}
}

// parseSuperExpression parses 'super' expressions
//
// SuperExpression
//
//	: super
//	;
func parseSuperExpression(parser *Parser) ast.Expression {
	eatToken(parser, lexer.TokenSuperKeyword)
	return &ast.SuperExpression{}
}

// parseNewExpression parses 'new' expressions
//
// NewExpression
//
//	: new MemberExpression Arguments
//	;
func parseNewExpression(parser *Parser) ast.Expression {
	eatToken(parser, lexer.TokenNewKeyword)

	callee := parseMemberExpression(parser)
	arguments := parseArguments(parser)

	return &ast.NewExpression{
		Callee:    callee,
		Arguments: arguments,
	}
}
