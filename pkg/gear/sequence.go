package gear

import (
	"github.com/joseph-beck/gear/pkg/err"
)

type Sequence struct {
	Value []Expression
}

func (s Sequence) Type() ExpressionType {
	return SequenceExpression
}

func (s Sequence) Evaluate(input string) (Result, error) {
	if len(input) == 0 {
		return Result{
			Remaining: input,
		}, err.EndOfInput
	}

	tree := NewCST("sequence")
	for _, expr := range s.Value {
		r, err := expr.Evaluate(input)

		if err != nil {
			return Result{
				Remaining: input,
			}, err
		}

		input = r.Remaining
		tree.Add(r.CST)
	}

	return Result{
		Remaining: input,
		CST:       tree,
	}, nil
}
