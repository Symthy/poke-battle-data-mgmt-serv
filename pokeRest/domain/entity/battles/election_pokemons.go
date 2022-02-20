package battles

type ElectionPokemons[T int | uint] struct {
	pokemonIds []T
	size       int
}

func NewElectionPokemons[T int | uint](pokemonIds []T) ElectionPokemons[T] {
	return ElectionPokemons[T]{
		pokemonIds: pokemonIds,
		size:       len(pokemonIds),
	}
}

func (e ElectionPokemons[T]) PokemonIds() []T {
	return e.pokemonIds
}

func (e ElectionPokemons[T]) Size() int {
	return e.size
}
