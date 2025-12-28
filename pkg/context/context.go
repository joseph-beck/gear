package context

import (
	"github.com/joseph-beck/gear/pkg/rule"
)

type Context struct {
	input    string
	position uint
	rules    []rule.Rule
}

func New() *Context {
	return &Context{
		input:    "",
		position: 0,
		rules:    []rule.Rule{},
	}
}

func (ctx *Context) Input() string {
	return ctx.input
}

func (ctx *Context) SetInput(input string) {
	ctx.input = input
}

func (ctx *Context) Position() uint {
	return ctx.position
}

func (ctx *Context) SetPosition(pos uint) {
	ctx.position = pos
}

func (ctx *Context) Rules() []rule.Rule {
	return ctx.rules
}

func (ctx *Context) SetRules(rules []rule.Rule) {
	ctx.rules = rules
}
