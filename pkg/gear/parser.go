package gear

import (
	"github.com/joseph-beck/gear/pkg/err"
)

type Parser struct {
	grammar Grammar
}

func New(g ...Grammar) Parser {
	if len(g) > 0 {
		return Parser{
			grammar: g[0],
		}
	}

	return Parser{
		grammar: NewGrammar(),
	}
}

func (p *Parser) SetGrammar(g Grammar) {
	p.grammar = g
}

func (p Parser) Parse(input string, rule string) (Result, error) {
	r, ok := p.grammar.Get(rule)

	if !ok {
		return Result{}, err.RuleNotFound
	}

	context := NewContext(contextCfg{
		input:   input,
		grammar: &p.grammar,
	})

	res, err := r.Expression.Evaluate(context)
	if err != nil {
		return Result{}, err
	}

	tree := NewCST(rule)
	tree.Add(res.CST)

	return Result{
		CST: tree,
	}, nil
}

func (p *Parser) DefaultResolver(name string) (Expression, error) {
	r, ok := p.grammar.Get(name)

	if !ok {
		return nil, err.RuleNotFound
	}

	return r.Expression, nil
}
