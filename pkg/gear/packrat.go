package gear

import "fmt"

type PackratKey struct {
	expr Expression
	pos  uint
}

type PackratEntry struct {
	result Result
	err    error
}

type Packrat struct {
	memo map[PackratKey]PackratEntry
}

func NewPackrat() Packrat {
	return Packrat{
		memo: make(map[PackratKey]PackratEntry),
	}
}

func (p *Packrat) Get(expr Expression, pos uint) (Result, error, bool) {
	entry, ok := p.memo[PackratKey{
		expr: expr,
		pos:  pos,
	}]

	if !ok {
		return Result{}, nil, false
	}

	fmt.Println(expr, pos, entry.result, entry.err)

	return entry.result, entry.err, true
}

func (p *Packrat) Put(expr Expression, pos uint, result Result, err error) {
	p.memo[PackratKey{
		expr: expr,
		pos:  pos,
	}] = PackratEntry{
		result: result,
		err:    err,
	}
}
