package value

import "github.com/Symthy/PokeRest/pokeRest/errs"

// 努力値
type EffortValue struct {
	value uint8
	param PokemonParam
}

func NewEffortValue(value uint64, param PokemonParam) (*EffortValue, error) {
	if value < 0 || value > 252 {
		return nil, errs.ThrowErrorInvalidValue("EffortValue", "value:"+string(param), string(rune(value)))
	}
	return &EffortValue{uint8(value), param}, nil
}

func (e EffortValue) Value() uint8 {
	return e.value
}
