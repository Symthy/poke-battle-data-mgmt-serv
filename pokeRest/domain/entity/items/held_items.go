package items

type HeldItems struct {
	items []HeldItem
}

func (i HeldItems) Items() []HeldItem {
	return i.items
}
