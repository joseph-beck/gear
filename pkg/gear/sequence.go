package gear

type Sequence struct {
	Value []Expression
}

func (s *Sequence) Type() ExpressionType {
	return SequenceExpression
}

func (s *Sequence) Evaluate(ctx *Context, pos uint) (Result, error) {
	// Only use memoization if we're not in growth mode
	if !ctx.Seeding() {
		if r, err, ok := ctx.Packrat().Get(s, pos); ok {
			return r, err
		}
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

	// Only memoize if not growing
	if !ctx.Seeding() {
		ctx.Packrat().Put(s, pos, result, nil)
	}
	return result, nil
}
