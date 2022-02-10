package trainings

type TrainedPokemonAttackTarget struct {
	id uint
}

func NewTrainedPokemonAttackTarget(id uint) TrainedPokemonAttackTarget {
	return TrainedPokemonAttackTarget{
		id: id,
	}
}

func (t TrainedPokemonAttackTarget) Id() uint {
	return t.id
}
