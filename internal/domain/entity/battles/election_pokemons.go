package battles

import "github.com/Symthy/PokeRest/internal/pkg/lists"

type ElectionPokemons struct {
	pokemonIds []uint16
	size       uint8
}

func NewElectionPokemons(pokemonIds []uint64) *ElectionPokemons {
	// Todo: validate id upper limit
	return &ElectionPokemons{
		pokemonIds: lists.ConvertTypeUint64To16(pokemonIds),
		size:       uint8(len(pokemonIds)),
	}
}

func (e ElectionPokemons) Ids() []uint16 {
	return e.pokemonIds
}

func (e ElectionPokemons) Get(i int) *uint16 {
	if i >= int(e.Size()) {
		return nil
	}
	id := e.pokemonIds[i]
	return &id
}

func (e ElectionPokemons) Size() uint8 {
	return e.size
}
