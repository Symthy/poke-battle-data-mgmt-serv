package value

import "github.com/Symthy/PokeRest/pokeRest/errs"

// 努力値
type EffortValue struct {
	value int
	param PokemonParam
}

func NewEffortValue(value int, param PokemonParam) (*EffortValue, error) {
	v := value
	if value < 0 || value > 252 {
		return nil, errs.ThrowErrorInvalidValue("EffortValue", "value:"+string(param), string(rune(value)))
	}
	return &EffortValue{v, param}, nil
}

func (e EffortValue) Value() int {
	return e.value
}
