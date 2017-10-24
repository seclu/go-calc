package main

import (
	"fmt"
	"./lexer"
	"./parser"
)

func main() {
	fmt.Println("Calculator")

	p := parser.Parser{lexer.Lexer{}}
}
