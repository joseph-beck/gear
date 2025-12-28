package gear

import (
	"github.com/joseph-beck/gear/pkg/err"
)

type OneOrMore struct {
	Value Expression
}

func (z OneOrMore) Type() ExpressionType {
	return OneOrMoreExpression
}

func (z OneOrMore) Evaluate(context *Context) (Result, error) {
	input := context.Remaining()

	if len(input) == 0 {
		return Result{}, err.EndOfInput
	}

	tree := NewCST("one_or_more")

	for {
		r, err := z.Value.Evaluate(context)

		if err != nil {
			break
		}

		tree.Add(r.CST)
	}

	if len(tree.Children) == 0 {
		return Result{}, err.FailedToMatch
	}

	return Result{
		CST: tree,
	}, nil
}
