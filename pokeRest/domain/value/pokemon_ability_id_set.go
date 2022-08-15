package value

import "github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"

type PokemonAbilityIdSet struct {
	primary   identifier.AbilityId
	secondary identifier.AbilityId
	hidden    identifier.AbilityId
}

func NewPokemonAbilityIdSet(abilityIdPrimary, abilityIdSecondary, hiddenAbilityId identifier.AbilityId) *PokemonAbilityIdSet {
	return &PokemonAbilityIdSet{
		primary:   abilityIdPrimary,
		secondary: abilityIdSecondary,
		hidden:    hiddenAbilityId,
	}
}

func (a PokemonAbilityIdSet) GetIds() (identifier.AbilityId, identifier.AbilityId, identifier.AbilityId) {
	return a.primary, a.secondary, a.hidden
}
