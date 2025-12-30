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
		expectedResult Result
		expectedError  error
	}{
		"error rule not found": {
			input:          "abc",
			rule:           "rule",
			grammar:        NewGrammar(),
			expectedResult: Result{},
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
			expectedResult: Result{},
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
			expectedResult: Result{},
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
			expectedResult: Result{
				CST: CST{
					Value: "rule_a",
					Children: []CST{
						{
							Value: "char",
							Children: []CST{
								{
									Value: "a",
								},
							},
							Label: Label{
								Expression: true,
							},
						},
					},
					Label: Label{
						Expression: true,
					},
				},
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
			expectedResult: Result{
				CST: CST{
					Value: "digit",
					Children: []CST{
						{
							Value: "choice",
							Children: []CST{
								{
									Value: "char",
									Children: []CST{
										{
											Value: "1",
										},
									},
									Label: Label{
										Expression: true,
									},
								},
							},
							Label: Label{
								Expression: true,
							},
						},
					},
					Label: Label{
						Expression: true,
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

			assert.Equal(t, test.expectedResult.CST, output.CST)

			assert.Equal(t, test.expectedError, err)
		})
	}
}
