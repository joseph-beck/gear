package label

type Label struct {
	productive bool
	hidden     bool
}

func New() Label {
	return Label{}
}
