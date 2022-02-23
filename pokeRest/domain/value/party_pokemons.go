package value

type PartyPokemonIds struct {
	pokemonsIds []*int
	size        int
}

func NewPartyPokemonIds(pokemonIds ...int) PartyPokemonIds {
	ids := make([]*int, 6, 6)
	for i, p := range pokemonIds {
		if i < 6 {
			ids[i] = &p
		}
	}
	return PartyPokemonIds{
		pokemonsIds: ids,
		size:        len(ids),
	}
}

func (p PartyPokemonIds) Ids() []*int {
	return p.pokemonsIds
}

func (p PartyPokemonIds) Get(i int) *int {
	if i >= p.size {
		return nil
	}
	return p.pokemonsIds[i]
}

func (p PartyPokemonIds) Size() int {
	return p.size
}
