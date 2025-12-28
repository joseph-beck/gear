package expression

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOneOrMoreType(t *testing.T) {
	expr := OneOrMore{}

	assert.Equal(t, OneOrMoreExpression, expr.Type())
}
