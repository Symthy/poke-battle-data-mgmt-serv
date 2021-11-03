package value

import "github.com/Symthy/PokeRest/pokeRest/exception"

type OptionaAbilitylId struct {
	id *int
}

func NewOptionalAbilityId(id *int) OptionaAbilitylId {
	return OptionaAbilitylId{id: id}
}

func NewOptionalAbilityIdEmpty() OptionaAbilitylId {
	return OptionaAbilitylId{id: nil}
}

func (o OptionaAbilitylId) Get() (*int, error) {
	if o.IsEmpty() {
		return nil, exception.NewNoValueError("Ability ID")
	}
	return o.id, nil
}

func (o OptionaAbilitylId) IsEmpty() bool {
	return o.id == nil
}
