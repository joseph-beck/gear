package gear

type Context struct {
	input   string
	pos     uint
	grammar *Grammar
	packrat *Packrat
	history *History
	rule    string
	depth   int
}

func NewContext(input string) *Context {
	history := NewHistory()
	grammar := NewGrammar()
	packrat := NewPackrat()

	return &Context{
		input:   input,
		pos:     0,
		grammar: &grammar,
		packrat: &packrat,
		history: &history,
		rule:    "",
		depth:   0,
	}
}

func (ctx *Context) Clone() *Context {
	grammar := ctx.grammar.Clone()
	packrat := ctx.packrat.Clone()
	history := ctx.history.Clone()

	return &Context{
		input:   ctx.input,
		pos:     ctx.pos,
		grammar: &grammar,
		packrat: &packrat,
		history: &history,
		rule:    ctx.rule,
		depth:   ctx.depth,
	}
}

func (ctx *Context) Update(other *Context) {
	ctx.pos = other.pos
	ctx.packrat = other.packrat
	ctx.history = other.history
	ctx.rule = other.rule
	ctx.depth = other.depth
}

func (ctx *Context) Remaining() string {
	return ctx.input[ctx.pos:]
}
