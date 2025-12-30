package gear

type Grammar struct {
	Rules []Rule
}

func NewGrammar() Grammar {
	return Grammar{}
}

func (g Grammar) Get(name string) (Rule, bool) {
	for _, r := range g.Rules {
		if r.Name == name {
			return r, true
		}
	}

	return Rule{}, false
}

func (g *Grammar) Add(r Rule) {
	g.Rules = append(g.Rules, r)
}
