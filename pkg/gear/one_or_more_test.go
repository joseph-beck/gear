package gear

import (
	"testing"

	"github.com/joseph-beck/gear/pkg/errs"
	"github.com/stretchr/testify/assert"
)

func TestOneOrMoreType(t *testing.T) {
	expr := OneOrMore{}

	assert.Equal(t, OneOrMoreExpression, expr.Type())
}

func TestOneOrMoreEvaluate(t *testing.T) {
	tests := map[string]struct {
		input          string
		expr           Expression
		expectedResult Result
		expectedError  error
	}{
		"match a with input aaa": {
			input: "aaa",
			expr: &OneOrMore{
				Value: &Char{
					Value: 'a',
				},
			},
			expectedResult: Result{
				CST: CST{
					Value: "one_or_more",
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
			expr: &OneOrMore{
				Value: &Char{
					Value: 'a',
				},
			},
			expectedResult: Result{
				CST: CST{
					Value: "one_or_more",
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
			expr: &OneOrMore{
				Value: &Char{
					Value: 'a',
				},
			},
			expectedResult: Result{
				CST: CST{
					Value: "one_or_more",
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
		"fail match empty input": {
			input: "",
			expr: &OneOrMore{
				Value: &Char{
					Value: 'a',
				},
			},
			expectedResult: Result{},
			expectedError:  errs.EndOfInput,
		},
		"fail match a with input b": {
			input: "b",
			expr: &OneOrMore{
				Value: &Char{
					Value: 'a',
				},
			},
			expectedResult: Result{},
			expectedError:  errs.FailedToMatch,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := NewContext(test.input)

			output, err := test.expr.Evaluate(ctx, 0)

			assert.Equal(t, test.expectedResult.CST, output.CST)

			assert.Equal(t, test.expectedError, err)
		})
	}
}
