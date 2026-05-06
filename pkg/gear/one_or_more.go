package gear

type OneOrMore struct {
	Value Expression
}

func (o *OneOrMore) Type() ExpressionType {
	return OneOrMoreExpression
}

func (o *OneOrMore) Evaluate(ctx *Context) (Result, error) {
	if ctx.pos >= uint(len(ctx.input)) {
		return Result{}, ErrEndOfInput
	}

	children := make([]CST, 0)

	for {
		res, err := o.Value.Evaluate(ctx)
		if err != nil {
			break
		}

		children = append(children, res.CST)
	}

	if len(children) == 0 {
		return Result{}, ErrFailedToMatch
	}

	return Result{
		CST: NewCST(CSTParam{
			Value:    "one_or_more",
			Children: children,
			Label: NewLabel(LabelParam{
				Expression: true,
			}),
		}),
	}, nil
}
