package gear

type Choice struct {
	Value []Expression
}

func NewChoice(value []Expression) Expression {
	return &Choice{
		Value: value,
	}
}

func (c *Choice) Type() ExpressionType {
	return ChoiceExpression
}

func (c *Choice) Evaluate(ctx *Context) (Result, error) {
	if ctx.pos >= uint(len(ctx.input)) {
		return Result{}, ErrEndOfInput
	}

	for i, expr := range c.Value {
		currentCtx := ctx.Clone()

		artifact := NewArtifact(currentCtx.rule, i, currentCtx.depth)
		history := currentCtx.history

		if history.Prod(artifact) {
			continue
		}

		history.Preserve(artifact)

		result, err := expr.Evaluate(currentCtx)
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

		ctx.Update(currentCtx)

		ctx.depth++

		return Result{
			CST: tree,
		}, nil
	}

	return Result{}, ErrFailedToMatch
}
