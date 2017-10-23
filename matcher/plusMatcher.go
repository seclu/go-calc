package matcher

type PlusMatcher struct {}

func (m PlusMatcher) GetName() string {
	return "PlusOperator"
}

func (m PlusMatcher) IsMatch(lex string) bool {
	return lex == "+"
}
