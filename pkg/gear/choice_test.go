package gear

import (
	"testing"

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
				CST: CST{
					Value: "choice",
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
				CST: CST{
					Value: "choice",
					Children: []CST{
						{
							Value: "char",
							Children: []CST{
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
			expectedResult: Result{},
			expectedError:  err.FailedToMatch,
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
			expectedResult: Result{},
			expectedError:  err.EndOfInput,
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
				CST: CST{
					Value: "choice",
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
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			context := &Context{
				input: test.input,
			}
			output, err := test.expr.Evaluate(context)

			assert.Equal(t, test.expectedResult, output)

			assert.Equal(t, test.expectedError, err)
		})
	}
}
