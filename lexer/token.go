package lexer

type Token struct {
	tType string
	value string
	priority int32
}

func (t *Token) Type() string {
	return t.tType
}

func (t *Token) Value() string {
	return t.value
}

func (t *Token) Priority() int32 {
	return t.priority
}

func (t *Token) String() string {
	return t.Value()
}