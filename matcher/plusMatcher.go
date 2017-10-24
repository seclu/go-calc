package matcher

type PlusMatcher struct {}

func (m PlusMatcher) GetName() string {
	return "PlusOperator"
}

func (m PlusMatcher) GetPriority() int32 {
	return 0
}

func (m PlusMatcher) IsMatch(lex string) bool {
	return lex == "+"
}
