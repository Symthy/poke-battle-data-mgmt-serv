package items

type HeldItem struct {
	id uint
}

func (i HeldItem) Id() uint {
	return i.id
}
