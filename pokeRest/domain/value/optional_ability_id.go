package value

import "github.com/Symthy/PokeRest/pokeRest/exception"

type OptionalAbilityId struct {
	id *int
}

func NewOptionalAbilityId(id *int) OptionalAbilityId {
	return OptionalAbilityId{id: id}
}

func NewOptionalAbilityIdEmpty() OptionalAbilityId {
	return OptionalAbilityId{id: nil}
}

func (o OptionalAbilityId) Get() (*int, error) {
	if o.isEmpty() {
		return nil, exception.NewNoValueError("Ability ID")
	}
	return o.id, nil
}

func (o OptionalAbilityId) isEmpty() bool {
	return o.id == nil
}
