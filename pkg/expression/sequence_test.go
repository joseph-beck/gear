package expression

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSequenceType(t *testing.T) {
	expr := Sequence{}

	assert.Equal(t, SequenceExpression, expr.Type())
}
