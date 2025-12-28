package expression

type NamedRule struct {
	value string
}

func (n NamedRule) Type() ExpressionType {
	return NamedRuleExpression
}

func (n NamedRule) Evaluate() (string, error) {
	return "", nil
}
