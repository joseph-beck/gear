package gear

import "fmt"

type Artifact struct {
	rule  string
	index int
	depth int
}

func NewArtifact(rule string, index int, depth int) Artifact {
	return Artifact{
		rule:  rule,
		index: index,
		depth: depth,
	}
}

func (a Artifact) Equal(other Artifact) bool {
	return a.rule == other.rule && a.index == other.index && a.depth == other.depth
}

func (a Artifact) String() string {
	return fmt.Sprintf("%s, %d, %d", a.rule, a.index, a.depth)
}

type History []Artifact

func NewHistory() History {
	return make(History, 0)
}

func (h History) Clone() History {
	clone := make(History, len(h))
	copy(clone, h)

	return clone
}

func (h *History) Preserve(artifact Artifact) {
	*h = append(*h, artifact)
}

func (h *History) Clear() {
	*h = (*h)[:0]
}

func (h *History) Prod(artifact Artifact) bool {
	for _, a := range *h {
		if a.Equal(artifact) {
			return true
		}
	}

	return false
}
