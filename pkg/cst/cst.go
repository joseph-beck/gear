package cst

import "github.com/joseph-beck/gear/pkg/label"

type CST struct {
	value    string
	children []CST
	label    label.Label
}

func New() CST {
	return CST{}
}

func (c CST) Value() string {
	return c.value
}
