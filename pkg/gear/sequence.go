package gear

type Sequence struct {
	Value []Expression
}

func (s *Sequence) Type() ExpressionType {
	return SequenceExpression
}

func (s *Sequence) Evaluate(ctx *Context) (Result, error) {
	children := make([]CST, 0)

	for _, expr := range s.Value {
		if ctx.pos >= uint(len(ctx.input)) {
			return Result{}, ErrEndOfInput
		}

		res, err := expr.Evaluate(ctx)
		if err != nil {
			return Result{}, ErrFailedToMatch
		}

		children = append(children, res.CST)
	}

	return Result{
		CST: NewCST(CSTParam{
			Value:    "sequence",
			Children: children,
			Label: NewLabel(LabelParam{
				Expression: true,
			}),
		}),
	}, nil
}
