package gear

type And struct {
	Expression Expression
}

func NewAnd(expr Expression) Expression {
	return &And{
		Expression: expr,
	}
}

func (a *And) Type() ExpressionType {
	return AndExpression
}

func (a *And) Evaluate(ctx *Context) (Result, error) {
	return Result{}, nil
}
