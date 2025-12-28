package expression

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharType(t *testing.T) {
	expr := Char{}

	assert.Equal(t, CharExpression, expr.Type())
}
