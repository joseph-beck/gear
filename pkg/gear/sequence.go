package gear

import "github.com/joseph-beck/gear/pkg/errs"

type Sequence struct {
	Value []Expression
}

func (s *Sequence) Type() ExpressionType {
	return SequenceExpression
}

func (s *Sequence) Evaluate(ctx *Context, pos uint) (Result, error) {
	if !ctx.Seeding() {
		if r, err, ok := ctx.Packrat().Get(s, pos); ok {
			return r, err
		}
	}

	if pos >= uint(len(ctx.Input())) {
		return Result{}, errs.EndOfInput
	}

	tree := NewCST("sequence")
	current := pos

	for _, expr := range s.Value {
		r, err := expr.Evaluate(ctx, current)
		if err != nil {
			if !ctx.Seeding() {
				ctx.Packrat().Put(s, pos, Result{}, err)
			}

			return Result{}, err
		}

		tree.Add(r.CST)
		current = r.Next
	}

	result := Result{
		Next: current,
		CST:  tree,
	}

	if !ctx.Seeding() {
		ctx.Packrat().Put(s, pos, result, nil)
	}

	return result, nil
}
