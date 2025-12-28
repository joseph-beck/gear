package gear

import (
	"testing"

	"github.com/joseph-beck/gear/pkg/err"
	"github.com/stretchr/testify/assert"
)

func TestZeroOrMoreType(t *testing.T) {
	expr := ZeroOrMore{}

	assert.Equal(t, ZeroOrMoreExpression, expr.Type())
}

func TestZeroOrMoreEvaluate(t *testing.T) {
	tests := map[string]struct {
		input          string
		expr           ZeroOrMore
		expectedResult Result
		expectedError  error
	}{
		"match a with input b": {
			input: "b",
			expr: ZeroOrMore{
				Value: Char{
					Value: 'a',
				},
			},
			expectedResult: Result{
				CST: CST{
					Value: "zero_or_more",
				},
			},
			expectedError: nil,
		},
		"match a with input aaa": {
			input: "aaa",
			expr: ZeroOrMore{
				Value: Char{
					Value: 'a',
				},
			},
			expectedResult: Result{
				CST: CST{
					Value: "zero_or_more",
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
									Value: "a",
								},
							},
						},
						{
							Value: "char",
							Children: []CST{
								{
									Value: "a",
								},
							},
						},
					},
				},
			},
			expectedError: nil,
		},
		"match a with input aaab": {
			input: "aaab",
			expr: ZeroOrMore{
				Value: Char{
					Value: 'a',
				},
			},
			expectedResult: Result{
				CST: CST{
					Value: "zero_or_more",
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
									Value: "a",
								},
							},
						},
						{
							Value: "char",
							Children: []CST{
								{
									Value: "a",
								},
							},
						},
					},
				},
			},
			expectedError: nil,
		},
		"match a with input aaba": {
			input: "aaba",
			expr: ZeroOrMore{
				Value: Char{
					Value: 'a',
				},
			},
			expectedResult: Result{
				CST: CST{
					Value: "zero_or_more",
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
									Value: "a",
								},
							},
						},
					},
				},
			},
			expectedError: nil,
		},
		"match zero a's with empty input": {
			input: "",
			expr: ZeroOrMore{
				Value: Char{
					Value: 'a',
				},
			},
			expectedResult: Result{},
			expectedError:  err.EndOfInput,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			context := &Context{
				input: test.input,
			}
			output, err := test.expr.Evaluate(context)

			assert.Equal(t, test.expectedResult, output)

			assert.Equal(t, test.expectedError, err)
		})
	}
}
