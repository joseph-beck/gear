package main

import (
	"fmt"

	"github.com/joseph-beck/gear/pkg/gear"
)

func main() {
	g := gear.NewGrammar(
		gear.GrammarParam{
			Rules: []gear.Rule{
				gear.NewRule("expr", gear.NewChoice(
					[]gear.Expression{
						gear.NewSequence(
							[]gear.Expression{
								gear.NewNamedRule("expr"),
								gear.NewChar('+'),
								gear.NewNamedRule("expr"),
							},
						),
						gear.NewNamedRule("digit"),
					},
				)),
				gear.NewRule("digit", gear.NewChoice(
					[]gear.Expression{
						gear.NewChar('1'), gear.NewChar('2'), gear.NewChar('3'),
					},
				)),
			},
		},
	)

	p := gear.New(gear.ParserParam{
		Grammar: g,
	})

	r, err := p.Parse("1+2+3", "expr")

	if err != nil {
		panic(err)
	}

	fmt.Println(r.CST)
}
