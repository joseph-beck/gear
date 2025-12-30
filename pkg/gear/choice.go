package gear

import "github.com/joseph-beck/gear/pkg/errs"

type Choice struct {
	Value []Expression
}

func (c *Choice) Type() ExpressionType {
	return ChoiceExpression
}

func (c *Choice) Evaluate(ctx *Context, pos uint) (Result, error) {
	if !ctx.Seeding() {
		if r, err, ok := ctx.Packrat().Get(c, pos); ok {
			return r, err
		}
	}

	for _, expr := range c.Value {
		r, err := expr.Evaluate(ctx, pos)
		if err != nil {
			continue
		}

		tree := NewCST("choice")
		tree.Add(r.CST)

		result := Result{
			Next: r.Next,
			CST:  tree,
		}

		if !ctx.Seeding() {
			ctx.Packrat().Put(c, pos, result, nil)
		}

		return result, nil
	}

	err := errs.FailedToMatch
	if !ctx.Seeding() {
		ctx.Packrat().Put(c, pos, Result{}, err)
	}

	return Result{}, err
}
