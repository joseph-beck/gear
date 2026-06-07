package peg

import (
	"testing"

	"github.com/joseph-beck/gear/pkg/gear"
	"github.com/stretchr/testify/assert"
)

func TestDirectLeftRecursion(t *testing.T) {
	tests := map[string]struct {
		input          string
		rule           string
		grammar        gear.Grammar
		expectedResult gear.ParserResult
		expectedError  error
	}{
		"simple direct left recursion": {
			input: "1+1+1",
			rule:  "expr",
			grammar: gear.NewGrammar(
				gear.GrammarParam{
					Rules: []gear.Rule{
						gear.NewRule("expr", gear.NewChoice(
							[]gear.Expression{
								gear.NewSequence(
									[]gear.Expression{
										gear.NewNamedRule("expr"),
										gear.NewChar('+'),
										gear.NewNamedRule("expr"),
									},
								),
								gear.NewNamedRule("digit"),
							},
						)),
						gear.NewRule("digit", gear.NewChoice(
							[]gear.Expression{
								gear.NewChar('1'),
							},
						)),
					},
				},
			),
			expectedResult: gear.ParserResult{
				CST: gear.NewCST(gear.CSTParam{
					Value: "expr",
					Children: []gear.CST{
						gear.NewCST(gear.CSTParam{
							Value:    "choice",
							Children: []gear.CST{},
							Label: gear.NewLabel(gear.LabelParam{
								Expression: true,
							}),
						}),
					},
					Label: gear.NewLabel(gear.LabelParam{
						Expression: true,
					}),
				}),
				Remaining: "",
			},
			expectedError: nil,
		},
		"multiple direct left recursion": {
			input: "1+1+1",
			rule:  "expr",
			grammar: gear.NewGrammar(
				gear.GrammarParam{
					Rules: []gear.Rule{
						gear.NewRule("expr", gear.NewChoice([]gear.Expression{
							gear.NewSequence(
								[]gear.Expression{
									gear.NewNamedRule("expr"),
									gear.NewChar('*'),
									gear.NewNamedRule("expr"),
								},
							),
							gear.NewSequence(
								[]gear.Expression{
									gear.NewNamedRule("expr"),
									gear.NewChar('+'),
									gear.NewNamedRule("expr"),
								},
							),
							gear.NewNamedRule("digit"),
						})),
						gear.NewRule("digit", gear.NewChoice(
							[]gear.Expression{
								gear.NewChar('1'),
							},
						)),
					},
				},
			),
			expectedResult: gear.ParserResult{
				CST: gear.NewCST(gear.CSTParam{
					Value: "expr",
					Children: []gear.CST{
						gear.NewCST(gear.CSTParam{
							Value:    "choice",
							Children: []gear.CST{},
							Label: gear.NewLabel(gear.LabelParam{
								Expression: true,
							}),
						}),
					},
					Label: gear.NewLabel(gear.LabelParam{
						Expression: true,
					}),
				}),
				Remaining: "",
			},
			expectedError: nil,
		},
		"rule not found": {
			input:   "1+1",
			rule:    "expr",
			grammar: gear.NewGrammar(),
			expectedResult: gear.ParserResult{
				Remaining: "1+1",
			},
			expectedError: gear.ErrRuleNotFound,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			parser := gear.New(gear.ParserParam{
				Grammar: test.grammar,
			})

			result, err := parser.Parse(test.input, test.rule)

			// assert.Equal(t, test.expectedResult.CST, output.CST)

			assert.Equal(t, test.expectedResult.Remaining, result.Remaining)

			assert.Equal(t, test.expectedError, err)
		})
	}
}

func TestIndirectLeftRecursion(t *testing.T) {
	tests := map[string]struct {
		input          string
		rule           string
		grammar        gear.Grammar
		expectedResult gear.ParserResult
		expectedError  error
	}{
		"simple indirect left recursion": {
			expectedError: gear.ErrRuleNotFound,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			parser := gear.New(gear.ParserParam{
				Grammar: test.grammar,
			})

			output, err := parser.Parse(test.input, test.rule)

			assert.Equal(t, test.expectedResult.CST, output.CST)

			assert.Equal(t, test.expectedResult.Remaining, output.Remaining)

			assert.Equal(t, test.expectedError, err)
		})
	}
}
