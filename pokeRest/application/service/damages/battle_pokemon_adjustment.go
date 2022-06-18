package damages

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type BattlePokemonAdjustment struct {
	pokemonId    identifier.PokemonId
	abilityId    identifier.AbilityId
	heldItemId   identifier.HeldItemId
	nature       value.PokemonNature
	effortValues *value.EffortValues
}

func NewBattlePokemonAdjustment(
	pokemonId identifier.PokemonId, abilityId identifier.AbilityId, heldItemId identifier.HeldItemId,
	nature value.PokemonNature, effortValues *value.EffortValues,
) *BattlePokemonAdjustment {
	return &BattlePokemonAdjustment{
		pokemonId:    pokemonId,
		abilityId:    abilityId,
		heldItemId:   heldItemId,
		nature:       nature,
		effortValues: effortValues,
	}
}

func (a BattlePokemonAdjustment) PokemonId() identifier.PokemonId {
	return a.pokemonId
}
func (a BattlePokemonAdjustment) AbilityId() identifier.AbilityId {
	return a.abilityId
}
func (a BattlePokemonAdjustment) HeldItemId() identifier.HeldItemId {
	return a.heldItemId
}
func (a BattlePokemonAdjustment) Nature() value.PokemonNature {
	return a.nature
}
func (a BattlePokemonAdjustment) EffortValues() *value.EffortValues {
	return a.effortValues
}
