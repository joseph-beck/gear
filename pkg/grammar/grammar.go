package grammar

import "github.com/joseph-beck/gear/pkg/rule"

type Grammar struct {
	rules []rule.Rule
}

func New() Grammar {
	return Grammar{}
}
