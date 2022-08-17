package parties

type Parties struct {
	items []Party
}

func NewParties(parties []Party) Parties {
	return Parties{items: parties}
}

func (p Parties) Items() []Party {
	return p.items
}
