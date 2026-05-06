package gear

type NamedRule struct {
	Value string
}

func (n *NamedRule) Type() ExpressionType {
	return NamedRuleExpression
}

func (n *NamedRule) Evaluate(ctx *Context, pos uint) (Result, error) {
	rule, ok := ctx.Grammar().Get(n.Value)
	if !ok {
		return Result{}, ErrRuleNotFound
	}

	ctx.SetDepth(ctx.depth + 1)
	ctx.SetRule(rule.Name)

	return rule.Expression.Evaluate(ctx, pos)
}
