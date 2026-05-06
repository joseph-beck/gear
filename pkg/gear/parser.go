package gear

type Parser struct {
	grammar *Grammar
}

type ParserParam struct {
	Grammar Grammar
}

type ParserResult struct {
	CST       CST
	Remaining string
}

func New(param ...ParserParam) Parser {
	if len(param) == 0 {
		grammar := NewGrammar()

		return Parser{
			grammar: &grammar,
		}
	}

	p := param[0]
	return Parser{
		grammar: &p.Grammar,
	}
}

func (p *Parser) Parse(input string, rule string) (ParserResult, error) {
	_, ok := p.grammar.Get(rule)

	if !ok {
		return ParserResult{}, ErrRuleNotFound
	}

	ctx := NewContext(input)
	ctx.grammar = p.grammar

	named := &NamedRule{
		Value: rule,
	}

	res, err := named.Evaluate(ctx)
	if err != nil {
		return ParserResult{}, err
	}

	remaining := input[ctx.pos:]

	return ParserResult{
		CST:       res.CST,
		Remaining: remaining,
	}, nil
}
