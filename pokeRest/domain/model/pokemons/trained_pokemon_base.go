package pokemons

type TrainedPokemonBase struct {
	id uint
}

func NewTrainedPokemonBase(id uint) TrainedPokemonBase {
	return TrainedPokemonBase{
		id: id,
	}
}

func (t TrainedPokemonBase) Id() uint {
	return t.id
}
