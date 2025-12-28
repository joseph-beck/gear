package expression

const (
	CharExpression = iota + 1
	ChoiceExpression
	SequenceExpression
	ZeroOrMoreExpression
	OneOrMoreExpression
)

type Expression interface {
	Type()
	Evaluate() (string, error)
}
