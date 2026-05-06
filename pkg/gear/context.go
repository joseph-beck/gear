package gear

type Context struct {
	input   string
	grammar *Grammar
	packrat Packrat
	seeding bool
	history History
	rule    string
	depth   int
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
		rule:    ctx.rule,
		depth:   ctx.depth,
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

func (ctx *Context) History() History {
	return ctx.history
}

func (ctx *Context) SetHistory(history History) {
	ctx.history = history
}

func (ctx *Context) Rule() string {
	return ctx.rule
}

func (ctx *Context) SetRule(rule string) {
	ctx.rule = rule
}

func (ctx *Context) Depth() int {
	return ctx.depth
}

func (ctx *Context) SetDepth(depth int) {
	ctx.depth = depth
}
