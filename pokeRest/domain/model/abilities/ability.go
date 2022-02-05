package abilities

type Ability struct {
	id uint
}

func (a Ability) Id() uint {
	return a.id
}
