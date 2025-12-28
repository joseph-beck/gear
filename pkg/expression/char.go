package expression

type Char struct {
	value rune
}

func (c Char) Type() ExpressionType {
	return CharExpression
}

func (c Char) Evaluate() (string, error) {
	return "", nil
}
