package expression

import (
	"github.com/joseph-beck/gear/pkg/cst"
	"github.com/joseph-beck/gear/pkg/err"
)

type Char struct {
	value rune
}

func (c Char) Type() ExpressionType {
	return CharExpression
}

func (c Char) Evaluate(input string) (Result, error) {
	if len(input) == 0 {
		return Result{
			remaining: input,
		}, err.EndOfInput
	}

	if rune(input[0]) == c.value {
		tree := cst.New("char")
		tree.Add(cst.New(string(c.value)))

		return Result{
			remaining: input[1:],
			cst:       tree,
		}, nil
	}

	return Result{
		remaining: input,
	}, err.FailedToMatch
}
