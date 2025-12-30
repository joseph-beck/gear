package gear

import "github.com/joseph-beck/gear/pkg/errs"

type NamedRule struct {
	Value string
}

func (n *NamedRule) Type() ExpressionType {
	return NamedRuleExpression
}

func (n *NamedRule) Evaluate(context *Context, pos uint) (Result, error) {
	if r, err, ok := context.Packrat().Get(n, pos); ok {
		return r, err
	}

	rule, ok := context.Grammar().Get(n.Value)

	if !ok {
		return Result{}, errs.RuleNotFound
	}

	r, err := rule.Expression.Evaluate(context, pos)
	if err != nil {
		return Result{}, err
	}

	tree := NewCST(rule.Name)
	tree.Children = append(tree.Children, r.CST)
	tree.Value = n.Value

	result := Result{
		Next: r.Next,
		CST:  tree,
	}

	context.Packrat().Put(n, pos, result, nil)
	return result, nil
}
