package pokemons

type PokemonList struct {
	pokemons []Pokemon
}

func NewPokemonList(pokemons []Pokemon) PokemonList {
	return PokemonList{pokemons: pokemons}
}

func (li PokemonList) Items() []Pokemon {
	return li.pokemons
}

func (li PokemonList) Count() int {
	return len(li.pokemons)
}

func (li PokemonList) IsEmpty() bool {
	return len(li.pokemons) == 0
}
