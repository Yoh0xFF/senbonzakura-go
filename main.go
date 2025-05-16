package main

import (
	"fmt"
	"github.com/yoh0xff/senbonzakura/parser"
	"github.com/yoh0xff/senbonzakura/visitor_s_expression"
)

func main() {
	// A simple program in our language
	source := `
		// Sample program that shows basic language features
		class Person {
			def constructor(name: string, age: number) {
				this.name = name;
				this.age = age;
			}
			
			def getName(): string {
				return this.name;
			}

			def getAge(): number {
				return this.age;
			}
		}
		
		let person: Person = new Person("John", 30);
		let name: string = person.getName();
		let age: number = person.getAge();
		
		if (age > 20) {
			// Adult
			let message: string = name + " is an adult.";
		} else {
			// Minor
			let message: string = name + " is a minor.";
		}
	`

	// Create a new parser
	p := parser.NewParser(source)

	// Parse the source code into an AST
	ast := parser.ParseRootStatement(p)

	// Create an S-expression visitor to generate the S-expression representation
	prettyConfig := visitor_s_expression.SExpressionConfig{
		Pretty:     true,
		IndentSize: 2,
	}
	visitor := visitor_s_expression.NewSExpressionVisitorWithConfig(prettyConfig)

	// Visit the AST to generate the S-expression
	ast.Accept(visitor)

	// Print the S-expression
	fmt.Println("S-Expression representation:")
	fmt.Println(visitor.String())
}
