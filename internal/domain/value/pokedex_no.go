package value

type PokedexId struct {
	pokedexNo uint16
	formNo    uint16
}

func NewPokedexNumber(pokedexNo, formNo uint64) (*PokedexId, error) {
	// if pokedexNo > xxx { // Todo: validate
	// 	return nil, errs.ThrowErrorInvalidValue("PokedexId", "pokedexNo", string(rune(pokedexNo)))
	// }
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
