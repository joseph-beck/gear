package expression

type Sequence struct {
	value []Expression
}

func (s Sequence) Type() ExpressionType {
	return SequenceExpression
}

func (s Sequence) Evaluate() (string, error) {
	return "", nil
}
