package main

import (
	"fmt"

	"github.com/joseph-beck/gear/pkg/gear"
)

func main() {
	g := gear.NewGrammar(
		gear.GrammarParam{
			Rules: []gear.Rule{
				gear.NewRule("digit", gear.NewChoice(
					[]gear.Expression{
						gear.NewChar('0'),
						gear.NewChar('1'),
						gear.NewChar('2'),
						gear.NewChar('3'),
						gear.NewChar('4'),
						gear.NewChar('5'),
						gear.NewChar('6'),
						gear.NewChar('7'),
						gear.NewChar('8'),
						gear.NewChar('9'),
					},
				)),
				gear.NewRule("number", gear.NewChoice(
					[]gear.Expression{
						gear.NewSequence(
							[]gear.Expression{
								gear.NewOptional(gear.NewNamedRule("number")),
								gear.NewNamedRule("digit"),
							},
						),
						gear.NewNamedRule("digit"),
					},
				)),
			},
		},
	)

	p := gear.New(gear.ParserParam{
		Grammar: g,
	})

	r, err := p.Parse("123123", "number")
	if err != nil {
		panic(err)
	}

	fmt.Println(r.CST)
}
