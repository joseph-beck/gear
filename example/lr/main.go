package main

import (
	"fmt"

	"github.com/joseph-beck/gear/pkg/gear"
)

func main() {
	p := gear.New()
	g := gear.NewGrammar()

	// digit ← '0' | '1' | ... | '9'
	digit := gear.NewRule("digit", &gear.Choice{
		Value: []gear.Expression{
			&gear.Char{Value: '0'},
			&gear.Char{Value: '1'},
			&gear.Char{Value: '2'},
			&gear.Char{Value: '3'},
			&gear.Char{Value: '4'},
			&gear.Char{Value: '5'},
			&gear.Char{Value: '6'},
			&gear.Char{Value: '7'},
			&gear.Char{Value: '8'},
			&gear.Char{Value: '9'},
		},
	})

	expr := gear.NewRule("expr", &gear.Choice{
		Value: []gear.Expression{
			&gear.Sequence{
				Value: []gear.Expression{
					&gear.NamedRule{
						Value: "expr",
					},
					&gear.Char{
						Value: '+',
					},
					&gear.NamedRule{
						Value: "digit",
					},
				},
			},
			&gear.NamedRule{
				Value: "digit",
			},
		},
	})

	g.Add(digit)
	g.Add(expr)

	p.SetGrammar(g)

	r, err := p.Parse("1+2+3", "expr")

	if err != nil {
		panic(err)
	}

	fmt.Println(r.CST)
}
