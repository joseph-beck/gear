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

			assert.Equal(t, test.expectedError, err)
		})
	}
}
