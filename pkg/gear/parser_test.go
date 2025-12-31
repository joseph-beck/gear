package gear

import (
	"testing"

	"github.com/joseph-beck/gear/pkg/errs"
	"github.com/stretchr/testify/assert"
)

func TestParserParse(t *testing.T) {
	tests := map[string]struct {
		input          string
		rule           string
		grammar        Grammar
		expectedResult ParserResult
		expectedError  error
	}{
		"error rule not found": {
			input:          "abc",
			rule:           "rule",
			grammar:        NewGrammar(),
			expectedResult: ParserResult{},
			expectedError:  errs.RuleNotFound,
		},
		"error failed to match": {
			input: "abc",
			rule:  "rule",
			grammar: func() Grammar {
				g := NewGrammar()

				r := NewRule("rule", &Char{
					Value: 'x',
				})

				g.Add(r)

				return g
			}(),
			expectedResult: ParserResult{},
			expectedError:  errs.FailedToMatch,
		},
		"error end of input": {
			input: "",
			rule:  "rule",
			grammar: func() Grammar {
				g := NewGrammar()

				r := NewRule("rule", &Char{
					Value: 'x',
				})

				g.Add(r)

				return g
			}(),
			expectedResult: ParserResult{},
			expectedError:  errs.EndOfInput,
		},
		"match char rule": {
			input: "a",
			rule:  "rule_a",
			grammar: func() Grammar {
				g := NewGrammar()

				r := NewRule("rule_a", &Char{
					Value: 'a',
				})

				g.Add(r)

				return g
			}(),
			expectedResult: ParserResult{
				CST: cst{
					value: "rule_a",
					children: []cst{
						{
							value: "char",
							children: []cst{
								{
									value: "a",
								},
							},
							label: label{
								expression: true,
							},
						},
					},
					label: label{
						expression: true,
					},
				},
				Remaining: "",
			},
			expectedError: nil,
		},
		"match char rule with remaining b": {
			input: "ab",
			rule:  "rule_a",
			grammar: func() Grammar {
				g := NewGrammar()

				r := NewRule("rule_a", &Char{
					Value: 'a',
				})

				g.Add(r)

				return g
			}(),
			expectedResult: ParserResult{
				CST: cst{
					value: "rule_a",
					children: []cst{
						{
							value: "char",
							children: []cst{
								{
									value: "a",
								},
							},
							label: label{
								expression: true,
							},
						},
					},
					label: label{
						expression: true,
					},
				},
				Remaining: "b",
			},
			expectedError: nil,
		},
		"match digit rule": {
			input: "123",
			rule:  "digit",
			grammar: func() Grammar {
				g := NewGrammar()

				digit := NewRule("digit", &Choice{
					Value: []Expression{
						&Char{
							Value: '0',
						},
						&Char{
							Value: '1',
						},
						&Char{
							Value: '2',
						},
						&Char{
							Value: '3',
						},
					},
				})

				g.Add(digit)

				return g
			}(),
			expectedResult: ParserResult{
				CST: cst{
					value: "digit",
					children: []cst{
						{
							value: "choice",
							children: []cst{
								{
									value: "char",
									children: []cst{
										{
											value: "1",
										},
									},
									label: label{
										expression: true,
									},
								},
							},
							label: label{
								expression: true,
							},
						},
					},
					label: label{
						expression: true,
					},
				},
				Remaining: "23",
			},
			expectedError: nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			parser := New(ParserParam{
				Grammar: test.grammar,
			})

			output, err := parser.Parse(test.input, test.rule)

			assert.Equal(t, test.expectedResult.CST, output.CST)

			assert.Equal(t, test.expectedError, err)
		})
	}
}
