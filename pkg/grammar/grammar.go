package grammar

import "github.com/joseph-beck/gear/pkg/rule"

type Grammar struct {
	Rules []rule.Rule
}

func New() Grammar {
	return Grammar{}
}

func (g Grammar) Get(name string) (rule.Rule, bool) {
	for _, r := range g.Rules {
		if r.Name == name {
			return r, true
		}
	}

	return rule.Rule{}, false
}

func (g *Grammar) Add(r rule.Rule) {
	g.Rules = append(g.Rules, r)
}
