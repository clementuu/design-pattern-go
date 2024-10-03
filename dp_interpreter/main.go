package main

import "fmt"

func main() {

	expression1 := &TerminalExpression{"hello world"}
	expression2 := &TerminalExpression{"goodbye world"}

	expression3 := &OrExpression{expression1, expression2}
	fmt.Println(expression3.Interpret())

	expression4 := &TerminalExpression{"hello everyone"}
	expression5 := &AndExpression{expression3, expression4}
	fmt.Println(expression5.Interpret())

}
