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
		"match a with input b": {
			input: "b",
			expr: ZeroOrMore{
				Value: Char{
					Value: 'a',
				},
			},
			expectedResult: Result{
				Remaining: "b",
				CST: cst.CST{
					Value: "zero_or_more",
				},
			},
			expectedError: nil,
		},
		"match a with input aaa": {
			input: "aaa",
			expr: ZeroOrMore{
				Value: Char{
					Value: 'a',
				},
			},
			expectedResult: Result{
				Remaining: "",
				CST: cst.CST{
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
		"match a with input aaab": {
			input: "aaab",
			expr: ZeroOrMore{
				Value: Char{
					Value: 'a',
				},
			},
			expectedResult: Result{
				Remaining: "b",
				CST: cst.CST{
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
		"match a with input aaba": {
			input: "aaba",
			expr: ZeroOrMore{
				Value: Char{
					Value: 'a',
				},
			},
			expectedResult: Result{
				Remaining: "ba",
				CST: cst.CST{
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
				Value: Char{
					Value: 'a',
				},
			},
			expectedResult: Result{
				Remaining: "",
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
