package gear

import "fmt"

type Key string

type PackratKey struct {
	// For NamedRules, this is the name of the rule.
	// For any other expression it's the pointer address.
	key Key
	// Position of the key in the input.
	pos uint
}

func NewPackratKey(expr Expression, pos uint) PackratKey {
	if rule, ok := expr.(*NamedRule); ok {
		return PackratKey{
			key: Key(rule.Value),
			pos: pos,
		}
	}

	return PackratKey{
		key: Key(fmt.Sprintf("%p", expr)),
		pos: pos,
	}
}

type PackratEntry struct {
	// Result of the packrat entry.
	result Result
	// Error of the packrat entry, if any.
	err error
	// Is this entry currently being seeded?
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
	entry, ok := p.memo[key]

	if ok && entry.seeding {
		return true
	}

	p.memo[key] = &PackratEntry{
		seeding: true,
	}

	return false
}

func (p *Packrat) Update(expr Expression, pos uint, result Result, err error) {
	key := NewPackratKey(expr, pos)
	entry, _ := p.memo[key]

	if entry != nil && entry.seeding {
		entry.result = result
		entry.err = err
	}
}

func (p *Packrat) Clear(expr Expression, pos uint) {
	key := NewPackratKey(expr, pos)

	for k := range p.memo {
		if k.pos == pos && k != key {
			delete(p.memo, k)
		}
	}
}
