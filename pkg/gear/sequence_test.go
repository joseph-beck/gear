package gear

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSequence_Type(t *testing.T) {
	expr := Sequence{}

	assert.Equal(t, SequenceExpression, expr.Type())
}

func TestSequence_Evaluate(t *testing.T) {
	tests := map[string]struct {
		input          string
		expr           Expression
		expectedResult Result
		expectedError  error
	}{
		"match ab with input ab": {
			input: "ab",
			expr: &Sequence{
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
					value: "sequence",
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
		"match ab with input abc": {
			input: "abc",
			expr: &Sequence{
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
					value: "sequence",
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
		"fail match ab with input a": {
			input: "a",
			expr: &Sequence{
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
		"fail match ab with input b": {
			input: "b",
			expr: &Sequence{
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
		"fail match ab with empty input": {
			input: "",
			expr: &Sequence{
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
