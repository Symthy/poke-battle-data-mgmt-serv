package trainings

type TrainedPokemonDeffenceTarget struct {
	id uint
}

func NewTrainedPokemonDeffenceTarget(id uint) TrainedPokemonDeffenceTarget {
	return TrainedPokemonDeffenceTarget{
		id: id,
	}
}

func (t TrainedPokemonDeffenceTarget) Id() uint {
	return t.id
}
