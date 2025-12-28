package expression

type Choice struct {
	value []Expression
}

func (c Choice) Type() ExpressionType {
	return ChoiceExpression
}

func (c Choice) Evaluate() (string, error) {
	return "", nil
}
