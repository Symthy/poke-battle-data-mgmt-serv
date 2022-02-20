package value

type PartyPokemonIds struct {
	pokemonsIds []int
	size        int
}

func NewPartyPokemonIds(pokemonIds ...int) PartyPokemonIds {
	ids := []int{}
	for i, p := range pokemonIds {
		if i < 6 {
			ids = append(ids, p)
		}
	}
	return PartyPokemonIds{
		pokemonsIds: ids,
		size:        len(ids),
	}
}

func (p PartyPokemonIds) PokemonsIds() []int {
	return p.pokemonsIds
}

func (p PartyPokemonIds) Size() int {
	return p.size
}
