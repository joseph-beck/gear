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
				Value: Char{
					Value: 'a',
				},
			},
			expectedResult: Result{
				Remaining: "",
				CST: cst.CST{
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
				Value: Char{
					Value: 'a',
				},
			},
			expectedResult: Result{
				Remaining: "b",
				CST: cst.CST{
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
				Value: Char{
					Value: 'a',
				},
			},
			expectedResult: Result{
				Remaining: "ba",
				CST: cst.CST{
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
				Value: Char{
					Value: 'a',
				},
			},
			expectedResult: Result{
				Remaining: "",
			},
			expectedError: err.EndOfInput,
		},
		"fail match a with input b": {
			input: "b",
			expr: OneOrMore{
				Value: Char{
					Value: 'a',
				},
			},
			expectedResult: Result{
				Remaining: "b",
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
