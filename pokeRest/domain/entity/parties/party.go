package parties

import "github.com/Symthy/PokeRest/pokeRest/domain/entity"

var _ entity.IDomain = (*Party)(nil)

type Party struct {
	id              uint
	trainedPokemons []string
}

func NewParty(id uint) Party {
	return Party{
		id: id,
	}
}

func (p Party) Id() uint {
	return p.id
}
