package gear

type Char struct {
	Value rune
}

func (c *Char) Type() ExpressionType {
	return CharExpression
}

func (c *Char) Evaluate(ctx *Context, pos uint) (Result, error) {
	input := ctx.Input()

	if pos >= uint(len(input)) {
		return Result{}, ErrEndOfInput
	}

	if rune(input[pos]) != c.Value {
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
		Next: pos + 1,
		CST:  tree,
	}

	return result, nil
}
