package rule

import "github.com/joseph-beck/gear/pkg/expression"

type Rule struct {
	name       string
	expression expression.Expression
}

func New() Rule {
	return Rule{}
}
