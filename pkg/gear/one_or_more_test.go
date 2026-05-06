package gear

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOneOrMore_Type(t *testing.T) {
	expr := OneOrMore{}

	assert.Equal(t, OneOrMoreExpression, expr.Type())
}

func TestOneOrMore_Evaluate(t *testing.T) {
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
					value: "one_or_more",
					children: []CST{
						{
							value: "char",
							children: []CST{
								{
									value: "a",
								},
							},
							label: label{
								expression: true,
							},
						},
						{
							value: "char",
							children: []CST{
								{
									value: "a",
								},
							},
							label: label{
								expression: true,
							},
						},
						{
							value: "char",
							children: []CST{
								{
									value: "a",
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
					value: "one_or_more",
					children: []CST{
						{
							value: "char",
							children: []CST{
								{
									value: "a",
								},
							},
							label: label{
								expression: true,
							},
						},
						{
							value: "char",
							children: []CST{
								{
									value: "a",
								},
							},
							label: label{
								expression: true,
							},
						},
						{
							value: "char",
							children: []CST{
								{
									value: "a",
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
					value: "one_or_more",
					children: []CST{
						{
							value: "char",
							children: []CST{
								{
									value: "a",
								},
							},
							label: label{
								expression: true,
							},
						},
						{
							value: "char",
							children: []CST{
								{
									value: "a",
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
			expectedError:  ErrEndOfInput,
		},
		"fail match a with input b": {
			input: "b",
			expr: &OneOrMore{
				Value: &Char{
					Value: 'a',
				},
			},
			expectedResult: Result{},
			expectedError:  ErrFailedToMatch,
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
