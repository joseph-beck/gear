package expression

import (
	"github.com/joseph-beck/gear/pkg/cst"
	"github.com/joseph-beck/gear/pkg/err"
)

type Choice struct {
	Value []Expression
}

func (c Choice) Type() ExpressionType {
	return ChoiceExpression
}

func (c Choice) Evaluate(input string) (Result, error) {
	if len(input) == 0 {
		return Result{
			Remaining: input,
		}, err.EndOfInput
	}

	for _, expr := range c.Value {
		r, err := expr.Evaluate(input)

		if err != nil {
			continue
		}

		tree := cst.New("choice")
		tree.Add(r.CST)

		return Result{
			Remaining: r.Remaining,
			CST:       tree,
		}, nil
	}

	return Result{
		Remaining: input,
	}, err.FailedToMatch
}
