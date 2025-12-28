package main

import (
	"fmt"

	"github.com/joseph-beck/gear/pkg/expression"
	"github.com/joseph-beck/gear/pkg/grammar"
	"github.com/joseph-beck/gear/pkg/parser"
	"github.com/joseph-beck/gear/pkg/rule"
)

func main() {
	p := parser.New()

	g := grammar.New()

	digit := rule.New("digit", expression.Choice{
		Value: []expression.Expression{
			expression.Char{Value: '0'},
			expression.Char{Value: '1'},
			expression.Char{Value: '2'},
			expression.Char{Value: '3'},
			expression.Char{Value: '4'},
			expression.Char{Value: '5'},
			expression.Char{Value: '6'},
			expression.Char{Value: '7'},
			expression.Char{Value: '8'},
			expression.Char{Value: '9'},
		},
	})
	number := rule.New("number", expression.OneOrMore{
		Value: expression.NamedRule{
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
