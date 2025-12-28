package expression

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChoiceType(t *testing.T) {
	expr := Choice{}

	assert.Equal(t, ChoiceExpression, expr.Type())
}
