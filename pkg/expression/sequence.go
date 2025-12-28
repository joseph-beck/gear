package expression

import (
	"github.com/joseph-beck/gear/pkg/cst"
	"github.com/joseph-beck/gear/pkg/err"
)

type Sequence struct {
	value []Expression
}

func (s Sequence) Type() ExpressionType {
	return SequenceExpression
}

func (s Sequence) Evaluate(input string) (Result, error) {
	if len(input) == 0 {
		return Result{
			remaining: input,
		}, err.EndOfInput
	}

	tree := cst.New("sequence")
	for _, expr := range s.value {
		r, err := expr.Evaluate(input)

		if err != nil {
			return Result{
				remaining: input,
			}, err
		}

		input = r.remaining
		tree.Add(r.cst)
	}

	return Result{
		remaining: input,
		cst:       tree,
	}, nil
}
