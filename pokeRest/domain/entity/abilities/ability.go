package abilities

type Ability struct {
	id uint
}

func NewAbility(id uint) Ability {
	return Ability{
		id: id,
	}
}

func (a Ability) Id() uint {
	return a.id
}
