package gear

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPackratKey(t *testing.T) {
	key := NewPackratKey("rule", 1)

	assert.Equal(t, "rule", key.rule)
	assert.Equal(t, uint(1), key.pos)
}

func TestNewPackratEntry(t *testing.T) {
	entry := NewPackratEntry(Result{
		CST: CST{
			value: "result",
		},
	}, nil)

	assert.Equal(t, "result", entry.result.CST.value)
	assert.Nil(t, entry.err)
}

func TestPackratEntry_Clone(t *testing.T) {
	entry := NewPackratEntry(Result{
		CST: CST{
			value: "result",
		},
	}, nil)

	cloned := entry.Clone()

	assert.Equal(t, "result", cloned.result.CST.value)
	assert.Nil(t, cloned.err)

	assert.NotSame(t, entry, cloned)
}

func TestNewPackrat(t *testing.T) {
	packrat := NewPackrat()

	assert.NotNil(t, packrat.memo)
	assert.Empty(t, packrat.memo)
}

func TestPackrat_Set(t *testing.T) {
}

func TestPackrat_Get(t *testing.T) {
}

func TestPackrat_Clone(t *testing.T) {
}
