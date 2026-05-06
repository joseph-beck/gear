package gear

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHistory(t *testing.T) {
	history := NewHistory()

	assert.Equal(t, history, make(History, 0))
}

func TestHistory_Preserve(t *testing.T) {
	history := NewHistory()
	artifact := NewArtifact("rule", 0, 0)
	history.Preserve(artifact)

	assert.Equal(t, history, History{artifact})
}

func TestHistory_Clear(t *testing.T) {
	history := NewHistory()
	artifact := NewArtifact("rule", 0, 0)
	history.Preserve(artifact)
	history.Clear()

	assert.Equal(t, history, make(History, 0))
}

func TestHistory_Prod(t *testing.T) {
	history := NewHistory()
	artifact := NewArtifact("rule", 0, 0)
	history.Preserve(artifact)

	assert.True(t, history.Prod(artifact))
	assert.False(t, history.Prod(NewArtifact("rule", 1, 0)))
}

func TestNewArtifact(t *testing.T) {
	artifact := NewArtifact("rule", 0, 0)

	assert.Equal(t, "rule", artifact.rule)
	assert.Equal(t, 0, artifact.index)
	assert.Equal(t, 0, artifact.depth)
}

func TestArtifact_Equal(t *testing.T) {
	artifact1 := NewArtifact("rule", 0, 0)
	artifact2 := NewArtifact("rule", 0, 0)
	artifact3 := NewArtifact("rule", 1, 0)

	assert.True(t, artifact1.Equal(artifact2))
	assert.False(t, artifact1.Equal(artifact3))
}

func TestArtifact_String(t *testing.T) {
	artifact := NewArtifact("rule", 0, 0)

	assert.Equal(t, "rule, 0, 0", artifact.String())
}
