package gear

type Optional struct {
	Expression Expression
}

func NewOptional(expr Expression) Expression {
	return &Optional{
		Expression: expr,
	}
}

func (o *Optional) Type() ExpressionType {
	return OptionalExpression
}

func (o *Optional) Evaluate(ctx *Context) (Result, error) {
	return Result{}, nil
}
