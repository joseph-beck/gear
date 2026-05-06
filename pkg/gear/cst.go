package gear

type CST struct {
	value    string
	children []CST
	label    label
}

type CSTParam struct {
	Value    string
	Children []CST
	Label    label
}

func NewCST(param ...CSTParam) CST {
	if len(param) == 0 {
		return CST{}
	}

	p := param[0]

	return CST{
		value:    p.Value,
		children: p.Children,
		label:    p.Label,
	}
}

func (c *CST) Add(cst CST) {
	c.children = append(c.children, cst)
}
