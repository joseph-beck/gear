package gear

type NamedRule struct {
	Value string
}

func NewNamedRule(value string) Expression {
	return &NamedRule{
		Value: value,
	}
}

func (n *NamedRule) Type() ExpressionType {
	return NamedRuleExpression
}

func (n *NamedRule) Evaluate(ctx *Context) (Result, error) {
	rule, ok := ctx.grammar.Get(n.Value)
	if !ok {
		return Result{}, ErrRuleNotFound
	}

	ctx.rule = rule.Name

	res, err := rule.Expression.Evaluate(ctx)
	if err != nil {
		ctx.packrat.Set(PackratKey{
			rule: n.Value,
			pos:  ctx.pos,
		}, &PackratEntry{
			result: Result{},
			err:    err,
		})

		return Result{}, err
	}

	result := Result{
		CST: NewCST(CSTParam{
			Value: rule.Name,
			Children: []CST{
				res.CST,
			},
			Label: NewLabel(LabelParam{
				Expression: true,
			}),
		}),
	}

	ctx.packrat.Set(PackratKey{
		rule: n.Value,
		pos:  ctx.pos,
	}, &PackratEntry{
		result: result,
		err:    nil,
	})

	return result, nil
}
