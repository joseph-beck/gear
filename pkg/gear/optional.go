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
	tree := NewCST(CSTParam{
		Value:    "optional",
		Children: []CST{},
		Label: NewLabel(LabelParam{
			Expression: true,
		}),
	})

	result := Result{
		CST: tree,
	}

	clonedCtx := ctx.Clone()
	r, err := o.Expression.Evaluate(clonedCtx)
	if err != nil {
		return result, nil
	}

	ctx.Update(clonedCtx)

	result.CST.Add(r.CST)

	return result, nil
}
