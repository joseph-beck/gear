package gear

type Grammar struct {
	rules []Rule
}

type GrammarParam struct {
	Rules []Rule
}

func NewGrammar(param ...GrammarParam) Grammar {
	if len(param) == 0 {
		return Grammar{}
	}

	p := param[0]
	return Grammar{
		rules: p.Rules,
	}
}

func (g Grammar) Get(name string) (Rule, bool) {
	for _, r := range g.rules {
		if r.Name == name {
			return r, true
		}
	}

	return Rule{}, false
}

func (g *Grammar) Add(r Rule) {
	g.rules = append(g.rules, r)
}
