package expression

import (
	"github.com/joseph-beck/gear/pkg/cst"
	"github.com/joseph-beck/gear/pkg/err"
)

type OneOrMore struct {
	value Expression
}

func (z OneOrMore) Type() ExpressionType {
	return OneOrMoreExpression
}

func (z OneOrMore) Evaluate(input string) (Result, error) {
	if len(input) == 0 {
		return Result{
			remaining: input,
		}, err.EndOfInput
	}

	tree := cst.New("one_or_more")

	for {
		r, err := z.value.Evaluate(input)

		if err != nil {
			break
		}

		input = r.remaining
		tree.Add(r.cst)
	}

	if len(tree.Children) == 0 {
		return Result{
			remaining: input,
		}, err.FailedToMatch
	}

	return Result{
		remaining: input,
		cst:       tree,
	}, nil
}
