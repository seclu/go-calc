package matcher

import "strconv"

type IntMatcher struct {}

func (m IntMatcher) GetName() string {
	return "Integer"
}

func (m IntMatcher) GetPriority() int32 {
	return 0
}

func (m IntMatcher) IsMatch(lex string) bool {
	if _, err := strconv.Atoi(lex); err == nil {
		return true
	}

	return false
}