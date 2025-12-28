package expression

import (
	"testing"

	"github.com/joseph-beck/gear/pkg/cst"
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
		expectedResult Result
		expectedError  error
	}{
		"match named rule_a with input a": {
			input: "a",
			expr: NamedRule{
				Value: "rule_a",
				Resolve: func(name string) (Expression, error) {
					if name == "rule_a" {
						return Char{
							Value: 'a',
						}, nil
					}
					return nil, nil
				},
			},
			expectedResult: Result{
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
		"no match named rule_a with input b": {
			input: "b",
			expr: NamedRule{
				Value: "rule_a",
				Resolve: func(name string) (Expression, error) {
					if name == "rule_a" {
						return Char{
							Value: 'a',
						}, nil
					}
					return nil, nil
				},
			},
			expectedResult: Result{
				Remaining: "b",
			},
			expectedError: err.FailedToMatch,
		},
		"nested named rule_b within named rule_a": {
			input: "b",
			expr: NamedRule{
				Value: "rule_a",
				Resolve: func(name string) (Expression, error) {
					if name == "rule_a" {
						return NamedRule{
							Value: "rule_b",
							Resolve: func(name string) (Expression, error) {
								if name == "rule_b" {
									return Char{
										Value: 'b',
									}, nil
								}
								return nil, nil
							},
						}, nil
					}
					return nil, nil
				},
			},
			expectedResult: Result{
				Remaining: "",
				CST: cst.CST{
					Value: "rule_a",
					Children: []cst.CST{
						{
							Value: "rule_b",
							Children: []cst.CST{
								{
									Value: "char",
									Children: []cst.CST{
										{
											Value: "b",
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
		"named sequence rule_a with input aaa": {
			input: "aaa",
			expr: NamedRule{
				Value: "rule_a",
				Resolve: func(name string) (Expression, error) {
					if name == "rule_a" {
						return ZeroOrMore{
							Value: Char{
								Value: 'a',
							},
						}, nil
					}
					return nil, nil
				},
			},
			expectedResult: Result{
				Remaining: "",
				CST: cst.CST{
					Value: "rule_a",
					Children: []cst.CST{
						{
							Value: "zero_or_more",
							Children: []cst.CST{
								{
									Value: "char",
									Children: []cst.CST{
										{
											Value: "a",
										},
									},
								},
								{
									Value: "char",
									Children: []cst.CST{
										{
											Value: "a",
										},
									},
								},
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
				},
			},
			expectedError: nil,
		},
		"resolve error": {
			input: "a",
			expr: NamedRule{
				Value: "rule_a",
				Resolve: func(name string) (Expression, error) {
					return nil, err.RuleNotFound
				},
			},
			expectedResult: Result{
				Remaining: "a",
			},
			expectedError: err.RuleNotFound,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			output, err := test.expr.Evaluate(test.input)

			assert.Equal(t, test.expectedResult, output)

			assert.Equal(t, test.expectedError, err)
		})
	}
}
