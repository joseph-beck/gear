package gear

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
	Evaluate(*Context, uint) (Result, error)
}

type Result struct {
	Next uint
	CST  cst
}
