package gear

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChoice_Type(t *testing.T) {
	expr := Choice{}

	assert.Equal(t, ChoiceExpression, expr.Type())
}

func TestChoice_Evaluate(t *testing.T) {
	tests := map[string]struct {
		input          string
		expr           Expression
		expectedResult Result
		expectedError  error
	}{
		"match a or b with input a": {
			input: "a",
			expr: &Choice{
				Value: []Expression{
					&Char{
						Value: 'a',
					},
					&Char{
						Value: 'b',
					},
				},
			},
			expectedResult: Result{
				CST: CST{
					value: "choice",
					children: []CST{
						{
							value: "char",
							children: []CST{
								{
									value: "a",
								},
							},
							label: label{
								expression: true,
							},
						},
					},
					label: label{
						expression: true,
					},
				},
			},
			expectedError: nil,
		},
		"match a or b with input b": {
			input: "b",
			expr: &Choice{
				Value: []Expression{
					&Char{
						Value: 'a',
					},
					&Char{
						Value: 'b',
					},
				},
			},
			expectedResult: Result{
				CST: CST{
					value: "choice",
					children: []CST{
						{
							value: "char",
							children: []CST{
								{
									value: "b",
								},
							},
							label: label{
								expression: true,
							},
						},
					},
					label: label{
						expression: true,
					},
				},
			},
			expectedError: nil,
		},
		"fail match a or b with input c": {
			input: "c",
			expr: &Choice{
				Value: []Expression{
					&Char{
						Value: 'a',
					},
					&Char{
						Value: 'b',
					},
				},
			},
			expectedResult: Result{},
			expectedError:  ErrFailedToMatch,
		},
		"fail match a or b with empty input": {
			input: "",
			expr: &Choice{
				Value: []Expression{
					&Char{
						Value: 'a',
					},
					&Char{
						Value: 'b',
					},
				},
			},
			expectedResult: Result{},
			expectedError:  ErrEndOfInput,
		},
		"match a or b with input ab": {
			input: "ab",
			expr: &Choice{
				Value: []Expression{
					&Char{
						Value: 'a',
					},
					&Char{
						Value: 'b',
					},
				},
			},
			expectedResult: Result{
				CST: CST{
					value: "choice",
					children: []CST{
						{
							value: "char",
							children: []CST{
								{
									value: "a",
								},
							},
							label: label{
								expression: true,
							},
						},
					},
					label: label{
						expression: true,
					},
				},
			},
			expectedError: nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := NewContext(test.input)

			output, err := test.expr.Evaluate(ctx)

			assert.Equal(t, test.expectedResult.CST, output.CST)

			assert.Equal(t, test.expectedError, err)
		})
	}
}
