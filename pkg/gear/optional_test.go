package gear

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptional_Type(t *testing.T) {
	expr := NewOptional(Letter)

	assert.Equal(t, OptionalExpression, expr.Type())
}

func TestOptional_Evaluate(t *testing.T) {
	tests := map[string]struct {
		input             string
		expr              Expression
		expectedResult    Result
		expectedRemaining string
		expectedError     error
	}{
		"optional matches, literal consumes input": {
			input: "ab",
			expr:  NewOptional(NewLiteral("ab")),
			expectedResult: Result{
				CST: CST{
					value: "optional",
					children: []CST{
						{
							value: "literal",
							children: []CST{
								{
									value: "ab",
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
			expectedRemaining: "",
			expectedError:     nil,
		},
		"optional does not match, literal does not consume input": {
			input: "ab",
			expr:  NewOptional(NewLiteral("ac")),
			expectedResult: Result{
				CST: CST{
					value:    "optional",
					children: []CST{},
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
