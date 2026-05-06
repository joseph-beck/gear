package gear

type PackratKey struct {
	rule string
	pos  uint
}

func NewPackratKey(rule string, pos uint) PackratKey {
	return PackratKey{
		rule: rule,
		pos:  pos,
	}
}

type PackratEntry struct {
	result Result
	err    error
}

func NewPackratEntry(result Result, err error) *PackratEntry {
	return &PackratEntry{
		result: result,
		err:    err,
	}
}

func (p *PackratEntry) Clone() *PackratEntry {
	return &PackratEntry{
		result: p.result,
		err:    p.err,
	}
}

type Packrat struct {
	memo map[PackratKey]*PackratEntry
}

func NewPackrat() *Packrat {
	return &Packrat{
		memo: make(map[PackratKey]*PackratEntry),
	}
}

func (p *Packrat) Get(key PackratKey) (*PackratEntry, bool) {
	entry, ok := p.memo[key]

	return entry, ok
}

func (p *Packrat) Set(key PackratKey, entry *PackratEntry) {
	p.memo[key] = entry
}

func (p *Packrat) Clone() Packrat {
	newMemo := make(map[PackratKey]*PackratEntry)

	for key, entry := range p.memo {
		newMemo[key] = entry.Clone()
	}

	return Packrat{
		memo: newMemo,
	}
}
