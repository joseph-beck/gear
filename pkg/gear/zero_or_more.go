package gear

type ZeroOrMore struct {
	Value Expression
}

func (z *ZeroOrMore) Type() ExpressionType {
	return ZeroOrMoreExpression
}

func (z *ZeroOrMore) Evaluate(ctx *Context, pos uint) (Result, error) {
	// Only use memoization if we're not in growth mode
	if !ctx.Seeding() {
		if r, err, ok := ctx.Packrat().Get(z, pos); ok {
			return r, err
		}
	}

	tree := NewCST("zero_or_more")
	current := pos

	for {
		r, err := z.Value.Evaluate(ctx, current)
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
		ctx.Packrat().Put(z, pos, result, nil)
	}
	return result, nil
}
