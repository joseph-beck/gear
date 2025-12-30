package gear

type Sequence struct {
	Value []Expression
}

func (s *Sequence) Type() ExpressionType {
	return SequenceExpression
}

func (s *Sequence) Evaluate(context *Context, pos uint) (Result, error) {
	if r, err, ok := context.Packrat().Get(s, pos); ok {
		return r, err
	}

	tree := NewCST("sequence")
	current := pos

	for _, expr := range s.Value {
		r, err := expr.Evaluate(context, current)
		if err != nil {
			context.Packrat().Put(s, pos, Result{}, err)
			return Result{}, err
		}

		tree.Add(r.CST)
		current = r.Next
	}

	result := Result{
		Next: current,
		CST:  tree,
	}

	context.Packrat().Put(s, pos, result, nil)
	return result, nil
}
