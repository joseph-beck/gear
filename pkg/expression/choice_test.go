package expression

import (
	"testing"

	"github.com/joseph-beck/gear/pkg/cst"
	"github.com/joseph-beck/gear/pkg/err"
	"github.com/stretchr/testify/assert"
)

func TestChoiceType(t *testing.T) {
	expr := Choice{}

	assert.Equal(t, ChoiceExpression, expr.Type())
}

func TestChoiceEvaluate(t *testing.T) {
	tests := map[string]struct {
		input          string
		expr           Choice
		expectedResult Result
		expectedError  error
	}{
		"match a or b with input a": {
			input: "a",
			expr: Choice{
				Value: []Expression{
					Char{
						Value: 'a',
					},
					Char{
						Value: 'b',
					},
				},
			},
			expectedResult: Result{
				Remaining: "",
				CST: cst.CST{
					Value: "choice",
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
		"match a or b with input b": {
			input: "b",
			expr: Choice{
				Value: []Expression{
					Char{
						Value: 'a',
					},
					Char{
						Value: 'b',
					},
				},
			},
			expectedResult: Result{
				Remaining: "",
				CST: cst.CST{
					Value: "choice",
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
			expectedError: nil,
		},
		"fail match a or b with input c": {
			input: "c",
			expr: Choice{
				Value: []Expression{
					Char{
						Value: 'a',
					},
					Char{
						Value: 'b',
					},
				},
			},
			expectedResult: Result{
				Remaining: "c",
			},
			expectedError: err.FailedToMatch,
		},
		"fail match a or b with empty input": {
			input: "",
			expr: Choice{
				Value: []Expression{
					Char{
						Value: 'a',
					},
					Char{
						Value: 'b',
					},
				},
			},
			expectedResult: Result{
				Remaining: "",
			},
			expectedError: err.EndOfInput,
		},
		"match a or b with input ab": {
			input: "ab",
			expr: Choice{
				Value: []Expression{
					Char{
						Value: 'a',
					},
					Char{
						Value: 'b',
					},
				},
			},
			expectedResult: Result{
				Remaining: "b",
				CST: cst.CST{
					Value: "choice",
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
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			output, err := test.expr.Evaluate(test.input)

			assert.Equal(t, test.expectedResult, output)

			assert.Equal(t, test.expectedError, err)
		})
	}
}
