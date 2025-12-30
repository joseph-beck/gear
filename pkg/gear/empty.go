package gear

type Empty struct {
}

func (e Empty) Type() ExpressionType {
	return EmptyExpression
}

func (e Empty) Evaluate(input string) (Result, error) {
	return Result{}, nil
}
