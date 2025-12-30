package gear

import "github.com/joseph-beck/gear/pkg/errs"

type Parser struct {
	grammar Grammar
}

type ParserParam struct {
	Grammar Grammar
}

func New(param ...ParserParam) Parser {
	if len(param) == 0 {

		return Parser{
			grammar: NewGrammar(),
		}
	}

	p := param[0]
	return Parser{
		grammar: p.Grammar,
	}
}

func (p *Parser) SetGrammar(g Grammar) {
	p.grammar = g
}

func (p *Parser) Parse(input string, rule string) (Result, error) {
	_, ok := p.grammar.Get(rule)

	if !ok {
		return Result{}, errs.RuleNotFound
	}

	ctx := NewContext(input)
	ctx.grammar = &p.grammar

	named := &NamedRule{
		Value: rule,
	}

	res, err := named.Evaluate(ctx, 0)
	if err != nil {
		return Result{}, err
	}

	return res, nil
}
