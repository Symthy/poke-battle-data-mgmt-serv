package pokemons

type Pokemons struct {
	pokemons []*Pokemon
}

func NewPokemons(pokemons []*Pokemon) *Pokemons {
	return &Pokemons{pokemons: pokemons}
}

func (p Pokemons) Items() []*Pokemon {
	return p.pokemons
}

func (p Pokemons) Count() int {
	return len(p.pokemons)
}

func (p Pokemons) IsEmpty() bool {
	return len(p.pokemons) == 0
}
