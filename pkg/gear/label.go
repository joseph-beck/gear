package gear

type Label struct {
	productive bool
	hidden     bool
}

func NewLabel() Label {
	return Label{}
}
