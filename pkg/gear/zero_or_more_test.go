package gear

import (
	"testing"

	"github.com/joseph-beck/gear/pkg/errs"
	"github.com/stretchr/testify/assert"
)

func TestZeroOrMoreType(t *testing.T) {
	expr := ZeroOrMore{}

	assert.Equal(t, ZeroOrMoreExpression, expr.Type())
}

func TestZeroOrMoreEvaluate(t *testing.T) {
	tests := map[string]struct {
		input          string
		expr           Expression
		expectedResult Result
		expectedError  error
	}{
		"match a with input b": {
			input: "b",
			expr: &ZeroOrMore{
				Value: &Char{
					Value: 'a',
				},
			},
			expectedResult: Result{
				CST: CST{
					Value: "zero_or_more",
					Label: Label{
						Expression: true,
					},
				},
			},
			expectedError: nil,
		},
		"match a with input aaa": {
			input: "aaa",
			expr: &ZeroOrMore{
				Value: &Char{
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
							Label: Label{
								Expression: true,
							},
						},
						{
							Value: "char",
							Children: []CST{
								{
									Value: "a",
								},
							},
							Label: Label{
								Expression: true,
							},
						},
						{
							Value: "char",
							Children: []CST{
								{
									Value: "a",
								},
							},
							Label: Label{
								Expression: true,
							},
						},
					},
					Label: Label{
						Expression: true,
					},
				},
			},
			expectedError: nil,
		},
		"match a with input aaab": {
			input: "aaab",
			expr: &ZeroOrMore{
				Value: &Char{
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
							Label: Label{
								Expression: true,
							},
						},
						{
							Value: "char",
							Children: []CST{
								{
									Value: "a",
								},
							},
							Label: Label{
								Expression: true,
							},
						},
						{
							Value: "char",
							Children: []CST{
								{
									Value: "a",
								},
							},
							Label: Label{
								Expression: true,
							},
						},
					},
					Label: Label{
						Expression: true,
					},
				},
			},
			expectedError: nil,
		},
		"match a with input aaba": {
			input: "aaba",
			expr: &ZeroOrMore{
				Value: &Char{
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
							Label: Label{
								Expression: true,
							},
						},
						{
							Value: "char",
							Children: []CST{
								{
									Value: "a",
								},
							},
							Label: Label{
								Expression: true,
							},
						},
					},
					Label: Label{
						Expression: true,
					},
				},
			},
			expectedError: nil,
		},
		"match zero a's with empty input": {
			input: "",
			expr: &ZeroOrMore{
				Value: &Char{
					Value: 'a',
				},
			},
			expectedResult: Result{},
			expectedError:  errs.EndOfInput,
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
