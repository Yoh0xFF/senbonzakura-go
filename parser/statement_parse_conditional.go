package parser

import (
	"github.com/yoh0xff/senbonzakura/ast"
	"github.com/yoh0xff/senbonzakura/lexer"
)

// parseIfStatement parses conditional statements
//
// ConditionalStatement
//
//	: if '(' Expression ')' Statement [else Statement]
//	;
func parseIfStatement(parser *Parser) ast.Statement {
	eatToken(parser, lexer.TokenIfKeyword)

	// Parse the condition
	eatToken(parser, lexer.TokenOpeningParenthesis)
	condition := ParseRootExpression(parser)
	eatToken(parser, lexer.TokenClosingParenthesis)

	// Parse the consequent (then block)
	consequent := parseStatement(parser)
	consequentBlock, ok := consequent.(*ast.BlockStatement)
	if !ok {
		// If the consequent is not a block statement, wrap it in one
		consequentBlock = &ast.BlockStatement{
			Body: []ast.Statement{consequent},
		}
	}

	// Check for an else clause
	var alternativeBlock *ast.BlockStatement
	if isNextTokenOfType(parser, lexer.TokenElseKeyword) {
		eatToken(parser, lexer.TokenElseKeyword)
		alternative := parseStatement(parser)

		// If the alternative is already a block statement, use it directly
		altBlock, ok := alternative.(*ast.BlockStatement)
		if ok {
			alternativeBlock = altBlock
		} else {
			// Otherwise, wrap it in a block statement
			alternativeBlock = &ast.BlockStatement{
				Body: []ast.Statement{alternative},
			}
		}
	}

	return &ast.IfStatement{
		Condition:   condition,
		Consequent:  consequentBlock,
		Alternative: alternativeBlock, // will be nil if there's no else clause
	}
}
