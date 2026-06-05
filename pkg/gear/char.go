package gear

type Char struct {
	Value rune
}

func NewChar(value rune) Expression {
	return &Char{
		Value: value,
	}
}

func (c *Char) Type() ExpressionType {
	return CharExpression
}

func (c *Char) Evaluate(ctx *Context) (Result, error) {
	if ctx.pos >= uint(len(ctx.input)) {
		return Result{}, ErrEndOfInput
	}

	if rune(ctx.input[ctx.pos]) != c.Value {
		return Result{}, ErrFailedToMatch
	}

	tree := NewCST(CSTParam{
		Value: "char",
		Label: NewLabel(LabelParam{
			Expression: true,
		}),
	})

	tree.Add(NewCST(CSTParam{
		Value: string(c.Value),
	}))

	result := Result{
		CST: tree,
	}

	ctx.pos++
	ctx.history.Clear()

	return result, nil
}
