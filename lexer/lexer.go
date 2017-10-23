package lexer

import (
	"strings"
	"../matcher"
)

var matchers []matcher.Matcher

type Lexer struct
{
	 Input string
}

func init() {
	matchers = append(matchers, matcher.IntMatcher{})
	matchers = append(matchers, matcher.WhiteSpaceMatcher{})
	matchers = append(matchers, matcher.PlusMatcher{})
}

func (l *Lexer) Lex() []Token {
	var tokens []Token

	for _, value := range strings.Split(l.Input, "") {
		var lex string
		var token Token

		for _, m := range matchers {
			if m.IsMatch(value) {
				lex = m.GetName()
			}
		}

		token = Token{lex: lex, value: value}

		tokens = append(tokens, token)
	}

	return tokens
}