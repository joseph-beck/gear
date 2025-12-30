package gear

import (
	"testing"

	"github.com/joseph-beck/gear/pkg/errs"
	"github.com/stretchr/testify/assert"
)

func TestNamedRuleType(t *testing.T) {
	expr := NamedRule{}

	assert.Equal(t, NamedRuleExpression, expr.Type())
}

func TestNamedRuleEvaluate(t *testing.T) {
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
			expectedError:  errs.FailedToMatch,
		},
		"named sequence rule_a with input aaa": {
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
									Label: Label{
										Expression: true,
									},
								},
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
					Label: Label{
						Expression: true,
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
			expectedError:  errs.RuleNotFound,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := NewContext(test.input)
			ctx.grammar = test.grammar

			output, err := test.expr.Evaluate(ctx, 0)

			assert.Equal(t, test.expectedResult.CST, output.CST)

			assert.Equal(t, test.expectedError, err)
		})
	}
}
