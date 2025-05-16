package visitor_s_expression

import (
	"fmt"
	"github.com/yoh0xff/senbonzakura/ast"
)

func visitStatement(visitor *SExpressionVisitor, statement ast.Statement) any {
	switch statement.NodeType() {
	case ast.NodeProgramStatement:
		return visitProgramStatement(visitor, statement.(*ast.ProgramStatement))
	case ast.NodeBlockStatement:
		return visitBlockStatement(visitor, statement.(*ast.BlockStatement))
	case ast.NodeEmptyStatement:
		return visitEmptyStatement(visitor)
	case ast.NodeExpressionStatement:
		return visitExpressionStatement(visitor, statement.(*ast.ExpressionStatement))
	case ast.NodeVariableDeclarationStatement:
		return visitVariableDeclarationStatement(visitor, statement.(*ast.VariableDeclarationStatement))
	case ast.NodeIfStatement:
		return visitConditionalStatement(visitor, statement.(*ast.IfStatement))
	case ast.NodeWhileStatement:
		return visitWhileStatement(visitor, statement.(*ast.WhileStatement))
	case ast.NodeDoWhileStatement:
		return visitDoWhileStatement(visitor, statement.(*ast.DoWhileStatement))
	case ast.NodeForStatement:
		return visitForStatement(visitor, statement.(*ast.ForStatement))
	case ast.NodeFunctionDeclarationStatement:
		return visitFunctionDeclarationStatement(visitor, statement.(*ast.FunctionDeclarationStatement))
	case ast.NodeReturnStatement:
		return visitReturnStatement(visitor, statement.(*ast.ReturnStatement))
	case ast.NodeClassDeclarationStatement:
		return visitClassDeclarationStatement(visitor, statement.(*ast.ClassDeclarationStatement))
	default:
		return fmt.Errorf("unknown statement type: %T", statement)
	}
}

func visitProgramStatement(visitor *SExpressionVisitor, statement *ast.ProgramStatement) error {
	visitor.beginExpression("program")

	if len(statement.Body) > 0 {
		for _, stmt := range statement.Body {
			visitor.writeSpaceOrNewLine()
			stmt.Accept(visitor)
		}
	}

	visitor.endExpression()
	return nil
}

func visitBlockStatement(visitor *SExpressionVisitor, statement *ast.BlockStatement) error {
	visitor.beginExpression("block")

	if len(statement.Body) > 0 {
		for _, stmt := range statement.Body {
			visitor.writeSpaceOrNewLine()
			stmt.Accept(visitor)
		}
	}

	visitor.endExpression()
	return nil
}

func visitEmptyStatement(visitor *SExpressionVisitor) error {
	visitor.beginExpression("empty")
	visitor.endExpression()
	return nil
}

func visitExpressionStatement(visitor *SExpressionVisitor, statement *ast.ExpressionStatement) error {
	visitor.beginExpression("expr")
	visitor.writeSpaceOrNewLine()
	statement.Expression.Accept(visitor)
	visitor.endExpression()
	return nil
}

func visitVariableDeclarationStatement(visitor *SExpressionVisitor, statement *ast.VariableDeclarationStatement) error {
	visitor.beginExpression("let")

	for _, variable := range statement.Variables {
		visitor.writeSpaceOrNewLine()
		variable.Accept(visitor)
	}

	visitor.endExpression()
	return nil
}

func visitConditionalStatement(visitor *SExpressionVisitor, statement *ast.IfStatement) error {
	visitor.beginExpression("if")

	// Process condition
	visitor.writeSpaceOrNewLine()
	statement.Condition.Accept(visitor)

	// Process consequent
	visitor.writeSpaceOrNewLine()
	statement.Consequent.Accept(visitor)

	// Process alternative if present
	if statement.Alternative != nil {
		visitor.writeSpaceOrNewLine()
		statement.Alternative.Accept(visitor)
	}

	visitor.endExpression()
	return nil
}

func visitWhileStatement(visitor *SExpressionVisitor, statement *ast.WhileStatement) error {
	visitor.beginExpression("while")

	// Process condition
	visitor.writeSpaceOrNewLine()
	statement.Condition.Accept(visitor)

	// Process body
	visitor.writeSpaceOrNewLine()
	statement.Body.Accept(visitor)

	visitor.endExpression()
	return nil
}

func visitDoWhileStatement(visitor *SExpressionVisitor, statement *ast.DoWhileStatement) error {
	visitor.beginExpression("do-while")

	// Process body first (unlike while, do-while executes body first)
	visitor.writeSpaceOrNewLine()
	statement.Body.Accept(visitor)

	// Process condition
	visitor.writeSpaceOrNewLine()
	statement.Condition.Accept(visitor)

	visitor.endExpression()
	return nil
}

func visitForStatement(visitor *SExpressionVisitor, statement *ast.ForStatement) error {
	visitor.beginExpression("for")

	// Process initializer if present
	if statement.Initializer != nil {
		visitor.writeSpaceOrNewLine()
		statement.Initializer.Accept(visitor)
	}

	// Process condition if present
	if statement.Condition != nil {
		visitor.writeSpaceOrNewLine()
		statement.Condition.Accept(visitor)
	}

	// Process increment if present
	if statement.Increment != nil {
		visitor.writeSpaceOrNewLine()
		statement.Increment.Accept(visitor)
	}

	// Process body
	visitor.writeSpaceOrNewLine()
	statement.Body.Accept(visitor)

	visitor.endExpression()
	return nil
}

func visitFunctionDeclarationStatement(visitor *SExpressionVisitor, statement *ast.FunctionDeclarationStatement) error {
	visitor.beginExpression("def")

	// Process function name
	visitor.writeSpaceOrNewLine()
	statement.Name.Accept(visitor)

	// Process parameters
	if len(statement.Parameters) > 0 {
		visitor.writeSpaceOrNewLine()
		visitor.beginExpression("params")

		for _, param := range statement.Parameters {
			visitor.writeSpaceOrNewLine()
			visitor.beginExpression("param")

			visitor.writeSpaceOrNewLine()
			param.Name.Accept(visitor)

			visitor.writeSpaceOrNewLine()
			visitor.beginExpression("type")
			visitType(visitor, param.Type)
			visitor.endExpression()

			visitor.endExpression()
		}

		visitor.endExpression()
	}

	// Process return type
	visitor.writeSpaceOrNewLine()
	visitor.beginExpression("return_type")
	visitType(visitor, statement.ReturnType)
	visitor.endExpression()

	// Process function body
	visitor.writeSpaceOrNewLine()
	statement.Body.Accept(visitor)

	visitor.endExpression()
	return nil
}

func visitReturnStatement(visitor *SExpressionVisitor, statement *ast.ReturnStatement) error {
	visitor.beginExpression("return")

	// Process return argument if present
	if statement.Argument != nil {
		visitor.writeSpaceOrNewLine()
		statement.Argument.Accept(visitor)
	}

	visitor.endExpression()
	return nil
}

func visitClassDeclarationStatement(visitor *SExpressionVisitor, statement *ast.ClassDeclarationStatement) error {
	visitor.beginExpression("class")

	// Process class name
	visitor.writeSpaceOrNewLine()
	statement.Name.Accept(visitor)

	// Process superclass if present
	if statement.SuperClass != nil {
		visitor.writeSpaceOrNewLine()
		visitor.beginExpression("extends")
		visitor.writeSpaceOrNewLine()
		statement.SuperClass.Accept(visitor)
		visitor.endExpression()
	}

	// Process class body
	visitor.writeSpaceOrNewLine()
	statement.Body.Accept(visitor)

	visitor.endExpression()
	return nil
}

// Helper function to visit type annotations
func visitType(visitor *SExpressionVisitor, typeAnnotation ast.Type) {
	switch t := typeAnnotation.(type) {
	case *ast.PrimitiveType:
		visitor.writeString(t.String())
	case *ast.ArrayType:
		visitor.beginExpression("array")
		visitor.writeSpaceOrNewLine()
		visitType(visitor, t.ElementType)
		visitor.endExpression()
	case *ast.FunctionType:
		visitor.beginExpression("function")

		// Parameters
		visitor.writeSpaceOrNewLine()
		visitor.beginExpression("params")
		for _, param := range t.Params {
			visitor.writeSpaceOrNewLine()
			visitType(visitor, param)
		}
		visitor.endExpression()

		// Return type
		visitor.writeSpaceOrNewLine()
		visitType(visitor, t.ReturnType)
		visitor.endExpression()
	case *ast.ClassType:
		visitor.beginExpression("class-type")
		visitor.writeSpaceOrNewLine()
		visitor.writeString(t.Name)
		if t.SuperClass != nil {
			visitor.writeSpaceOrNewLine()
			visitor.beginExpression("extends")
			visitor.writeSpaceOrNewLine()
			visitor.writeString(*t.SuperClass)
			visitor.endExpression()
		}
		visitor.endExpression()
	case *ast.GenericType:
		visitor.beginExpression("generic")
		visitor.writeSpaceOrNewLine()
		visitor.writeString(t.Base)
		visitor.writeSpaceOrNewLine()
		visitor.beginExpression("args")
		for _, arg := range t.TypeArgs {
			visitor.writeSpaceOrNewLine()
			visitType(visitor, arg)
		}
		visitor.endExpression()
		visitor.endExpression()
	case *ast.VoidType:
		visitor.writeString("void")
	default:
		panic(fmt.Errorf("unknown type annotation: %T", typeAnnotation))
	}
}
