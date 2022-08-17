package value

type PartyPokemonIds struct {
	pokemonsIds []*uint16
	size        uint8
}

func NewPartyPokemonIds(pokemonIds []uint64) *PartyPokemonIds {
	// Todo: validate id
	ids := make([]*uint16, 6)
	for i, pid := range pokemonIds {
		id := uint16(pid)
		if i < 6 {
			ids[i] = &id
		}
	}
	return &PartyPokemonIds{
		pokemonsIds: ids,
		size:        uint8(len(ids)),
	}
}

func (p PartyPokemonIds) Ids() []*uint16 {
	return p.pokemonsIds
}

func (p PartyPokemonIds) Get(i uint) *uint16 {
	if i >= uint(p.size) {
		return nil
	}
	return p.pokemonsIds[i]
}

func (p PartyPokemonIds) Size() uint8 {
	return p.size
}
