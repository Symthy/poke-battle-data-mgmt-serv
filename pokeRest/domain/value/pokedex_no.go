package value

import "github.com/Symthy/PokeRest/pokeRest/errs"

type PokedexId struct {
	pokedexNo uint16
	formNo    uint16
}

func NewPokedexNumber(pokedexNo, formNo uint64) (*PokedexId, error) {
	if pokedexNo < 0 {
		return nil, errs.ThrowErrorInvalidValue("PokedexId", "pokedexNo", string(rune(pokedexNo)))
	}
	if pokedexNo < 0 {
		return nil, errs.ThrowErrorInvalidValue("PokedexId", "formNo", string(rune(pokedexNo)))
	}
	return &PokedexId{
		pokedexNo: uint16(pokedexNo),
		formNo:    uint16(formNo),
	}, nil
}

func (p PokedexId) PokedexNo() uint16 {
	return p.pokedexNo
}

func (p PokedexId) FormNo() uint16 {
	return p.formNo
}
