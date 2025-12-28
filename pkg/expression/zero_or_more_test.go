package expression

import (
	"testing"

	"github.com/joseph-beck/gear/pkg/cst"
	"github.com/joseph-beck/gear/pkg/err"
	"github.com/stretchr/testify/assert"
)

func TestZeroOrMoreType(t *testing.T) {
	expr := ZeroOrMore{}

	assert.Equal(t, ZeroOrMoreExpression, expr.Type())
}

func TestZeroOrMoreEvaluate(t *testing.T) {
	tests := map[string]struct {
		input          string
		expr           ZeroOrMore
		expectedResult Result
		expectedError  error
	}{
		"match zero a's": {
			input: "b",
			expr: ZeroOrMore{
				value: Char{
					value: 'a',
				},
			},
			expectedResult: Result{
				remaining: "b",
				cst: cst.CST{
					Value: "zero_or_more",
				},
			},
			expectedError: nil,
		},
		"match three a against aaa": {
			input: "aaa",
			expr: ZeroOrMore{
				value: Char{
					value: 'a',
				},
			},
			expectedResult: Result{
				remaining: "",
				cst: cst.CST{
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
			expectedError: nil,
		},
		"match three a against aaab": {
			input: "aaab",
			expr: ZeroOrMore{
				value: Char{
					value: 'a',
				},
			},
			expectedResult: Result{
				remaining: "b",
				cst: cst.CST{
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
			expectedError: nil,
		},
		"match three a against aaba": {
			input: "aaba",
			expr: ZeroOrMore{
				value: Char{
					value: 'a',
				},
			},
			expectedResult: Result{
				remaining: "ba",
				cst: cst.CST{
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
					},
				},
			},
			expectedError: nil,
		},
		"match zero a's with empty input": {
			input: "",
			expr: ZeroOrMore{
				value: Char{
					value: 'a',
				},
			},
			expectedResult: Result{
				remaining: "",
			},
			expectedError: err.EndOfInput,
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
