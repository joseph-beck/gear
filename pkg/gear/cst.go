package gear

type CST struct {
	Value    string
	Children []CST
	Label    Label
}

type cstParam struct {
	value    string
	children []CST
	label    Label
}

func NewCST(param ...cstParam) CST {
	if len(param) == 0 {
		return CST{}
	}

	p := param[0]
	return CST{
		Value:    p.value,
		Children: p.children,
		Label:    p.label,
	}
}

func (c *CST) Add(cst CST) {
	c.Children = append(c.Children, cst)
}
