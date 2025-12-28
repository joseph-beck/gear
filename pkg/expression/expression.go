package expression

type ExpressionType int

const (
	CharExpression ExpressionType = iota
	ChoiceExpression
	SequenceExpression
	ZeroOrMoreExpression
	OneOrMoreExpression
	NamedRuleExpression
)

type Expression interface {
	Type() ExpressionType
	Evaluate() (string, error)
}
