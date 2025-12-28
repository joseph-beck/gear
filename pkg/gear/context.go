package gear

type Context struct {
	input    string
	position uint
	grammar  *Grammar
}

type contextCfg struct {
	input   string
	grammar *Grammar
}

func NewContext(cfg ...contextCfg) *Context {
	if len(cfg) > 0 {
		return &Context{
			input:    cfg[0].input,
			position: 0,
			grammar:  cfg[0].grammar,
		}
	}

	return &Context{
		input:    "",
		position: 0,
		grammar:  &Grammar{},
	}
}

func (ctx *Context) Clone() *Context {
	return &Context{
		input:    ctx.input,
		position: ctx.position,
		grammar:  ctx.grammar,
	}
}

func (ctx Context) Input() string {
	return ctx.input
}

func (ctx *Context) SetInput(input string) {
	ctx.input = input
}

func (ctx *Context) Remaining() string {
	return ctx.input[ctx.position:]
}

func (ctx Context) Position() uint {
	return ctx.position
}

func (ctx *Context) SetPosition(pos uint) {
	ctx.position = pos
}

func (ctx *Context) Grammar() *Grammar {
	return ctx.grammar
}
