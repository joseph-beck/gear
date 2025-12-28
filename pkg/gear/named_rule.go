package gear

import "github.com/joseph-beck/gear/pkg/err"

type NamedRule struct {
	Value string
}

func (n NamedRule) Type() ExpressionType {
	return NamedRuleExpression
}

func (n NamedRule) Evaluate(context *Context) (Result, error) {
	rule, ok := context.Grammar().Get(n.Value)

	if !ok {
		return Result{}, err.RuleNotFound
	}

	r, err := rule.Expression.Evaluate(context)
	if err != nil {
		return Result{}, err
	}

	tree := NewCST(rule.Name)
	tree.Children = append(tree.Children, r.CST)
	tree.Value = n.Value

	return Result{
		CST: tree,
	}, nil
}
