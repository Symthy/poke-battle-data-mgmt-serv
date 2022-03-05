package value

type PartyPokemonIds struct {
	pokemonsIds []*uint
	size        int
}

func NewPartyPokemonIds(pokemonIds []uint) PartyPokemonIds {
	ids := make([]*uint, 6)
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

func (p PartyPokemonIds) Ids() []*uint {
	return p.pokemonsIds
}

func (p PartyPokemonIds) Get(i int) *uint {
	if i >= p.size {
		return nil
	}
	return p.pokemonsIds[i]
}

func (p PartyPokemonIds) Size() int {
	return p.size
}
