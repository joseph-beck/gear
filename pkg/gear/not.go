package gear

type Not struct {
	Expression Expression
}

func NewNot(expr Expression) Expression {
	return &Not{
		Expression: expr,
	}
}

func (n *Not) Type() ExpressionType {
	return NotExpression
}

func (n *Not) Evaluate(ctx *Context) (Result, error) {
	_, err := n.Expression.Evaluate(ctx.Clone())
	if err == nil {
		return Result{}, ErrFailedToMatch
	}

	tree := NewCST(CSTParam{
		Value: "not",
		Label: NewLabel(LabelParam{
			Expression: true,
		}),
	})

	result := Result{
		CST: tree,
	}

	return result, nil
}
