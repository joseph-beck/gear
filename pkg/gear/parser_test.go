package gear

import (
	"testing"

	"github.com/joseph-beck/gear/pkg/err"
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
			expectedError:  err.RuleNotFound,
		},
		"error failed to match": {
			input: "abc",
			rule:  "rule",
			grammar: func() Grammar {
				g := NewGrammar()

				r := NewRule("rule", Char{
					Value: 'x',
				})

				g.Add(r)

				return g
			}(),
			expectedResult: Result{},
			expectedError:  err.FailedToMatch,
		},
		"error end of input": {
			input: "",
			rule:  "rule",
			grammar: func() Grammar {
				g := NewGrammar()

				r := NewRule("rule", Char{
					Value: 'x',
				})

				g.Add(r)

				return g
			}(),
			expectedResult: Result{},
			expectedError:  err.EndOfInput,
		},
		"match char rule": {
			input: "a",
			rule:  "rule_a",
			grammar: func() Grammar {
				g := NewGrammar()

				r := NewRule("rule_a", Char{
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
						},
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

				digit := NewRule("digit", Choice{
					Value: []Expression{
						Char{
							Value: '0',
						},
						Char{
							Value: '1',
						},
						Char{
							Value: '2',
						},
						Char{
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
