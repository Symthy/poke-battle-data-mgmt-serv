package battles

type ElectionTrainedPokemons struct {
	pokemonIds []uint64
	size       uint8
}

func NewElectionTrainedPokemons(pokemonIds []uint64) *ElectionTrainedPokemons {
	// Todo: validate id upper limit
	return &ElectionTrainedPokemons{
		pokemonIds: pokemonIds,
		size:       uint8(len(pokemonIds)),
	}
}

func (e ElectionTrainedPokemons) Ids() []uint64 {
	return e.pokemonIds
}

func (e ElectionTrainedPokemons) Get(i int) *uint64 {
	if i >= int(e.Size()) {
		return nil
	}
	id := e.pokemonIds[i]
	return &id
}

func (e ElectionTrainedPokemons) Size() uint8 {
	return e.size
}
