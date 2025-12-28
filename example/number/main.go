package main

import (
	"fmt"

	"github.com/joseph-beck/gear/pkg/gear"
)

func main() {
	p := gear.New()

	g := gear.NewGrammar()

	digit := gear.NewRule("digit", gear.Choice{
		Value: []gear.Expression{
			gear.Char{Value: '0'},
			gear.Char{Value: '1'},
			gear.Char{Value: '2'},
			gear.Char{Value: '3'},
			gear.Char{Value: '4'},
			gear.Char{Value: '5'},
			gear.Char{Value: '6'},
			gear.Char{Value: '7'},
			gear.Char{Value: '8'},
			gear.Char{Value: '9'},
		},
	})
	number := gear.NewRule("number", gear.OneOrMore{
		Value: gear.NamedRule{
			Value:   "digit",
			Resolve: p.DefaultResolver,
		},
	})

	g.Add(digit)
	g.Add(number)

	p.SetGrammar(g)

	r, err := p.Parse("123", "number")
	if err != nil {
		panic(err)
	}

	fmt.Println(r.Remaining)
	fmt.Println(r.CST)
}
