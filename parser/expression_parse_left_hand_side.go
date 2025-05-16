package parser

import (
	"github.com/yoh0xff/senbonzakura/ast"
	"github.com/yoh0xff/senbonzakura/lexer"
)

// parseLeftHandSideExpression parses left-hand-side expressions
//
// LeftHandSideExpression
//
//	: CallMemberExpression
//	;
func parseLeftHandSideExpression(parser *Parser) ast.Expression {
	return parseCallMemberExpression(parser)
}

// parseCallMemberExpression parses call and member expressions
//
// CallMemberExpression
//
//	: MemberExpression
//	| CallExpression
//	;
func parseCallMemberExpression(parser *Parser) ast.Expression {
	// Member part might be part of a call
	member := parseMemberExpression(parser)

	// See if we have a call expression
	if isNextTokenOfType(parser, lexer.TokenOpeningParenthesis) {
		return parseCallExpression(parser, member)
	}

	// Simple member expression
	return member
}

// parseCallExpression parses call expressions
//
// # Generic call expression helper
//
// CallExpression
//
//	: Callee Arguments
//	;
//
// Callee
//
//	: MemberExpression
//	| CallExpression
//	;
func parseCallExpression(parser *Parser, callee ast.Expression) ast.Expression {
	callExpression := &ast.CallExpression{
		Callee:    callee,
		Arguments: parseArguments(parser),
	}

	// Check for chained calls
	if isNextTokenOfType(parser, lexer.TokenOpeningParenthesis) {
		return parseCallExpression(parser, callExpression)
	}

	return callExpression
}

// parseArguments parses function call arguments
//
// Arguments
//
//	: '(' [ArgumentList] ')'
//	;
func parseArguments(parser *Parser) []ast.Expression {
	eatToken(parser, lexer.TokenOpeningParenthesis)

	var arguments []ast.Expression
	if !isNextTokenOfType(parser, lexer.TokenClosingParenthesis) {
		arguments = parseArgumentsList(parser)
	} else {
		arguments = []ast.Expression{}
	}

	eatToken(parser, lexer.TokenClosingParenthesis)

	return arguments
}

// parseArgumentsList parses function call argument lists
//
// ArgumentList
//
//	: AssignmentExpression
//	| ArgumentList ',' AssignmentExpression
//	;
func parseArgumentsList(parser *Parser) []ast.Expression {
	var arguments []ast.Expression

	for {
		arguments = append(arguments, parseAssignmentExpression(parser))

		if isNextTokenOfType(parser, lexer.TokenComma) {
			eatToken(parser, lexer.TokenComma)
		} else {
			break
		}
	}

	return arguments
}

// parseMemberExpression parses member expressions
//
// MemberExpression
//
//	: PrimaryExpression
//	| MemberExpression '.' Identifier
//	| MemberExpression '[' Expression ']'
//	;
func parseMemberExpression(parser *Parser) ast.Expression {
	object := parsePrimaryExpression(parser)

	for isNextTokenAnyOfType(parser, []lexer.TokenType{lexer.TokenDot, lexer.TokenOpeningBracket}) {
		if isNextTokenOfType(parser, lexer.TokenDot) {
			eatToken(parser, lexer.TokenDot)
			property := parseIdentifierExpression(parser)

			object = &ast.MemberExpression{
				Computed: false,
				Object:   object,
				Property: property,
			}
		}

		if isNextTokenOfType(parser, lexer.TokenOpeningBracket) {
			eatToken(parser, lexer.TokenOpeningBracket)
			property := ParseRootExpression(parser)
			eatToken(parser, lexer.TokenClosingBracket)

			object = &ast.MemberExpression{
				Computed: true,
				Object:   object,
				Property: property,
			}
		}
	}

	return object
}
