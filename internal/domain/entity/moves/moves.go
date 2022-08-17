package moves

type Moves struct {
	items []*Move
}

func NewMoves(items []*Move) *Moves {
	return &Moves{items: items}
}

func (m Moves) Items() []*Move {
	return m.items
}
