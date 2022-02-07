package pokemons

type TrainedPokemonBase struct {
	id uint
}

func (t TrainedPokemonBase) Id() uint {
	return t.id
}
