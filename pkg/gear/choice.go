package gear

import "github.com/joseph-beck/gear/pkg/errs"

type Choice struct {
	Value []Expression
}

func (c *Choice) Type() ExpressionType {
	return ChoiceExpression
}

func (c *Choice) Evaluate(context *Context, pos uint) (Result, error) {
	if r, err, ok := context.Packrat().Get(c, pos); ok {
		return r, err
	}

	for _, expr := range c.Value {
		r, err := expr.Evaluate(context, pos)
		if err == nil {
			tree := NewCST("choice")
			tree.Add(r.CST)

			result := Result{
				Next: r.Next,
				CST:  tree,
			}

			context.Packrat().Put(c, pos, result, nil)
			return result, nil
		}
	}

	e := errs.FailedToMatch
	context.Packrat().Put(c, pos, Result{}, e)
	return Result{}, e
}
