package damages

import "github.com/Symthy/PokeRest/pokeRest/domain/value"

type BattlePokemonAdjustment struct {
	pokemonId    uint
	abilityId    uint
	heldItemId   uint
	nature       value.PokemonNature
	effortValues value.EffortValues
}

func NewBattlePokemonAdjustment(pokemonId uint, abilityId uint, heldItemId uint,
	nature value.PokemonNature, effortValues value.EffortValues) BattlePokemonAdjustment {
	return BattlePokemonAdjustment{
		pokemonId:    pokemonId,
		abilityId:    abilityId,
		heldItemId:   heldItemId,
		nature:       nature,
		effortValues: effortValues,
	}
}

func (a BattlePokemonAdjustment) PokemonId() uint {
	return a.pokemonId
}
func (a BattlePokemonAdjustment) AbilityId() uint {
	return a.abilityId
}
func (a BattlePokemonAdjustment) HeldItemId() uint {
	return a.heldItemId
}
func (a BattlePokemonAdjustment) Nature() value.PokemonNature {
	return a.nature
}
func (a BattlePokemonAdjustment) EffortValues() value.EffortValues {
	return a.effortValues
}
