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

func (s Sequence) Evaluate(context *Context) (Result, error) {
	input := context.Remaining()

	if len(input) == 0 {
		return Result{}, err.EndOfInput
	}

	tree := NewCST("sequence")
	for _, expr := range s.Value {
		r, err := expr.Evaluate(context)

		if err != nil {
			return Result{}, err
		}

		tree.Add(r.CST)
	}

	return Result{
		CST: tree,
	}, nil
}
