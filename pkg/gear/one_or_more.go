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

func (o *OneOrMore) Evaluate(ctx *Context, pos uint) (Result, error) {
	// Only use memoization if we're not in growth mode
	if !ctx.Seeding() {
		if r, err, ok := ctx.Packrat().Get(o, pos); ok {
			return r, err
		}
	}

	tree := NewCST("one_or_more")
	current := pos

	first, err := o.Value.Evaluate(ctx, current)
	if err != nil {
		if !ctx.Seeding() {
			ctx.Packrat().Put(o, pos, Result{}, err)
		}
		return Result{}, err
	}

	if first.Next == current {
		if !ctx.Seeding() {
			ctx.Packrat().Put(o, pos, Result{}, errs.FailedToMatch)
		}
		return Result{}, errs.FailedToMatch
	}

	tree.Add(first.CST)
	current = first.Next

	for {
		r, err := o.Value.Evaluate(ctx, current)
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

	// Only memoize if not growing
	if !ctx.Seeding() {
		ctx.Packrat().Put(o, pos, result, nil)
	}
	return result, nil
}
