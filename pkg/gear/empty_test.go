package gear

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmpty_Type(t *testing.T) {
	expr := Empty{}

	assert.Equal(t, EmptyExpression, expr.Type())
}

func TestEmpty_Evaluate(t *testing.T) {
	tests := map[string]struct {
		input          string
		expr           Expression
		expectedResult Result
		expectedError  error
	}{
		"match empty with empty input": {
			input:          "",
			expr:           &Empty{},
			expectedResult: Result{},
			expectedError:  nil,
		},
		"match empty with non-empty input": {
			input:          "abc",
			expr:           &Empty{},
			expectedResult: Result{},
			expectedError:  nil,
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
