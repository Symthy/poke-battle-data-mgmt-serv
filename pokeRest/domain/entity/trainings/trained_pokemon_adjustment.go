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

func (t TrainedPokemonAdjustment) Notify(note ITrainedPokemonAdjustNotification) {
	note.SetId(t.id)
	note.SetPokemonId(t.pokemonId)
	note.SetNature(t.nature)
	note.SetAbilityId(t.abilityId)
	note.SetHeldItemId(t.heldItemId)
	note.SetEffortValueH(t.effortValueH)
	note.SetEffortValueA(t.effortValueA)
	note.SetEffortValueB(t.effortValueB)
	note.SetEffortValueC(t.effortValueC)
	note.SetEffortValueD(t.effortValueD)
	note.SetEffortValueS(t.effortValueS)
	note.SetMoveId1(t.moveId1)
	note.SetMoveId2(t.moveId2)
	note.SetMoveId3(t.moveId3)
	note.SetMoveId4(t.moveId4)
}
