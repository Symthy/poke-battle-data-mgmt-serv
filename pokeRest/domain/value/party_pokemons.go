package value

type PartyPokemons struct {
	pokemonsIds []int
	size        int
}

func NewPartyPokemons(pokemonIds []*int) PartyPokemons {
	ids := []int{}
	for i, p := range pokemonIds {
		if p != nil && i < 6 {
			ids = append(ids, *p)
		}
	}
	return PartyPokemons{
		pokemonsIds: ids,
		size:        len(ids),
	}
}

func (p PartyPokemons) PokemonsIds() []int {
	return p.pokemonsIds
}

func (p PartyPokemons) Size() int {
	return p.size
}
