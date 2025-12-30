package gear

import (
	"testing"

	"github.com/joseph-beck/gear/pkg/errs"
	"github.com/stretchr/testify/assert"
)

func TestSequenceType(t *testing.T) {
	expr := Sequence{}

	assert.Equal(t, SequenceExpression, expr.Type())
}

func TestSequenceEvaluate(t *testing.T) {
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
					Value: "sequence",
					Children: []CST{
						{
							Value: "char",
							Children: []CST{
								{
									Value: "a",
								},
							},
						},
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
					Value: "sequence",
					Children: []CST{
						{
							Value: "char",
							Children: []CST{
								{
									Value: "a",
								},
							},
						},
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
			expectedError:  errs.EndOfInput,
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
			expectedError:  errs.FailedToMatch,
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
			expectedError:  errs.EndOfInput,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			context := NewContext(test.input)

			output, err := test.expr.Evaluate(context, 0)

			assert.Equal(t, test.expectedResult.CST, output.CST)

			assert.Equal(t, test.expectedError, err)
		})
	}
}
