package gear

type Literal struct {
	Value string
}

func (l *Literal) Type() ExpressionType {
	return LiteralExpression
}

func (l *Literal) Evaluate(ctx *Context) (Result, error) {
	if ctx.pos >= uint(len(ctx.input)) {
		return Result{}, ErrEndOfInput
	}

	if len(ctx.input[ctx.pos:]) < len(l.Value) {
		return Result{}, ErrEndOfInput
	}

	if ctx.input[ctx.pos:][0:len(l.Value)] != l.Value {
		return Result{}, ErrFailedToMatch
	}

	tree := NewCST(CSTParam{
		Value: "literal",
		Label: NewLabel(LabelParam{
			Expression: true,
		}),
	})

	tree.Add(NewCST(CSTParam{
		Value: l.Value,
	}))

	result := Result{
		CST: tree,
	}

	ctx.pos += uint(len(l.Value))
	ctx.history.Clear()

	return result, nil
}
