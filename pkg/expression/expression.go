package expression

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
	Evaluate() (string, error)
}
