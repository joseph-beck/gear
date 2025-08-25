package reflections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpectType(t *testing.T) {
	x := "string"
	y := ExpectType[string](x)

	assert.Equal(t, "string", y)

	assert.Panics(t, func() {
		ExpectType[int](x)
	})
}
