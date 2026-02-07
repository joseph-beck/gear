package gear

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParserParse(t *testing.T) {
	tests := map[string]struct {
		input          string
		rule           string
		grammar        Grammar
		expectedResult ParserResult
		expectedError  error
	}{
		"error rule not found": {
			input:          "abc",
			rule:           "rule",
			grammar:        NewGrammar(),
			expectedResult: ParserResult{},
			expectedError:  ErrRuleNotFound,
		},
		"error failed to match": {
			input: "abc",
			rule:  "rule",
			grammar: func() Grammar {
				g := NewGrammar()

				r := NewRule("rule", &Char{
					Value: 'x',
				})

				g.Add(r)

				return g
			}(),
			expectedResult: ParserResult{},
			expectedError:  ErrFailedToMatch,
		},
		"error end of input": {
			input: "",
			rule:  "rule",
			grammar: func() Grammar {
				g := NewGrammar()

				r := NewRule("rule", &Char{
					Value: 'x',
				})

				g.Add(r)

				return g
			}(),
			expectedResult: ParserResult{},
			expectedError:  ErrEndOfInput,
		},
		"match char rule": {
			input: "a",
			rule:  "rule_a",
			grammar: func() Grammar {
				g := NewGrammar()

				r := NewRule("rule_a", &Char{
					Value: 'a',
				})

				g.Add(r)

				return g
			}(),
			expectedResult: ParserResult{
				CST: cst{
					value: "rule_a",
					children: []cst{
						{
							value: "char",
							children: []cst{
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
				Remaining: "",
			},
			expectedError: nil,
		},
		"match char rule with remaining b": {
			input: "ab",
			rule:  "rule_a",
			grammar: func() Grammar {
				g := NewGrammar()

				r := NewRule("rule_a", &Char{
					Value: 'a',
				})

				g.Add(r)

				return g
			}(),
			expectedResult: ParserResult{
				CST: cst{
					value: "rule_a",
					children: []cst{
						{
							value: "char",
							children: []cst{
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
				Remaining: "b",
			},
			expectedError: nil,
		},
		"match digit rule": {
			input: "123",
			rule:  "digit",
			grammar: func() Grammar {
				g := NewGrammar()

				digit := NewRule("digit", &Choice{
					Value: []Expression{
						&Char{
							Value: '0',
						},
						&Char{
							Value: '1',
						},
						&Char{
							Value: '2',
						},
						&Char{
							Value: '3',
						},
					},
				})

				g.Add(digit)

				return g
			}(),
			expectedResult: ParserResult{
				CST: cst{
					value: "digit",
					children: []cst{
						{
							value: "choice",
							children: []cst{
								{
									value: "char",
									children: []cst{
										{
											value: "1",
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
					label: label{
						expression: true,
					},
				},
				Remaining: "23",
			},
			expectedError: nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			parser := New(ParserParam{
				Grammar: test.grammar,
			})

			output, err := parser.Parse(test.input, test.rule)

			assert.Equal(t, test.expectedResult.CST, output.CST)

			assert.Equal(t, test.expectedError, err)
		})
	}
}

func BenchmarkParserParseDirectLeftRecursion(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		g := NewGrammar()

		number := NewRule("number", &Choice{
			Value: []Expression{
				&Char{
					Value: '1',
				},
				&Char{
					Value: '2',
				},
				&Char{
					Value: '3',
				},
			},
		})

		expr := NewRule("expr", &Choice{
			Value: []Expression{
				&Sequence{
					Value: []Expression{
						&NamedRule{
							Value: "expr",
						},
						&Char{
							Value: '+',
						},
						&NamedRule{
							Value: "number",
						},
					},
				},
				&NamedRule{
					Value: "number",
				},
			},
		})

		g.Add(number)
		g.Add(expr)

		parser := New(ParserParam{
			Grammar: g,
		})

		_, _ = parser.Parse("1+2+3", "expr")
	}
}

func BenchmarkParserParseIndirectLeftRecursion(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		g := NewGrammar()

		number := NewRule("number", &Choice{
			Value: []Expression{
				&Char{
					Value: '1',
				},
				&Char{
					Value: '2',
				},
				&Char{
					Value: '3',
				},
			},
		})

		x := NewRule("x", &NamedRule{
			Value: "expr",
		})

		expr := NewRule("expr", &Choice{
			Value: []Expression{
				&Sequence{
					Value: []Expression{
						&NamedRule{
							Value: "x",
						},
						&Char{
							Value: '+',
						},
						&NamedRule{
							Value: "number",
						},
					},
				},
				&NamedRule{
					Value: "number",
				},
			},
		})

		g.Add(number)
		g.Add(x)
		g.Add(expr)

		parser := New(ParserParam{
			Grammar: g,
		})

		_, _ = parser.Parse("1+2+3", "expr")
	}
}

func BenchmarkParserParseArithmeticExpression(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		number := NewRule(
			"number",
			&OneOrMore{
				Value: Digit,
			},
		)

		factor := NewRule(
			"factor",
			&Choice{
				Value: []Expression{
					&NamedRule{
						Value: "number",
					},
					&Sequence{
						Value: []Expression{
							&Char{
								Value: '(',
							},
							&NamedRule{
								Value: "arithmetic_expression",
							},
							&Char{
								Value: ')',
							},
						},
					},
				},
			},
		)

		term := NewRule(
			"term",
			&Sequence{
				Value: []Expression{
					&NamedRule{
						Value: "factor",
					},
					&ZeroOrMore{
						Value: &Sequence{
							Value: []Expression{
								&Char{
									Value: '*',
								},
								&NamedRule{
									Value: "factor",
								},
							},
						},
					},
				},
			},
		)

		arithmeticExpression := NewRule(
			"arithmetic_expression", &Sequence{
				Value: []Expression{
					&NamedRule{
						Value: "term",
					},
					&ZeroOrMore{
						Value: &Sequence{
							Value: []Expression{
								&Char{
									Value: '+',
								},
								&NamedRule{
									Value: "term",
								},
							},
						},
					},
				},
			},
		)

		g := NewGrammar()

		g.Add(number)
		g.Add(factor)
		g.Add(term)
		g.Add(arithmeticExpression)

		parser := New(ParserParam{
			Grammar: g,
		})

		_, _ = parser.Parse("1+(2*3)", "arithmetic_expression")
	}
}

func BenchmarkNumber(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		number := NewRule(
			"number",
			&OneOrMore{
				Value: Digit,
			},
		)

		g := NewGrammar()

		g.Add(number)

		parser := New(ParserParam{
			Grammar: g,
		})

		_, _ = parser.Parse("100000000", "number")
	}
}

func BenchmarkWord(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		word := NewRule(
			"word",
			&OneOrMore{
				Value: Letter,
			},
		)

		g := NewGrammar()

		g.Add(word)

		parser := New(ParserParam{
			Grammar: g,
		})

		_, _ = parser.Parse("helloworld", "word")
	}
}

func BenchmarkSentence(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		sentence := NewRule(
			"sentence",
			&OneOrMore{
				Value: &Sequence{
					Value: []Expression{
						Letter,
						&Choice{
							Value: []Expression{
								&Char{
									Value: ' ',
								},
								&Char{
									Value: ',',
								},
								&Char{
									Value: '.',
								},
							},
						},
					},
				},
			},
		)

		g := NewGrammar()

		g.Add(sentence)

		parser := New(ParserParam{
			Grammar: g,
		})

		_, _ = parser.Parse(
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur at tellus eros. Donec semper tempus facilisis. Mauris elementum non quam in aliquet. Nulla massa ligula, vehicula in aliquam varius, venenatis ut quam. Nam feugiat enim nec euismod pretium. Proin accumsan iaculis dolor id dignissim. Sed mollis sit amet mi a tincidunt. Proin at felis eget nisl cursus venenatis molestie nec augue. Morbi ut mauris nec risus gravida fringilla vehicula vitae orci. Pellentesque accumsan et mauris id ultricies. Donec tincidunt lorem tortor, id imperdiet lorem imperdiet ut. Pellentesque non quam at sem suscipit vestibulum eget non risus. Mauris auctor volutpat leo, vel interdum velit posuere sed. Donec ultricies justo non leo vestibulum, eu vehicula arcu commodo.",
			"sentence",
		)
	}
}
