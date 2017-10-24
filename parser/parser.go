package parser

import (
	"../lexer"
	"gopkg.in/oleiade/lane.v1"
)

type Parser struct {
	lexer lexer.Lexer
}

func (p *Parser) Evaluate(expression string) string {
	tokens := p.lexer.Tokenize(expression)
	rpn := p.parse(tokens)
}

func (p *Parser) parse(tokens []lexer.Token) lane.Queue {
	outputQueue := lane.NewQueue()
	operatorStack := lane.NewStack()

	for _, token := range tokens {
		switch token.Type() {
		case "Integer":
			outputQueue.Enqueue(token)
		case "MultiplyOperator":
		case "PlusOperator":
			o1 := token
			o2 := operatorStack.Last()
			if o2 != nil {
				for p.hasOperatorInStack(operatorStack) && o1.Priority() < o2.Priority() {
					outputQueue.Enqueue(operatorStack.Pop())
				}
			}

			operatorStack.Push(o1)
		default:
			panic("Invalid type!")
		}
	}

	return outputQueue
}

func (p *Parser) hasOperatorInStack(stack lane.Stack) bool {
	if stack.Empty() {
		return false
	}

	last := stack.Last()
	if last == nil {
		return false
	}

	if last.Type() == "MultiplyOperator" || last.Type() == "PlusOperator" {
		return true
	}

	return false
}