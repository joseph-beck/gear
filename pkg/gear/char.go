package gear

import (
	"github.com/joseph-beck/gear/pkg/err"
)

type Char struct {
	Value rune
}

func (c Char) Type() ExpressionType {
	return CharExpression
}

func (c Char) Evaluate(context *Context) (Result, error) {
	input := context.Remaining()

	if len(input) == 0 {
		return Result{}, err.EndOfInput
	}

	if rune(input[0]) == c.Value {
		tree := NewCST("char")
		tree.Add(NewCST(string(c.Value)))

		context.SetPosition(context.Position() + 1)

		return Result{
			CST: tree,
		}, nil
	}

	return Result{}, err.FailedToMatch
}
