package lexer

type Token struct
{
	lex string
	value string
}

func (t *Token) Type() string {
	return t.lex
}

func (t *Token) Value() string {
	return t.value
}

func (t *Token) String() string {
	return t.Value()
}