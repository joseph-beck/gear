package gear

type Context struct {
	input string

	grammar *Grammar

	packrat Packrat

	seeding bool
}

func NewContext(input string) *Context {
	return &Context{
		input:   input,
		grammar: &Grammar{},
		packrat: NewPackrat(),
		seeding: false,
	}
}

func (ctx *Context) Clone() *Context {
	return &Context{
		input:   ctx.input,
		grammar: ctx.grammar,
		packrat: ctx.packrat,
		seeding: ctx.seeding,
	}
}

func (ctx Context) Input() string {
	return ctx.input
}

func (ctx *Context) SetInput(input string) {
	ctx.input = input
}

func (ctx *Context) Grammar() *Grammar {
	return ctx.grammar
}

func (ctx *Context) Packrat() *Packrat {
	return &ctx.packrat
}

func (ctx *Context) Seeding() bool {
	return ctx.seeding
}

func (ctx *Context) SetSeeding(seeding bool) {
	ctx.seeding = seeding
}
