package gear

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChar_Type(t *testing.T) {
	expr := Char{}

	assert.Equal(t, CharExpression, expr.Type())
}

func TestChar_Evaluate(t *testing.T) {
	tests := map[string]struct {
		input          string
		expr           Expression
		expectedResult Result
		expectedError  error
	}{
		"match a with a": {
			input: "a",
			expr: &Char{
				Value: 'a',
			},
			expectedResult: Result{
				CST: CST{
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
			expectedError: nil,
		},
		"fail match b with a": {
			input: "b",
			expr: &Char{
				Value: 'a',
			},
			expectedResult: Result{},
			expectedError:  ErrFailedToMatch,
		},
		"fail match empty input": {
			input: "",
			expr: &Char{
				Value: 'a',
			},
			expectedResult: Result{},
			expectedError:  ErrEndOfInput,
		},
		"match a with input ab": {
			input: "ab",
			expr: &Char{
				Value: 'a',
			},
			expectedResult: Result{
				CST: CST{
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
