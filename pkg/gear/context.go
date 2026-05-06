package gear

type Context struct {
	input   string
	pos     uint
	grammar *Grammar
	packrat Packrat
	history *History
	rule    string
	depth   int
}

func NewContext(input string) *Context {
	return &Context{
		input:   input,
		pos:     0,
		grammar: &Grammar{},
		packrat: NewPackrat(),
		history: &History{},
		rule:    "",
		depth:   0,
	}
}

func (ctx *Context) Clone() *Context {
	return &Context{
		input:   ctx.input,
		pos:     ctx.pos,
		grammar: ctx.grammar,
		packrat: ctx.packrat,
		history: ctx.history,
		rule:    ctx.rule,
		depth:   ctx.depth,
	}
}
