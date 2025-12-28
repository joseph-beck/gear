package expression

import (
	"testing"

	"github.com/joseph-beck/gear/pkg/cst"
	"github.com/joseph-beck/gear/pkg/err"
	"github.com/stretchr/testify/assert"
)

func TestOneOrMoreType(t *testing.T) {
	expr := OneOrMore{}

	assert.Equal(t, OneOrMoreExpression, expr.Type())
}

func TestOneOrMoreEvaluate(t *testing.T) {
	tests := map[string]struct {
		input          string
		expr           OneOrMore
		expectedResult Result
		expectedError  error
	}{
		"match a with input aaa": {
			input: "aaa",
			expr: OneOrMore{
				value: Char{
					value: 'a',
				},
			},
			expectedResult: Result{
				remaining: "",
				cst: cst.CST{
					Value: "one_or_more",
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
			expectedError: nil,
		},
		"match a with input aaab": {
			input: "aaab",
			expr: OneOrMore{
				value: Char{
					value: 'a',
				},
			},
			expectedResult: Result{
				remaining: "b",
				cst: cst.CST{
					Value: "one_or_more",
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
			expectedError: nil,
		},
		"match a with input aaba": {
			input: "aaba",
			expr: OneOrMore{
				value: Char{
					value: 'a',
				},
			},
			expectedResult: Result{
				remaining: "ba",
				cst: cst.CST{
					Value: "one_or_more",
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
					},
				},
			},
			expectedError: nil,
		},
		"fail match empty input": {
			input: "",
			expr: OneOrMore{
				value: Char{
					value: 'a',
				},
			},
			expectedResult: Result{
				remaining: "",
			},
			expectedError: err.EndOfInput,
		},
		"fail match a with input b": {
			input: "b",
			expr: OneOrMore{
				value: Char{
					value: 'a',
				},
			},
			expectedResult: Result{
				remaining: "b",
			},
			expectedError: err.FailedToMatch,
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
