package expression

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNamedRuleType(t *testing.T) {
	expr := NamedRule{}

	assert.Equal(t, NamedRuleExpression, expr.Type())
}
