package expression

import (
	"github.com/joseph-beck/gear/pkg/cst"
	"github.com/joseph-beck/gear/pkg/err"
)

type ZeroOrMore struct {
	value Expression
}

func (z ZeroOrMore) Type() ExpressionType {
	return ZeroOrMoreExpression
}

func (z ZeroOrMore) Evaluate(input string) (Result, error) {
	if len(input) == 0 {
		return Result{
			remaining: input,
		}, err.EndOfInput
	}

	tree := cst.New("zero_or_more")

	for {
		r, err := z.value.Evaluate(input)

		if err != nil {
			break
		}

		input = r.remaining
		tree.Add(r.cst)
	}

	return Result{
		remaining: input,
		cst:       tree,
	}, nil
}
