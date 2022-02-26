package trainings

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

var _ entity.IDomain[identifier.TrainedAdjustmentId] = (*TrainedPokemonAdjustment)(nil)

type TrainedPokemonAdjustment struct {
	id           identifier.TrainedAdjustmentId
	pokemonId    identifier.PokemonId
	nature       value.PokemonNature
	abilityId    identifier.AbilityId
	heldItemId   identifier.HeldItemId
	effortValueH value.EffortValue
	effortValueA value.EffortValue
	effortValueB value.EffortValue
	effortValueC value.EffortValue
	effortValueD value.EffortValue
	effortValueS value.EffortValue
	moveId1      identifier.MoveId
	moveId2      identifier.MoveId
	moveId3      identifier.MoveId
	moveId4      identifier.MoveId
}

func NewTrainedPokemonAdjustment(
	id identifier.TrainedAdjustmentId, pokemonId identifier.PokemonId, nature string,
	abilityId identifier.AbilityId, heldItemId identifier.HeldItemId, effortValueH int,
	effortValueA int, effortValueB int, effortValueC int, effortValueD int, effortValueS int,
	moveId1 identifier.MoveId, moveId2 identifier.MoveId, moveId3 identifier.MoveId, moveId4 identifier.MoveId,
) TrainedPokemonAdjustment {
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

// Todo: refactor Notification
func (t TrainedPokemonAdjustment) Id() identifier.TrainedAdjustmentId {
	return t.id
}
func (t TrainedPokemonAdjustment) PokemonId() identifier.PokemonId {
	return t.pokemonId
}
func (t TrainedPokemonAdjustment) Nature() value.PokemonNature {
	return t.nature
}
func (t TrainedPokemonAdjustment) AbilityId() identifier.AbilityId {
	return t.abilityId
}
func (t TrainedPokemonAdjustment) HeldItemId() identifier.HeldItemId {
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

func (t TrainedPokemonAdjustment) MoveIdFirst() identifier.MoveId {
	return t.moveId1
}
func (t TrainedPokemonAdjustment) MoveIdSecond() identifier.MoveId {
	return t.moveId2
}
func (t TrainedPokemonAdjustment) MoveIdThird() identifier.MoveId {
	return t.moveId3
}
func (t TrainedPokemonAdjustment) MoveIdForth() identifier.MoveId {
	return t.moveId4
}
