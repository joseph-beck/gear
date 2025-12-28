package expression

import "github.com/joseph-beck/gear/pkg/cst"

type NamedRule struct {
	value   string
	resolve func(name string) (Expression, error)
}

func (n NamedRule) Type() ExpressionType {
	return NamedRuleExpression
}

func (n NamedRule) Evaluate(input string) (Result, error) {
	rule, err := n.resolve(n.value)

	if err != nil {
		return Result{
			remaining: input,
		}, err
	}

	r, err := rule.Evaluate(input)
	if err != nil {
		return Result{
			remaining: input,
		}, err
	}

	tree := cst.New(n.value)
	tree.Children = append(tree.Children, r.cst)
	tree.Value = n.value

	return Result{
		remaining: r.remaining,
		cst:       tree,
	}, nil
}
