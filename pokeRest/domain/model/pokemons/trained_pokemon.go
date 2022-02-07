package pokemons

type TrainedPokemon struct {
	id uint
}

func (t TrainedPokemon) Id() uint {
	return t.id
}
