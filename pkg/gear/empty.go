package gear

type Empty struct {
}

func NewEmpty() Expression {
	return &Empty{}
}

func (e Empty) Type() ExpressionType {
	return EmptyExpression
}

func (e Empty) Evaluate(ctx *Context) (Result, error) {
	return Result{}, nil
}
