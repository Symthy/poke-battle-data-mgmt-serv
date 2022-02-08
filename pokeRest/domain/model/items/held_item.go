package items

type HeldItem struct {
	id uint
}

func NewHeldItem(id uint) HeldItem {
	return HeldItem{
		id: id,
	}
}

func (i HeldItem) Id() uint {
	return i.id
}
