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
		current_ctx := ctx.Clone()

		artifact := NewArtifact(current_ctx.rule, i, current_ctx.depth)

		history := current_ctx.history

		if history.Prod(artifact) {
			continue
		}

		history.Preserve(artifact)

		result, err := expr.Evaluate(current_ctx)
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
			CST: tree,
		}, nil
	}

	return Result{}, ErrFailedToMatch
}
