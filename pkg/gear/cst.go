package gear

type cst struct {
	value    string
	children []cst
	label    label
}

type CSTParam struct {
	Value    string
	Children []cst
	Label    label
}

func NewCST(param ...CSTParam) cst {
	if len(param) == 0 {
		return cst{}
	}

	p := param[0]
	return cst{
		value:    p.Value,
		children: p.Children,
		label:    p.Label,
	}
}

func (c *cst) Add(cst cst) {
	c.children = append(c.children, cst)
}
