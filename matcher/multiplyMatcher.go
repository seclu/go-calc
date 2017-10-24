package matcher

type MultiplyMatcher struct {}

func (m MultiplyMatcher) GetName() string {
	return "MultiplyOperator"
}

func (m MultiplyMatcher) GetPriority() int32 {
	return 1
}

func (m MultiplyMatcher) IsMatch(lex string) bool {
	return lex == "*"
}