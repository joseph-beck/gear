package expression

import (
	"testing"

	"github.com/joseph-beck/gear/pkg/cst"
	"github.com/joseph-beck/gear/pkg/err"
	"github.com/stretchr/testify/assert"
)

func TestCharType(t *testing.T) {
	expr := Char{}

	assert.Equal(t, CharExpression, expr.Type())
}

func TestCharEvaluate(t *testing.T) {
	tests := map[string]struct {
		input          string
		expr           Char
		expectedResult Result
		expectedError  error
	}{
		"match a with a": {
			input: "a",
			expr: Char{
				value: 'a',
			},
			expectedResult: Result{
				remaining: "",
				cst: cst.CST{
					Value: "char",
					Children: []cst.CST{
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
			expr: Char{
				value: 'a',
			},
			expectedResult: Result{
				remaining: "b",
			},
			expectedError: err.FailedToMatch,
		},
		"fail match empty input": {
			input: "",
			expr: Char{
				value: 'a',
			},
			expectedResult: Result{
				remaining: "",
			},
			expectedError: err.EndOfInput,
		},
		"match a with input ab": {
			input: "ab",
			expr: Char{
				value: 'a',
			},
			expectedResult: Result{
				remaining: "b",
				cst: cst.CST{
					Value: "char",
					Children: []cst.CST{
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
			output, err := test.expr.Evaluate(test.input)

			assert.Equal(t, test.expectedResult, output)

			assert.Equal(t, test.expectedError, err)
		})
	}
}
