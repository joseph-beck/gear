package gear

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNamedRule_Type(t *testing.T) {
	expr := NamedRule{}

	assert.Equal(t, NamedRuleExpression, expr.Type())
}

func TestNamedRule_Evaluate(t *testing.T) {
	tests := map[string]struct {
		input          string
		expr           Expression
		grammar        *Grammar
		expectedResult Result
		expectedError  error
	}{
		"match named rule_a with input a": {
			input: "a",
			expr: &NamedRule{
				Value: "rule_a",
			},
			grammar: func() *Grammar {
				g := &Grammar{}
				g.Add(NewRule("rule_a", &Char{
					Value: 'a',
				}))
				return g
			}(),
			expectedResult: Result{
				CST: CST{
					value: "rule_a",
					children: []CST{
						{
							value: "char",
							children: []CST{
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
			},
			expectedError: nil,
		},
		"error failed to match rule_a with input b": {
			input: "b",
			expr: &NamedRule{
				Value: "rule_a",
			},
			grammar: func() *Grammar {
				g := &Grammar{}
				g.Add(NewRule("rule_a", &Char{
					Value: 'a',
				}))
				return g
			}(),
			expectedResult: Result{},
			expectedError:  ErrFailedToMatch,
		},
		"match named sequence rule_a with input aaa": {
			input: "aaa",
			expr: &NamedRule{
				Value: "rule_a",
			},
			grammar: func() *Grammar {
				g := &Grammar{}
				g.Add(NewRule("rule_a", &ZeroOrMore{
					Value: &Char{
						Value: 'a',
					},
				}))
				return g
			}(),
			expectedResult: Result{
				CST: CST{
					value: "rule_a",
					children: []CST{
						{
							value: "zero_or_more",
							children: []CST{
								{
									value: "char",
									children: []CST{
										{
											value: "a",
										},
									},
									label: label{
										expression: true,
									},
								},
								{
									value: "char",
									children: []CST{
										{
											value: "a",
										},
									},
									label: label{
										expression: true,
									},
								},
								{
									value: "char",
									children: []CST{
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
					},
					label: label{
						expression: true,
					},
				},
			},
			expectedError: nil,
		},
		"error rule not found": {
			input: "a",
			expr: &NamedRule{
				Value: "rule_a",
			},
			grammar:        &Grammar{},
			expectedResult: Result{},
			expectedError:  ErrRuleNotFound,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := NewContext(test.input)
			ctx.grammar = test.grammar

			output, err := test.expr.Evaluate(ctx)

			assert.Equal(t, test.expectedResult.CST, output.CST)

			assert.Equal(t, test.expectedError, err)
		})
	}
}
