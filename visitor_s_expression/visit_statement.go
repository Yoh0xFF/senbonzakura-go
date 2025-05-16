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
	visitor.BeginExpr("program")

	if len(statement.Body) > 0 {
		for _, stmt := range statement.Body {
			visitor.WriteSpaceOrNewLine()
			stmt.Accept(visitor)
		}
	}

	visitor.EndExpr()
	return nil
}

func visitBlockStatement(visitor *SExpressionVisitor, statement *ast.BlockStatement) error {
	visitor.BeginExpr("block")

	if len(statement.Body) > 0 {
		for _, stmt := range statement.Body {
			visitor.WriteSpaceOrNewLine()
			stmt.Accept(visitor)
		}
	}

	visitor.EndExpr()
	return nil
}

func visitEmptyStatement(visitor *SExpressionVisitor) error {
	visitor.BeginExpr("empty")
	visitor.EndExpr()
	return nil
}

func visitExpressionStatement(visitor *SExpressionVisitor, statement *ast.ExpressionStatement) error {
	visitor.BeginExpr("expr")
	visitor.WriteSpaceOrNewLine()
	statement.Expression.Accept(visitor)
	visitor.EndExpr()
	return nil
}

func visitVariableDeclarationStatement(visitor *SExpressionVisitor, statement *ast.VariableDeclarationStatement) error {
	visitor.BeginExpr("let")

	for _, variable := range statement.Variables {
		visitor.WriteSpaceOrNewLine()
		variable.Accept(visitor)
	}

	visitor.EndExpr()
	return nil
}

func visitConditionalStatement(visitor *SExpressionVisitor, statement *ast.IfStatement) error {
	visitor.BeginExpr("if")

	// Process condition
	visitor.WriteSpaceOrNewLine()
	statement.Condition.Accept(visitor)

	// Process consequent
	visitor.WriteSpaceOrNewLine()
	statement.Consequent.Accept(visitor)

	// Process alternative if present
	if statement.Alternative != nil {
		visitor.WriteSpaceOrNewLine()
		statement.Alternative.Accept(visitor)
	}

	visitor.EndExpr()
	return nil
}

func visitWhileStatement(visitor *SExpressionVisitor, statement *ast.WhileStatement) error {
	visitor.BeginExpr("while")

	// Process condition
	visitor.WriteSpaceOrNewLine()
	statement.Condition.Accept(visitor)

	// Process body
	visitor.WriteSpaceOrNewLine()
	statement.Body.Accept(visitor)

	visitor.EndExpr()
	return nil
}

func visitDoWhileStatement(visitor *SExpressionVisitor, statement *ast.DoWhileStatement) error {
	visitor.BeginExpr("do-while")

	// Process body first (unlike while, do-while executes body first)
	visitor.WriteSpaceOrNewLine()
	statement.Body.Accept(visitor)

	// Process condition
	visitor.WriteSpaceOrNewLine()
	statement.Condition.Accept(visitor)

	visitor.EndExpr()
	return nil
}

func visitForStatement(visitor *SExpressionVisitor, statement *ast.ForStatement) error {
	visitor.BeginExpr("for")

	// Process initializer if present
	if statement.Initializer != nil {
		visitor.WriteSpaceOrNewLine()
		statement.Initializer.Accept(visitor)
	}

	// Process condition if present
	if statement.Condition != nil {
		visitor.WriteSpaceOrNewLine()
		statement.Condition.Accept(visitor)
	}

	// Process increment if present
	if statement.Increment != nil {
		visitor.WriteSpaceOrNewLine()
		statement.Increment.Accept(visitor)
	}

	// Process body
	visitor.WriteSpaceOrNewLine()
	statement.Body.Accept(visitor)

	visitor.EndExpr()
	return nil
}

func visitFunctionDeclarationStatement(visitor *SExpressionVisitor, statement *ast.FunctionDeclarationStatement) error {
	visitor.BeginExpr("def")

	// Process function name
	visitor.WriteSpaceOrNewLine()
	statement.Name.Accept(visitor)

	// Process parameters
	if len(statement.Parameters) > 0 {
		visitor.WriteSpaceOrNewLine()
		visitor.BeginExpr("params")

		for _, param := range statement.Parameters {
			visitor.WriteSpaceOrNewLine()
			visitor.BeginExpr("param")

			visitor.WriteSpaceOrNewLine()
			param.Name.Accept(visitor)

			visitor.WriteSpaceOrNewLine()
			visitor.BeginExpr("type")
			visitType(visitor, param.Type)
			visitor.EndExpr()

			visitor.EndExpr()
		}

		visitor.EndExpr()
	}

	// Process return type
	visitor.WriteSpaceOrNewLine()
	visitor.BeginExpr("return_type")
	visitType(visitor, statement.ReturnType)
	visitor.EndExpr()

	// Process function body
	visitor.WriteSpaceOrNewLine()
	statement.Body.Accept(visitor)

	visitor.EndExpr()
	return nil
}

func visitReturnStatement(visitor *SExpressionVisitor, statement *ast.ReturnStatement) error {
	visitor.BeginExpr("return")

	// Process return argument if present
	if statement.Argument != nil {
		visitor.WriteSpaceOrNewLine()
		statement.Argument.Accept(visitor)
	}

	visitor.EndExpr()
	return nil
}

func visitClassDeclarationStatement(visitor *SExpressionVisitor, statement *ast.ClassDeclarationStatement) error {
	visitor.BeginExpr("class")

	// Process class name
	visitor.WriteSpaceOrNewLine()
	statement.Name.Accept(visitor)

	// Process superclass if present
	if statement.SuperClass != nil {
		visitor.WriteSpaceOrNewLine()
		visitor.BeginExpr("extends")
		visitor.WriteSpaceOrNewLine()
		statement.SuperClass.Accept(visitor)
		visitor.EndExpr()
	}

	// Process class body
	visitor.WriteSpaceOrNewLine()
	statement.Body.Accept(visitor)

	visitor.EndExpr()
	return nil
}

// Helper function to visit type annotations
func visitType(visitor *SExpressionVisitor, typeAnnotation ast.Type) {
	switch t := typeAnnotation.(type) {
	case *ast.PrimitiveType:
		visitor.WriteString(t.String())
	case *ast.ArrayType:
		visitor.BeginExpr("array")
		visitor.WriteSpaceOrNewLine()
		visitType(visitor, t.ElementType)
		visitor.EndExpr()
	case *ast.FunctionType:
		visitor.BeginExpr("function")

		// Parameters
		visitor.WriteSpaceOrNewLine()
		visitor.BeginExpr("params")
		for _, param := range t.Params {
			visitor.WriteSpaceOrNewLine()
			visitType(visitor, param)
		}
		visitor.EndExpr()

		// Return type
		visitor.WriteSpaceOrNewLine()
		visitType(visitor, t.ReturnType)
		visitor.EndExpr()
	case *ast.ClassType:
		visitor.BeginExpr("class-type")
		visitor.WriteSpaceOrNewLine()
		visitor.WriteString(t.Name)
		if t.SuperClass != nil {
			visitor.WriteSpaceOrNewLine()
			visitor.BeginExpr("extends")
			visitor.WriteSpaceOrNewLine()
			visitor.WriteString(*t.SuperClass)
			visitor.EndExpr()
		}
		visitor.EndExpr()
	case *ast.GenericType:
		visitor.BeginExpr("generic")
		visitor.WriteSpaceOrNewLine()
		visitor.WriteString(t.Base)
		visitor.WriteSpaceOrNewLine()
		visitor.BeginExpr("args")
		for _, arg := range t.TypeArgs {
			visitor.WriteSpaceOrNewLine()
			visitType(visitor, arg)
		}
		visitor.EndExpr()
		visitor.EndExpr()
	case *ast.VoidType:
		visitor.WriteString("void")
	default:
		panic(fmt.Errorf("unknown type annotation: %T", typeAnnotation))
	}
}
