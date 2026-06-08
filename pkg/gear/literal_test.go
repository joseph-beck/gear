package gear

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLiteral_Type(t *testing.T) {
	expr := Literal{}

	assert.Equal(t, LiteralExpression, expr.Type())
}

func TestLiteral_Evaluate(t *testing.T) {
	tests := map[string]struct {
		input          string
		expr           Expression
		expectedResult Result
		expectedError  error
	}{
		"match abc with abc": {
			input: "abc",
			expr: &Literal{
				Value: "abc",
			},
			expectedResult: Result{
				CST: CST{
					value: "literal",
					children: []CST{
						{
							value: "abc",
						},
					},
					label: label{
						expression: true,
					},
				},
			},
			expectedError: nil,
		},
		"fail match abc with ab with end of input": {
			input: "ab",
			expr: &Literal{
				Value: "abc",
			},
			expectedResult: Result{},
			expectedError:  ErrEndOfInput,
		},
		"fail match abc with abd with failed to match": {
			input: "abd",
			expr: &Literal{
				Value: "abc",
			},
			expectedResult: Result{},
			expectedError:  ErrFailedToMatch,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := NewContext(test.input)

			result, err := test.expr.Evaluate(ctx)

			assert.Equal(t, test.expectedResult, result)
			assert.Equal(t, test.expectedError, err)
		})
	}
}
