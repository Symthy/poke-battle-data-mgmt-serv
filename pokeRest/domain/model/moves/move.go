package moves

type Move struct {
	id uint
}

func NewMove(id uint) Move {
	return Move{
		id: id,
	}
}

func (m Move) Id() uint {
	return m.id
}
