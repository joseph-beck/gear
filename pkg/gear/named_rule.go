package gear

import "github.com/joseph-beck/gear/pkg/errs"

type NamedRule struct {
	Value string
}

func (n *NamedRule) Type() ExpressionType {
	return NamedRuleExpression
}

func (n *NamedRule) Evaluate(ctx *Context, pos uint) (Result, error) {
	if r, err, ok := ctx.Packrat().Get(n, pos); ok {
		return r, err
	}

	if ctx.Packrat().Mark(n, pos) {
		return Result{}, errs.FailedToMatch
	}

	rule, ok := ctx.Grammar().Get(n.Value)
	if !ok {
		return Result{}, errs.RuleNotFound
	}

	r, err := rule.Expression.Evaluate(ctx, pos)

	if err != nil {
		ctx.Packrat().Put(n, pos, Result{}, err)
		return Result{}, err
	}

	tree := NewCST(cstParam{
		value: rule.Name,
		label: NewLabel(labelParam{
			expression: true,
		}),
	})
	tree.Children = append(tree.Children, r.CST)
	tree.Value = n.Value

	result := Result{
		Next: r.Next,
		CST:  tree,
	}

	ctx.Packrat().Update(n, pos, result, nil)
	ctx.Packrat().Put(n, pos, result, nil)

	for {
		last := result
		ctx.Packrat().Clear(n, pos)
		ctx.Packrat().Put(n, pos, last, nil)

		ctx.SetSeeding(true)

		r, err := rule.Expression.Evaluate(ctx, pos)

		if err != nil || r.Next <= result.Next {
			ctx.Packrat().Put(n, pos, result, nil)
			break
		}

		ctx.SetSeeding(false)

		tree := NewCST(cstParam{
			value: rule.Name,
			label: NewLabel(labelParam{
				expression: true,
			}),
		})
		tree.Children = append(tree.Children, r.CST)
		tree.Value = n.Value

		result = Result{
			Next: r.Next,
			CST:  tree,
		}

		ctx.Packrat().Put(n, pos, result, nil)
	}

	return result, nil
}
