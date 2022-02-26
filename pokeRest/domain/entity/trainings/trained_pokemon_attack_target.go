package trainings

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

var _ entity.IDomain[identifier.TrainedAttackTargetId] = (*TrainedPokemonAttackTarget)(nil)

type TrainedPokemonAttackTarget struct {
	id                   identifier.TrainedAttackTargetId
	trainedPokemonId     identifier.TrainedPokemonId
	moveId               identifier.MoveId
	targetPokemonId      identifier.PokemonId
	targetPokemonNature  value.PokemonNature
	targetPokemonEffortH value.EffortValue
	targetPokemonEffortB value.EffortValue
	targetPokemonEffortD value.EffortValue
}

// Todo: factory
func NewTrainedPokemonAttackTarget(
	id identifier.TrainedAttackTargetId, trainedPokemonId identifier.TrainedPokemonId,
	moveId identifier.MoveId, targetPokemonId identifier.PokemonId, targetPokemonNature string,
	targetPokemonEffortH int, targetPokemonEffortB int, targetPokemonEffortD int) TrainedPokemonAttackTarget {
	return TrainedPokemonAttackTarget{
		id:                   id,
		trainedPokemonId:     trainedPokemonId,
		moveId:               moveId,
		targetPokemonId:      targetPokemonId,
		targetPokemonNature:  value.NewPokemonNature(targetPokemonNature),
		targetPokemonEffortH: value.NewEffortValue(targetPokemonEffortH),
		targetPokemonEffortB: value.NewEffortValue(targetPokemonEffortB),
	}
}

// Todo: refactor Notification
func (t TrainedPokemonAttackTarget) Id() identifier.TrainedAttackTargetId {
	return t.id
}

func (t TrainedPokemonAttackTarget) TrainedPokemonId() identifier.TrainedPokemonId {
	return t.trainedPokemonId
}

func (t TrainedPokemonAttackTarget) MoveId() identifier.MoveId {
	return t.moveId
}

func (t TrainedPokemonAttackTarget) TargetPokemonId() identifier.PokemonId {
	return t.targetPokemonId
}

func (t TrainedPokemonAttackTarget) TargetPokemonEffortH() value.EffortValue {
	return t.targetPokemonEffortH
}

func (t TrainedPokemonAttackTarget) TargetPokemonEffortB() value.EffortValue {
	return t.targetPokemonEffortB
}

func (t TrainedPokemonAttackTarget) TargetPokemonEffortD() value.EffortValue {
	return t.targetPokemonEffortD
}
