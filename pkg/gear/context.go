package gear

type Context struct {
	input    string
	position uint
	grammar  *Grammar
}

func NewContext() *Context {
	return &Context{
		input:    "",
		position: 0,
		grammar:  &Grammar{},
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

func (ctx *Context) Grammar() *Grammar {
	return ctx.grammar
}
