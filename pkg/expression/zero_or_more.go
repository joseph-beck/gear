package expression

type ZeroOrMore struct {
	value Expression
}

func (z ZeroOrMore) Type() int {
	return ZeroOrMoreExpression
}

func (z ZeroOrMore) Evaluate() (string, error) {
	return "", nil
}