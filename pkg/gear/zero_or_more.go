package gear

import (
	"github.com/joseph-beck/gear/pkg/err"
)

type ZeroOrMore struct {
	Value Expression
}

func (z ZeroOrMore) Type() ExpressionType {
	return ZeroOrMoreExpression
}

func (z ZeroOrMore) Evaluate(context *Context) (Result, error) {
	input := context.Remaining()

	if len(input) == 0 {
		return Result{}, err.EndOfInput
	}

	tree := NewCST("zero_or_more")

	for {
		r, err := z.Value.Evaluate(context)

		if err != nil {
			break
		}

		tree.Add(r.CST)
	}

	return Result{
		CST: tree,
	}, nil
}
