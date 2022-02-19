package battles

type ElectionPokemons struct {
	pokemonIds []int
	size       int
}

func NewElectionPokemons(pokemonIds []int) ElectionPokemons {
	return ElectionPokemons{
		pokemonIds: pokemonIds,
		size:       len(pokemonIds),
	}
}

func (e ElectionPokemons) PokemonIds() []int {
	return e.pokemonIds
}

func (e ElectionPokemons) Size() int {
	return e.size
}
