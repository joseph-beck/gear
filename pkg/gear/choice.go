package gear

type Choice struct {
	Value []Expression
}

func (c *Choice) Type() ExpressionType {
	return ChoiceExpression
}

func (c *Choice) Evaluate(ctx *Context, pos uint) (Result, error) {
	for i, expr := range c.Value {
		current_ctx := ctx.Clone()

		artifact := NewArtifact(current_ctx.Rule(), i, current_ctx.Depth())

		history := current_ctx.History()

		if history.Prod(artifact) {
			continue
		}

		history.Preserve(artifact)

		current_ctx.SetHistory(history)

		result, err := expr.Evaluate(current_ctx, pos)
		if err != nil {
			continue
		}

		tree := NewCST(CSTParam{
			Value: "choice",
			Label: NewLabel(LabelParam{
				Expression: true,
			}),
		})
		tree.Add(result.CST)

		return Result{
			CST:  tree,
			Next: result.Next,
		}, nil
	}

	return Result{}, ErrFailedToMatch
}
