package matcher

type Matcher interface {
	GetName() string
	IsMatch(lex string) bool
	GetPriority() int32
}
