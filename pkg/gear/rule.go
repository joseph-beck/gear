package gear

type Rule struct {
	Name       string
	Expression Expression
}

func NewRule(name string, expr Expression) Rule {
	return Rule{
		Name:       name,
		Expression: expr,
	}
}
