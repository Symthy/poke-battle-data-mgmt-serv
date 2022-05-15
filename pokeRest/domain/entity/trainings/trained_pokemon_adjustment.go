package trainings

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

var _ entity.IDomain[identifier.TrainedAdjustmentId, uint64] = (*TrainedPokemonAdjustment)(nil)

type TrainedPokemonAdjustment struct {
	id           identifier.TrainedAdjustmentId
	pokemonId    identifier.PokemonId
	nature       value.PokemonNature
	abilityId    identifier.AbilityId
	heldItemId   identifier.HeldItemId
	effortValues value.EffortValues
	moveSet      value.PokemonMoveIdSet
}

func NewTrainedPokemonAdjustment(
	id identifier.TrainedAdjustmentId, pokemonId identifier.PokemonId, nature value.PokemonNature,
	abilityId identifier.AbilityId, heldItemId identifier.HeldItemId, effortValues value.EffortValues,
	moveSet value.PokemonMoveIdSet,
) TrainedPokemonAdjustment {
	// Todo: add and validate
	return TrainedPokemonAdjustment{
		id:           id,
		pokemonId:    pokemonId,
		nature:       nature,
		abilityId:    abilityId,
		heldItemId:   heldItemId,
		effortValues: effortValues,
		moveSet:      moveSet,
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
	note.SetEffortValueH(t.effortValues.H())
	note.SetEffortValueA(t.effortValues.A())
	note.SetEffortValueB(t.effortValues.B())
	note.SetEffortValueC(t.effortValues.C())
	note.SetEffortValueD(t.effortValues.D())
	note.SetEffortValueS(t.effortValues.S())
	moveId1, moveId2, moveId3, moveId4 := t.moveSet.GetMoveIds()
	note.SetMoveId1(moveId1)
	note.SetMoveId2(moveId2)
	note.SetMoveId3(moveId3)
	note.SetMoveId4(moveId4)
}
