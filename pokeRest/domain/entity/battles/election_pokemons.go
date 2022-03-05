package battles

type ElectionPokemons struct {
	pokemonIds []uint
	size       int
}

func NewElectionPokemons(pokemonIds []uint) ElectionPokemons {
	return ElectionPokemons{
		pokemonIds: pokemonIds,
		size:       len(pokemonIds),
	}
}

func (e ElectionPokemons) Ids() []uint {
	return e.pokemonIds
}

func (e ElectionPokemons) Get(i int) *uint {
	if i >= e.Size() {
		return nil
	}
	id := e.pokemonIds[i]
	return &id
}

func (e ElectionPokemons) Size() int {
	return e.size
}
