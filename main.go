package main

import (
	"fmt"
	"./lexer"
	"./parser"
)

func main() {
	p := parser.Parser{lexer.Lexer{}}
	fmt.Println(p.Evaluate("2 + 2 * 2"))
}
