package value

import "github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"

type PokemonAbilitySet struct {
	abilityIdPrimary   identifier.AbilityId
	abilityIdSecondary identifier.AbilityId
	hiddenAbilityId    identifier.AbilityId
}

func NewPokemonAbilitySet(abilityIdPrimary, abilityIdSecondary, hiddenAbilityId identifier.AbilityId) PokemonAbilitySet {
	return PokemonAbilitySet{abilityIdPrimary, abilityIdSecondary, hiddenAbilityId}
}

func (a PokemonAbilitySet) GetAbilityIds() (identifier.AbilityId, identifier.AbilityId, identifier.AbilityId) {
	return a.abilityIdPrimary, a.abilityIdSecondary, a.hiddenAbilityId
}
