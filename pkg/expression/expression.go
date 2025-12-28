package expression

import (
	"github.com/joseph-beck/gear/pkg/cst"
)

type ExpressionType int

const (
	EmptyExpression ExpressionType = iota
	CharExpression
	ChoiceExpression
	SequenceExpression
	ZeroOrMoreExpression
	OneOrMoreExpression
	NamedRuleExpression
)

type Expression interface {
	Type() ExpressionType
	Evaluate(string) (Result, error)
}

type Result struct {
	Remaining string
	CST       cst.CST
}
