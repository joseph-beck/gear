package gear

import (
	"testing"

	"github.com/joseph-beck/gear/pkg/errs"
	"github.com/stretchr/testify/assert"
)

func TestCharType(t *testing.T) {
	expr := Char{}

	assert.Equal(t, CharExpression, expr.Type())
}

func TestCharEvaluate(t *testing.T) {
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
				Next: 1,
				CST: CST{
					Value: "char",
					Children: []CST{
						{
							Value: "a",
						},
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
			expectedError:  errs.FailedToMatch,
		},
		"fail match empty input": {
			input: "",
			expr: &Char{
				Value: 'a',
			},
			expectedResult: Result{},
			expectedError:  errs.EndOfInput,
		},
		"match a with input ab": {
			input: "ab",
			expr: &Char{
				Value: 'a',
			},
			expectedResult: Result{
				CST: CST{
					Value: "char",
					Children: []CST{
						{
							Value: "a",
						},
					},
				},
			},
			expectedError: nil,
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
