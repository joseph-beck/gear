package gear

type Label struct {
	Hidden     bool
	Expression bool
}

type labelParam struct {
	hidden     bool
	expression bool
}

func NewLabel(param ...labelParam) Label {
	if len(param) == 0 {
		return Label{}
	}

	p := param[0]
	return Label{
		Hidden:     p.hidden,
		Expression: p.expression,
	}
}
