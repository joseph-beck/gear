package gear

type label struct {
	hidden     bool
	expression bool
}

type LabelParam struct {
	Hidden     bool
	Expression bool
}

func NewLabel(param ...LabelParam) label {
	if len(param) == 0 {
		return label{}
	}

	p := param[0]
	return label{
		hidden:     p.Hidden,
		expression: p.Expression,
	}
}
