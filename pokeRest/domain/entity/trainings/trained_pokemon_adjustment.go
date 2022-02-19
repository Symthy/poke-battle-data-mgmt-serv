package trainings

import "github.com/Symthy/PokeRest/pokeRest/domain/value"

type TrainedPokemonAdjustment struct {
	id           uint
	pokemonId    int
	nature       value.PokemonNature
	abilityId    *int
	heldItemId   *int
	effortValueH value.EffortValue
	effortValueA value.EffortValue
	effortValueB value.EffortValue
	effortValueC value.EffortValue
	effortValueD value.EffortValue
	effortValueS value.EffortValue
	moveId1      *int
	moveId2      *int
	moveId3      *int
	moveId4      *int
}

func NewTrainedPokemonAdjustment(
	id uint, pokemonId int, nature string, abilityId *int, heldItemId *int,
	effortValueH int, effortValueA int, effortValueB int, effortValueC int, effortValueD int,
	effortValueS int, moveId1 *int, moveId2 *int, moveId3 *int, moveId4 *int) TrainedPokemonAdjustment {
	// Todo: add and validate
	return TrainedPokemonAdjustment{
		id:           id,
		pokemonId:    pokemonId,
		nature:       value.NewPokemonNature(nature),
		abilityId:    abilityId,
		heldItemId:   heldItemId,
		effortValueH: value.NewEffortValue(effortValueH),
		effortValueA: value.NewEffortValue(effortValueA),
		effortValueB: value.NewEffortValue(effortValueB),
		effortValueC: value.NewEffortValue(effortValueC),
		effortValueD: value.NewEffortValue(effortValueD),
		effortValueS: value.NewEffortValue(effortValueS),
		moveId1:      moveId1,
		moveId2:      moveId2,
		moveId3:      moveId3,
		moveId4:      moveId4,
	}
}

func (t TrainedPokemonAdjustment) Id() uint {
	return t.id
}
func (t TrainedPokemonAdjustment) PokemonId() int {
	return t.pokemonId
}
func (t TrainedPokemonAdjustment) Nature() value.PokemonNature {
	return t.nature
}
func (t TrainedPokemonAdjustment) AbilityId() *int {
	return t.abilityId
}
func (t TrainedPokemonAdjustment) HeldItemId() *int {
	return t.heldItemId
}

func (t TrainedPokemonAdjustment) EffortValueH() value.EffortValue {
	return t.effortValueH
}
func (t TrainedPokemonAdjustment) EffortValueA() value.EffortValue {
	return t.effortValueA
}
func (t TrainedPokemonAdjustment) EffortValueB() value.EffortValue {
	return t.effortValueB
}
func (t TrainedPokemonAdjustment) EffortValueC() value.EffortValue {
	return t.effortValueC
}
func (t TrainedPokemonAdjustment) EffortValueD() value.EffortValue {
	return t.effortValueD
}
func (t TrainedPokemonAdjustment) EffortValueS() value.EffortValue {
	return t.effortValueS
}

func (t TrainedPokemonAdjustment) MoveIdFirst() *int {
	return t.moveId1
}
func (t TrainedPokemonAdjustment) MoveIdSecond() *int {
	return t.moveId2
}
func (t TrainedPokemonAdjustment) MoveIdThird() *int {
	return t.moveId3
}
func (t TrainedPokemonAdjustment) MoveIdForth() *int {
	return t.moveId4
}
