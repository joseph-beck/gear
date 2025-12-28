package expression

type Sequence struct {
	value []Expression
}

func (s Sequence) Type() int {
	return SequenceExpression
}

func (s Sequence) Evaluate() (string, error) {
	return "", nil
}
