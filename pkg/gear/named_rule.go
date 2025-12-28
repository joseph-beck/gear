package gear

type NamedRule struct {
	Value   string
	Resolve func(name string) (Expression, error)
}

func (n NamedRule) Type() ExpressionType {
	return NamedRuleExpression
}

func (n NamedRule) Evaluate(input string) (Result, error) {
	rule, err := n.Resolve(n.Value)

	if err != nil {
		return Result{
			Remaining: input,
		}, err
	}

	r, err := rule.Evaluate(input)
	if err != nil {
		return Result{
			Remaining: input,
		}, err
	}

	tree := NewCST(n.Value)
	tree.Children = append(tree.Children, r.CST)
	tree.Value = n.Value

	return Result{
		Remaining: r.Remaining,
		CST:       tree,
	}, nil
}
