package expression

import (
	"github.com/joseph-beck/gear/pkg/cst"
	"github.com/joseph-beck/gear/pkg/err"
)

type Choice struct {
	value []Expression
}

func (c Choice) Type() ExpressionType {
	return ChoiceExpression
}

func (c Choice) Evaluate(input string) (Result, error) {
	if len(input) == 0 {
		return Result{
			remaining: input,
		}, err.EndOfInput
	}

	for _, expr := range c.value {
		r, err := expr.Evaluate(input)

		if err != nil {
			continue
		}

		tree := cst.New("choice")
		tree.Add(r.cst)

		return Result{
			remaining: r.remaining,
			cst:       tree,
		}, nil
	}

	return Result{
		remaining: input,
	}, err.FailedToMatch
}
