package expression

import (
	"github.com/joseph-beck/gear/pkg/cst"
	"github.com/joseph-beck/gear/pkg/err"
)

type OneOrMore struct {
	Value Expression
}

func (z OneOrMore) Type() ExpressionType {
	return OneOrMoreExpression
}

func (z OneOrMore) Evaluate(input string) (Result, error) {
	if len(input) == 0 {
		return Result{
			Remaining: input,
		}, err.EndOfInput
	}

	tree := cst.New("one_or_more")

	for {
		r, err := z.Value.Evaluate(input)

		if err != nil {
			break
		}

		input = r.Remaining
		tree.Add(r.CST)
	}

	if len(tree.Children) == 0 {
		return Result{
			Remaining: input,
		}, err.FailedToMatch
	}

	return Result{
		Remaining: input,
		CST:       tree,
	}, nil
}
