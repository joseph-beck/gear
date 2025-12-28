package parser

import (
	"github.com/joseph-beck/gear/pkg/cst"
	"github.com/joseph-beck/gear/pkg/err"
	"github.com/joseph-beck/gear/pkg/expression"
	"github.com/joseph-beck/gear/pkg/grammar"
)

type Parser struct {
	grammar grammar.Grammar
}

func New(g ...grammar.Grammar) Parser {
	if len(g) > 0 {
		return Parser{
			grammar: g[0],
		}
	}

	return Parser{
		grammar: grammar.New(),
	}
}

func (p Parser) Parse(input string, rule string) (expression.Result, error) {
	r, ok := p.grammar.Get(rule)

	if !ok {
		return expression.Result{}, err.RuleNotFound
	}

	res, err := r.Expression.Evaluate(input)
	if err != nil {
		return expression.Result{}, err
	}

	tree := cst.New(rule)
	tree.Add(res.CST)

	return expression.Result{
		Remaining: res.Remaining,
		CST:       tree,
	}, nil
}

func (p *Parser) DefaultResolver(name string) (expression.Expression, error) {
	r, ok := p.grammar.Get(name)

	if !ok {
		return nil, err.RuleNotFound
	}

	return r.Expression, nil
}
