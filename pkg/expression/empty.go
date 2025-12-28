package expression

type Empty struct{}

func (e Empty) Type() ExpressionType {
	return EmptyExpression
}

func (e Empty) Evaluate() (string, error) {
	return "", nil
}
