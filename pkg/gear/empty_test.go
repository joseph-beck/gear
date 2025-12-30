package gear

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyType(t *testing.T) {
	expr := Empty{}

	assert.Equal(t, EmptyExpression, expr.Type())
}
