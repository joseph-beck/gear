package expression

type Choice struct {
	value []Expression
}

func (c Choice) Type() int {
	return ChoiceExpression
}

func (c Choice) Evaluate() (string, error) {
	return "", nil
}
