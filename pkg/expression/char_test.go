package expression

import (
	"testing"

	"github.com/joseph-beck/gear/pkg/cst"
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
		expectError    bool
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
			expectError: false,
		},
		"fail match b with a": {
			input: "b",
			expr: Char{
				value: 'a',
			},
			expectedResult: Result{
				remaining: "b",
			},
			expectError: true,
		},
		"fail match empty input": {
			input: "",
			expr: Char{
				value: 'a',
			},
			expectedResult: Result{
				remaining: "",
			},
			expectError: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			output, err := test.expr.Evaluate(test.input)

			assert.Equal(t, test.expectedResult, output)

			if test.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
