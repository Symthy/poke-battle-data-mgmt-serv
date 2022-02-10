package moves

type Moves struct {
	items []Move
}

func (m Moves) Items() []Move {
	return m.items
}
