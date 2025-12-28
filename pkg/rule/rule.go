package rule

import "github.com/joseph-beck/gear/pkg/expression"

type Rule struct {
	Name       string
	Expression expression.Expression
}

func New(name string, expr expression.Expression) Rule {
	return Rule{
		Name:       name,
		Expression: expr,
	}
}
