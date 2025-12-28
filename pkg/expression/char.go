package expression

import (
	"github.com/joseph-beck/gear/pkg/cst"
	"github.com/joseph-beck/gear/pkg/err"
)

type Char struct {
	Value rune
}

func (c Char) Type() ExpressionType {
	return CharExpression
}

func (c Char) Evaluate(input string) (Result, error) {
	if len(input) == 0 {
		return Result{
			Remaining: input,
		}, err.EndOfInput
	}

	if rune(input[0]) == c.Value {
		tree := cst.New("char")
		tree.Add(cst.New(string(c.Value)))

		return Result{
			Remaining: input[1:],
			CST:       tree,
		}, nil
	}

	return Result{
		Remaining: input,
	}, err.FailedToMatch
}
