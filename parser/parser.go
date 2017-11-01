package parser

import (
	"../lexer"
	"gopkg.in/oleiade/lane.v1"
	"errors"
	"strconv"
)

type Parser struct {
	Lexer lexer.Lexer
}

func (p *Parser) Evaluate(expression string) string {
	tokens := p.Lexer.Tokenize(expression)
	rpn, _ := p.parse(tokens)
	result := p.evaluateRPN(*rpn)

	return strconv.Itoa(result)
}

func (p *Parser) parse(tokens []lexer.Token) (*lane.Queue, error) {
	outputQueue := lane.NewQueue()
	operatorStack := lane.NewStack()

	for _, token := range tokens {
		switch token.Type() {
		case "Integer":
			outputQueue.Enqueue(token)
			break
		case "MultiplyOperator", "PlusOperator":
			o1 := token
			var o2 lexer.Token

			if operatorStack.First() != nil {
				o2 = operatorStack.First().(lexer.Token)
			}

			for p.hasOperatorInStack(*operatorStack) && o1.Priority() < o2.Priority() {
				outputQueue.Enqueue(operatorStack.Pop())
			}

			operatorStack.Push(o1)
			break
		case "WhiteSpace": break
		default:
			return lane.NewQueue(), errors.New("invalid type")
		}
	}

	for i := 0 ; i <= operatorStack.Size() ; i++ {
		outputQueue.Enqueue(operatorStack.Pop())
	}

	return outputQueue, nil
}

func (p *Parser) evaluateRPN(expressionTokens lane.Queue) int {
	stack := lane.NewStack()

	for expressionTokens.Head() != nil {
		t := expressionTokens.Pop().(lexer.Token)
		tokenValue := t.Value()

		if t.Type() == "Integer" {
			tokenValue, _ := strconv.Atoi(tokenValue)
			stack.Push(tokenValue)
			continue
		}

		switch tokenValue {
		case "+":
			stack.Push(stack.Pop().(int) + stack.Pop().(int))
			break
		case "*":
			stack.Push(stack.Pop().(int) * stack.Pop().(int))
			break
		}
	}

	return stack.First().(int)
}

func (p *Parser) hasOperatorInStack(stack lane.Stack) bool {
	if stack.Empty() {
		return false
	}


	if stack.Last() == nil {
		return false
	}

	var last lexer.Token = stack.Last().(lexer.Token)

	if last.Type() == "MultiplyOperator" || last.Type() == "PlusOperator" {
		return true
	}

	return false
}