package main

import (
	"fmt"

	"github.com/joseph-beck/gear/pkg/gear"
)

func main() {
	digit := &gear.Choice{
		Value: []gear.Expression{
			&gear.Char{
				Value: '0',
			},
			&gear.Char{
				Value: '1',
			},
			&gear.Char{
				Value: '2',
			},
			&gear.Char{
				Value: '3',
			},
			&gear.Char{
				Value: '4',
			},
			&gear.Char{
				Value: '5',
			},
			&gear.Char{
				Value: '6',
			},
			&gear.Char{
				Value: '7',
			},
			&gear.Char{
				Value: '8',
			},
			&gear.Char{
				Value: '9',
			},
		},
	}

	number := gear.NewRule(
		"number",
		&gear.OneOrMore{
			Value: digit,
		},
	)

	factor := gear.NewRule(
		"factor",
		&gear.Choice{
			Value: []gear.Expression{
				&gear.NamedRule{
					Value: "number",
				},
				&gear.Sequence{
					Value: []gear.Expression{
						&gear.Char{
							Value: '(',
						},
						&gear.NamedRule{
							Value: "arithmetic_expression",
						},
						&gear.Char{
							Value: ')',
						},
					},
				},
			},
		},
	)

	term := gear.NewRule(
		"term",
		&gear.Sequence{
			Value: []gear.Expression{
				&gear.NamedRule{
					Value: "factor",
				},
				&gear.ZeroOrMore{
					Value: &gear.Sequence{
						Value: []gear.Expression{
							&gear.Char{
								Value: '*',
							},
							&gear.NamedRule{
								Value: "factor",
							},
						},
					},
				},
			},
		},
	)

	arithmeticExpression := gear.NewRule(
		"arithmetic_expression", &gear.Sequence{
			Value: []gear.Expression{
				&gear.NamedRule{
					Value: "term",
				},
				&gear.ZeroOrMore{
					Value: &gear.Sequence{
						Value: []gear.Expression{
							&gear.Char{
								Value: '+',
							},
							&gear.NamedRule{
								Value: "term",
							},
						},
					},
				},
			},
		},
	)

	g := gear.NewGrammar(
		gear.GrammarParam{
			Rules: []gear.Rule{
				number,
				factor,
				term,
				arithmeticExpression,
			},
		},
	)

	parser := gear.New(gear.ParserParam{
		Grammar: g,
	})

	input := "1+(2*3)"
	result, err := parser.Parse(input, "arithmetic_expression")

	if err != nil {
		panic(err)
	}

	fmt.Println(result.CST)
}
