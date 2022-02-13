package items

type HeldItems struct {
	items []HeldItem
}

func NewHeldItems(items []HeldItem) HeldItems {
	return HeldItems{items: items}
}

func (i HeldItems) Items() []HeldItem {
	return i.items
}
