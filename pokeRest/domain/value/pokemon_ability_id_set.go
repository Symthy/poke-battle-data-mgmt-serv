package value

import "github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"

type PokemonAbilityIdSet struct {
	abilityIdPrimary   identifier.AbilityId
	abilityIdSecondary identifier.AbilityId
	hiddenAbilityId    identifier.AbilityId
}

func NewPokemonAbilityIdSet(abilityIdPrimary, abilityIdSecondary, hiddenAbilityId identifier.AbilityId) PokemonAbilityIdSet {
	return PokemonAbilityIdSet{abilityIdPrimary, abilityIdSecondary, hiddenAbilityId}
}

func (a PokemonAbilityIdSet) GetAbilityIds() (identifier.AbilityId, identifier.AbilityId, identifier.AbilityId) {
	return a.abilityIdPrimary, a.abilityIdSecondary, a.hiddenAbilityId
}
