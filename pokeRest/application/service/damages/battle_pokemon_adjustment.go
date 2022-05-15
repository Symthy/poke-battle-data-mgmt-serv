package damages

import "github.com/Symthy/PokeRest/pokeRest/domain/value"

type BattlePokemonAdjustment struct {
	pokemonId    uint16
	abilityId    uint16
	heldItemId   uint16
	nature       value.PokemonNature
	effortValues value.EffortValues
}

func NewBattlePokemonAdjustment(pokemonId uint16, abilityId uint16, heldItemId uint16,
	nature value.PokemonNature, effortValues value.EffortValues) BattlePokemonAdjustment {
	return BattlePokemonAdjustment{
		pokemonId:    pokemonId,
		abilityId:    abilityId,
		heldItemId:   heldItemId,
		nature:       nature,
		effortValues: effortValues,
	}
}

func (a BattlePokemonAdjustment) PokemonId() uint16 {
	return a.pokemonId
}
func (a BattlePokemonAdjustment) AbilityId() uint16 {
	return a.abilityId
}
func (a BattlePokemonAdjustment) HeldItemId() uint16 {
	return a.heldItemId
}
func (a BattlePokemonAdjustment) Nature() value.PokemonNature {
	return a.nature
}
func (a BattlePokemonAdjustment) EffortValues() value.EffortValues {
	return a.effortValues
}
