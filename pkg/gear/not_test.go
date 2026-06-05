package gear

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNot_Type(t *testing.T) {
	not := NewNot(Letter)

	assert.Equal(t, NotExpression, not.Type())
}

func TestNot_Evaluate(t *testing.T) {}
