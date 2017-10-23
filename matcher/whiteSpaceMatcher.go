package matcher

import "unicode"

type WhiteSpaceMatcher struct {}

func (m WhiteSpaceMatcher) GetName() string {
	return "WhiteSpace"
}

func (m WhiteSpaceMatcher) IsMatch(lex string) bool {
	r := []rune(lex)

	return unicode.IsSpace(r[0])
}