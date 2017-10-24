package lexer

import (
	"strings"
	"../matcher"
)

var matchers []matcher.Matcher

type Lexer struct {}

func init() {
	matchers = append(matchers, matcher.IntMatcher{})
	matchers = append(matchers, matcher.WhiteSpaceMatcher{})
	matchers = append(matchers, matcher.PlusMatcher{})
	matchers = append(matchers, matcher.MultiplyMatcher{})
}

func (l *Lexer) Tokenize(expression string) []Token {
	var tokens []Token

	for _, value := range strings.Split(expression, "") {
		var tType string
		var priority int32
		var token Token

		for _, m := range matchers {
			if m.IsMatch(value) {
				tType = m.GetName()
				priority = m.GetPriority()
			}
		}

		token = Token{tType: tType, value: value, priority: priority}

		tokens = append(tokens, token)
	}

	return tokens
}