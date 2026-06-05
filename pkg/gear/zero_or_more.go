package gear

type ZeroOrMore struct {
	Value Expression
}

func NewZeroOrMore(expr Expression) Expression {
	return &ZeroOrMore{
		Value: expr,
	}
}

func (z *ZeroOrMore) Type() ExpressionType {
	return ZeroOrMoreExpression
}

func (z *ZeroOrMore) Evaluate(ctx *Context) (Result, error) {
	children := make([]CST, 0)

	for {
		result, err := z.Value.Evaluate(ctx)
		if err != nil {
			break
		}

		children = append(children, result.CST)
	}

	return Result{
		CST: CST{
			value:    "zero_or_more",
			children: children,
			label: label{
				expression: true,
			},
		},
	}, nil
}
