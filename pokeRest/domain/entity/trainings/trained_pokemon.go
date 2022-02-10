package trainings

type TrainedPokemon struct {
	id uint
}

func NewTrainedPokemon(id uint) TrainedPokemon {
	return TrainedPokemon{
		id: id,
	}
}

func (t TrainedPokemon) Id() uint {
	return t.id
}
