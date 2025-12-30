package gear

import "github.com/joseph-beck/gear/pkg/errs"

type Char struct {
	Value rune
}

func (c *Char) Type() ExpressionType {
	return CharExpression
}

func (c *Char) Evaluate(ctx *Context, pos uint) (Result, error) {
	input := ctx.Input()

	if pos >= uint(len(input)) {
		return Result{}, errs.EndOfInput
	}

	if rune(input[pos]) != c.Value {
		return Result{}, errs.FailedToMatch
	}

	tree := NewCST("char")
	tree.Add(NewCST(string(c.Value)))

	result := Result{
		Next: pos + 1,
		CST:  tree,
	}

	return result, nil
}
