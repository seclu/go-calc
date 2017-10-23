package main

import (
	"fmt"
	"./lexer"
)

func main() {
	fmt.Println("Calculator")

	var l = lexer.Lexer{"1 + 2"}
	for _, token := range l.Lex() {
		fmt.Println(token.Value() + " - " + token.Type())
	}
}
