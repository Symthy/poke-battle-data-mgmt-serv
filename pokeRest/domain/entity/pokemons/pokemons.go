package pokemons

type Pokemons struct {
	pokemons []Pokemon
}

func NewPokemons(pokemons []Pokemon) Pokemons {
	return Pokemons{pokemons: pokemons}
}

func (li Pokemons) Items() []Pokemon {
	return li.pokemons
}

func (li Pokemons) Count() int {
	return len(li.pokemons)
}

func (li Pokemons) IsEmpty() bool {
	return len(li.pokemons) == 0
}
