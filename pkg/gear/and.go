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
	_, err := a.Expression.Evaluate(ctx.Clone())
	if err != nil {
		return Result{}, ErrFailedToMatch
	}

	tree := NewCST(CSTParam{
		Value: "and",
		Label: NewLabel(LabelParam{
			Expression: true,
		}),
	})

	result := Result{
		CST: tree,
	}

	return result, nil
}
