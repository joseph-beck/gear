package gear

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnd_Type(t *testing.T) {
	expr := NewAnd(NewLiteral("test"))

	assert.Equal(t, AndExpression, expr.Type())
}

func TestAnd_Evaluate(t *testing.T) {
	tests := map[string]struct {
		input             string
		expr              Expression
		expectedResult    Result
		expectedRemaining string
		expectedError     error
	}{
		"and matches, literal does not consume any input": {
			input: "ab",
			expr:  NewAnd(NewLiteral("ab")),
			expectedResult: Result{
				CST: CST{
					value: "and",
					label: label{
						expression: true,
					},
				},
			},
			expectedRemaining: "ab",
			expectedError:     nil,
		},
		"and does not match, literal does not consume any input": {
			input:             "ab",
			expr:              NewAnd(NewLiteral("ac")),
			expectedResult:    Result{},
			expectedRemaining: "ab",
			expectedError:     ErrFailedToMatch,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := NewContext(test.input)

			output, err := test.expr.Evaluate(ctx)

			assert.Equal(t, test.expectedResult.CST, output.CST)

			assert.Equal(t, test.expectedRemaining, ctx.Remaining())

			assert.Equal(t, test.expectedError, err)
		})
	}
}
