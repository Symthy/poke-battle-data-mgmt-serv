package pokemons

type TrainedPokemonMoveSet struct {
	id uint
}

func NewTrainedPokemonMoveSet(id uint) TrainedPokemonMoveSet {
	return TrainedPokemonMoveSet{
		id: id,
	}
}

func (t TrainedPokemonMoveSet) Id() uint {
	return t.id
}
