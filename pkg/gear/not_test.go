package gear

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNot_Type(t *testing.T) {
	not := NewNot(Letter)

	assert.Equal(t, NotExpression, not.Type())
}

func TestNot_Evaluate(t *testing.T) {
	tests := map[string]struct {
		input             string
		expr              Expression
		expectedResult    Result
		expectedRemaining string
		expectedError     error
	}{
		"not matches, literal does not consume any input, returns error": {
			input:             "ab",
			expr:              NewNot(NewLiteral("ab")),
			expectedResult:    Result{},
			expectedRemaining: "ab",
			expectedError:     ErrFailedToMatch,
		},
		"not does not match, literal does not consume any input, no error returned": {
			input: "ab",
			expr:  NewNot(NewLiteral("ac")),
			expectedResult: Result{
				CST: CST{
					value: "not",
					label: label{
						expression: true,
					},
				},
			},
			expectedRemaining: "ab",
			expectedError:     nil,
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
