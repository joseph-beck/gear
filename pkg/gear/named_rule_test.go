package gear

import (
	"testing"

	"github.com/joseph-beck/gear/pkg/err"
	"github.com/stretchr/testify/assert"
)

func TestNamedRuleType(t *testing.T) {
	expr := NamedRule{}

	assert.Equal(t, NamedRuleExpression, expr.Type())
}

func TestNamedRuleEvaluate(t *testing.T) {
	tests := map[string]struct {
		input          string
		expr           NamedRule
		grammar        *Grammar
		expectedResult Result
		expectedError  error
	}{
		"match named rule_a with input a": {
			input: "a",
			expr: NamedRule{
				Value: "rule_a",
			},
			grammar: func() *Grammar {
				g := &Grammar{}
				g.Add(NewRule("rule_a", Char{
					Value: 'a',
				}))
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
		"error failed to match rule_a with input b": {
			input: "b",
			expr: NamedRule{
				Value: "rule_a",
			},
			grammar: func() *Grammar {
				g := &Grammar{}
				g.Add(NewRule("rule_a", Char{
					Value: 'a',
				}))
				return g
			}(),
			expectedResult: Result{},
			expectedError:  err.FailedToMatch,
		},
		"named sequence rule_a with input aaa": {
			input: "aaa",
			expr: NamedRule{
				Value: "rule_a",
			},
			grammar: func() *Grammar {
				g := &Grammar{}
				g.Add(NewRule("rule_a", ZeroOrMore{
					Value: Char{
						Value: 'a',
					},
				}))
				return g
			}(),
			expectedResult: Result{
				CST: CST{
					Value: "rule_a",
					Children: []CST{
						{
							Value: "zero_or_more",
							Children: []CST{
								{
									Value: "char",
									Children: []CST{
										{
											Value: "a",
										},
									},
								},
								{
									Value: "char",
									Children: []CST{
										{
											Value: "a",
										},
									},
								},
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
				},
			},
			expectedError: nil,
		},
		"error rule not found": {
			input: "a",
			expr: NamedRule{
				Value: "rule_a",
			},
			grammar:        &Grammar{},
			expectedResult: Result{},
			expectedError:  err.RuleNotFound,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			context := &Context{
				input:   test.input,
				grammar: test.grammar,
			}
			output, err := test.expr.Evaluate(context)

			assert.Equal(t, test.expectedResult, output)

			assert.Equal(t, test.expectedError, err)
		})
	}
}
