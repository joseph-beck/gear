package gear

import (
	"github.com/joseph-beck/gear/pkg/err"
)

type Choice struct {
	Value []Expression
}

func (c Choice) Type() ExpressionType {
	return ChoiceExpression
}

func (c Choice) Evaluate(context *Context) (Result, error) {
	input := context.Remaining()

	if len(input) == 0 {
		return Result{}, err.EndOfInput
	}

	for _, expr := range c.Value {
		r, err := expr.Evaluate(context)

		if err != nil {
			continue
		}

		tree := NewCST("choice")
		tree.Add(r.CST)

		return Result{
			CST: tree,
		}, nil
	}

	return Result{}, err.FailedToMatch
}
