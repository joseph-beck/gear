package parser

import (
	"testing"

	"github.com/joseph-beck/gear/pkg/cst"
	"github.com/joseph-beck/gear/pkg/err"
	"github.com/joseph-beck/gear/pkg/expression"
	"github.com/joseph-beck/gear/pkg/grammar"
	"github.com/joseph-beck/gear/pkg/rule"
	"github.com/stretchr/testify/assert"
)

func TestParserParse(t *testing.T) {
	tests := map[string]struct {
		input          string
		rule           string
		grammar        grammar.Grammar
		expectedResult expression.Result
		expectedError  error
	}{
		"error rule not found": {
			input:          "abc",
			rule:           "rule",
			grammar:        grammar.New(),
			expectedResult: expression.Result{},
			expectedError:  err.RuleNotFound,
		},
		"error failed to match": {
			input: "abc",
			rule:  "rule",
			grammar: func() grammar.Grammar {
				g := grammar.New()

				r := rule.New("rule", expression.Char{
					Value: 'x',
				})

				g.Add(r)

				return g
			}(),
			expectedResult: expression.Result{},
			expectedError:  err.FailedToMatch,
		},
		"error end of input": {
			input: "",
			rule:  "rule",
			grammar: func() grammar.Grammar {
				g := grammar.New()

				r := rule.New("rule", expression.Char{
					Value: 'x',
				})

				g.Add(r)

				return g
			}(),
			expectedResult: expression.Result{},
			expectedError:  err.EndOfInput,
		},
		"match char rule": {
			input: "a",
			rule:  "rule_a",
			grammar: func() grammar.Grammar {
				g := grammar.New()

				r := rule.New("rule_a", expression.Char{
					Value: 'a',
				})

				g.Add(r)

				return g
			}(),
			expectedResult: expression.Result{
				Remaining: "",
				CST: cst.CST{
					Value: "rule_a",
					Children: []cst.CST{
						{
							Value: "char",
							Children: []cst.CST{
								{
									Value: "a",
								},
							},
						},
					},
				},
			},
			expectedError: nil,
		},
		"match digit rule": {
			input: "123",
			rule:  "digit",
			grammar: func() grammar.Grammar {
				g := grammar.New()

				digit := rule.New("digit", expression.Choice{
					Value: []expression.Expression{
						expression.Char{
							Value: '0',
						},
						expression.Char{
							Value: '1',
						},
						expression.Char{
							Value: '2',
						},
						expression.Char{
							Value: '3',
						},
					},
				})

				g.Add(digit)

				return g
			}(),
			expectedResult: expression.Result{
				Remaining: "23",
				CST: cst.CST{
					Value: "digit",
					Children: []cst.CST{
						{
							Value: "choice",
							Children: []cst.CST{
								{
									Value: "char",
									Children: []cst.CST{
										{
											Value: "1",
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			parser := New(test.grammar)

			output, err := parser.Parse(test.input, test.rule)

			assert.Equal(t, test.expectedResult, output)

			assert.Equal(t, test.expectedError, err)
		})
	}
}
