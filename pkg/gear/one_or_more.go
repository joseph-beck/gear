package gear

import (
	"github.com/joseph-beck/gear/pkg/errs"
)

type OneOrMore struct {
	Value Expression
}

func (o *OneOrMore) Type() ExpressionType {
	return OneOrMoreExpression
}

func (o *OneOrMore) Evaluate(context *Context, pos uint) (Result, error) {
	if r, err, ok := context.Packrat().Get(o, pos); ok {
		return r, err
	}

	tree := NewCST("one_or_more")
	current := pos

	first, err := o.Value.Evaluate(context, current)
	if err != nil {
		context.Packrat().Put(o, pos, Result{}, err)
		return Result{}, err
	}

	if first.Next == current {
		context.Packrat().Put(o, pos, Result{}, errs.FailedToMatch)
		return Result{}, errs.FailedToMatch
	}

	tree.Add(first.CST)
	current = first.Next

	for {
		r, err := o.Value.Evaluate(context, current)
		if err != nil {
			break
		}

		if r.Next == current {
			break
		}

		tree.Add(r.CST)
		current = r.Next
	}

	result := Result{
		Next: current,
		CST:  tree,
	}

	context.Packrat().Put(o, pos, result, nil)
	return result, nil
}
