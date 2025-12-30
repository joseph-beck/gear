package gear

import "github.com/joseph-beck/gear/pkg/errs"

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

func (p *Parser) Parse(input string, rule string) (Result, error) {
	r, ok := p.grammar.Get(rule)

	if !ok {
		return Result{}, errs.RuleNotFound
	}

	context := NewContext(input)
	context.grammar = &p.grammar

	res, err := r.Expression.Evaluate(context, 0)
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
		return nil, errs.RuleNotFound
	}

	return r.Expression, nil
}
