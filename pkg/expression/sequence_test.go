package expression

import (
	"testing"

	"github.com/joseph-beck/gear/pkg/cst"
	"github.com/joseph-beck/gear/pkg/err"
	"github.com/stretchr/testify/assert"
)

func TestSequenceType(t *testing.T) {
	expr := Sequence{}

	assert.Equal(t, SequenceExpression, expr.Type())
}

func TestSequenceEvaluate(t *testing.T) {
	tests := map[string]struct {
		input          string
		expr           Sequence
		expectedResult Result
		expectedError  error
	}{
		"match ab with input ab": {
			input: "ab",
			expr: Sequence{
				value: []Expression{
					Char{
						value: 'a',
					},
					Char{
						value: 'b',
					},
				},
			},
			expectedResult: Result{
				remaining: "",
				cst: cst.CST{
					Value: "sequence",
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
									Value: "b",
								},
							},
						},
					},
				},
			},
			expectedError: nil,
		},
		"match ab with input abc": {
			input: "abc",
			expr: Sequence{
				value: []Expression{
					Char{
						value: 'a',
					},
					Char{
						value: 'b',
					},
				},
			},
			expectedResult: Result{
				remaining: "c",
				cst: cst.CST{
					Value: "sequence",
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
									Value: "b",
								},
							},
						},
					},
				},
			},
			expectedError: nil,
		},
		"fail match ab with input a": {
			input: "a",
			expr: Sequence{
				value: []Expression{
					Char{
						value: 'a',
					},
					Char{
						value: 'b',
					},
				},
			},
			expectedResult: Result{
				remaining: "",
			},
			expectedError: err.EndOfInput,
		},
		"fail match ab with input b": {
			input: "b",
			expr: Sequence{
				value: []Expression{
					Char{
						value: 'a',
					},
					Char{
						value: 'b',
					},
				},
			},
			expectedResult: Result{
				remaining: "b",
			},
			expectedError: err.FailedToMatch,
		},
		"fail match ab with empty input": {
			input: "",
			expr: Sequence{
				value: []Expression{
					Char{
						value: 'a',
					},
					Char{
						value: 'b',
					},
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
