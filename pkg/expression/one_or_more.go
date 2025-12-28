package expression

type OneOrMore struct {
	value Expression
}

func (z OneOrMore) Type() int {
	return OneOrMoreExpression
}

func (z OneOrMore) Evaluate() (string, error) {
	return "", nil
}
