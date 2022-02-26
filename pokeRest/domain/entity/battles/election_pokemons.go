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

func (e ElectionPokemons[T]) Ids() []T {
	return e.pokemonIds
}

func (e ElectionPokemons[T]) Get(i int) *T {
	if i >= e.Size() {
		return nil
	}
	id := e.pokemonIds[i]
	return &id
}

func (e ElectionPokemons[T]) Size() int {
	return e.size
}
