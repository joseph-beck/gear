package gear

type CST struct {
	Value    string
	Children []CST
	Label    Label
}

func NewCST(value ...string) CST {
	if len(value) == 1 {
		return CST{
			Value: value[0],
		}
	}

	return CST{}
}

func (c *CST) Add(cst CST) {
	c.Children = append(c.Children, cst)
}
