package gear

import "fmt"

type PackratKey struct {
	// For NamedRules, this is "named:ruleName", for others it's the pointer address
	ruleKey string
	// Position of the key in the input
	pos uint
}

func NewPackratKey(expr Expression, pos uint) PackratKey {
	if nr, ok := expr.(*NamedRule); ok {
		return PackratKey{
			ruleKey: "named:" + nr.Value,
			pos:     pos,
		}
	}

	return PackratKey{
		ruleKey: fmt.Sprintf("%p", expr),
		pos:     pos,
	}
}

type PackratEntry struct {
	result Result

	err error

	seeding bool
}

type Packrat struct {
	memo map[PackratKey]*PackratEntry
}

func NewPackrat() Packrat {
	return Packrat{
		memo: make(map[PackratKey]*PackratEntry),
	}
}

func (p *Packrat) Get(expr Expression, pos uint) (Result, error, bool) {
	key := NewPackratKey(expr, pos)
	entry, ok := p.memo[key]

	if !ok {
		return Result{}, nil, false
	}

	return entry.result, entry.err, true
}

func (p *Packrat) Put(expr Expression, pos uint, result Result, err error) {
	key := NewPackratKey(expr, pos)
	p.memo[key] = &PackratEntry{
		result:  result,
		err:     err,
		seeding: false,
	}
}

func (p *Packrat) Mark(expr Expression, pos uint) bool {
	key := NewPackratKey(expr, pos)
	entry, exists := p.memo[key]

	if exists && entry.seeding {
		return true
	}

	p.memo[key] = &PackratEntry{
		seeding: true,
	}

	return false
}

func (p *Packrat) Update(expr Expression, pos uint, result Result, err error) {
	key := NewPackratKey(expr, pos)
	entry := p.memo[key]
	if entry != nil && entry.seeding {
		entry.result = result
		entry.err = err
	}
}

func (p *Packrat) Clear(pos uint, except Expression) {
	exceptKey := NewPackratKey(except, pos)
	for key := range p.memo {
		if key.pos == pos && key != exceptKey {
			delete(p.memo, key)
		}
	}
}
