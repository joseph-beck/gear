package cst

import "github.com/joseph-beck/gear/pkg/label"

type CST struct {
	Value    string
	Children []CST
	Label    label.Label
}

func New(value ...string) CST {
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
