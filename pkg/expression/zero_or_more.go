package expression

import (
	"github.com/joseph-beck/gear/pkg/cst"
	"github.com/joseph-beck/gear/pkg/err"
)

type ZeroOrMore struct {
	Value Expression
}

func (z ZeroOrMore) Type() ExpressionType {
	return ZeroOrMoreExpression
}

func (z ZeroOrMore) Evaluate(input string) (Result, error) {
	if len(input) == 0 {
		return Result{
			Remaining: input,
		}, err.EndOfInput
	}

	tree := cst.New("zero_or_more")

	for {
		r, err := z.Value.Evaluate(input)

		if err != nil {
			break
		}

		input = r.Remaining
		tree.Add(r.CST)
	}

	return Result{
		Remaining: input,
		CST:       tree,
	}, nil
}
