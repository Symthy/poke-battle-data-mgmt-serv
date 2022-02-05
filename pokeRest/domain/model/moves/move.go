package moves

type Move struct {
	id uint
}

func (m Move) Id() uint {
	return m.id
}
