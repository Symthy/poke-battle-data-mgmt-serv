package trainings

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

var _ entity.IDomain[identifier.TrainedDefenceTargetId] = (*TrainedPokemonDefenceTarget)(nil)

type TrainedPokemonDefenceTarget struct {
	id                   identifier.TrainedDefenceTargetId
	trainedPokemonId     identifier.TrainedPokemonId
	moveId               identifier.MoveId
	targetPokemonId      identifier.PokemonId
	targetPokemonNature  value.PokemonNature
	targetPokemonEffortA value.EffortValue
	targetPokemonEffortC value.EffortValue
}

func NewTrainedPokemonDefenceTarget(
	id identifier.TrainedDefenceTargetId, trainedPokemonId identifier.TrainedPokemonId,
	moveId identifier.MoveId, targetPokemonId identifier.PokemonId, targetPokemonNature string,
	targetPokemonEffortA int, targetPokemonEffortC int) TrainedPokemonDefenceTarget {
	return TrainedPokemonDefenceTarget{
		id:                   id,
		trainedPokemonId:     trainedPokemonId,
		moveId:               moveId,
		targetPokemonId:      targetPokemonId,
		targetPokemonNature:  value.NewPokemonNature(targetPokemonNature),
		targetPokemonEffortA: value.NewEffortValue(targetPokemonEffortA),
		targetPokemonEffortC: value.NewEffortValue(targetPokemonEffortC),
	}
}

// Todo: refactor Notification
func (t TrainedPokemonDefenceTarget) Id() identifier.TrainedDefenceTargetId {
	return t.id
}

func (t TrainedPokemonDefenceTarget) TrainedPokemonId() identifier.TrainedPokemonId {
	return t.trainedPokemonId
}

func (t TrainedPokemonDefenceTarget) MoveId() identifier.MoveId {
	return t.moveId
}

func (t TrainedPokemonDefenceTarget) TargetPokemonId() identifier.PokemonId {
	return t.targetPokemonId
}

func (t TrainedPokemonDefenceTarget) TargetPokemonEffortA() value.EffortValue {
	return t.targetPokemonEffortA
}

func (t TrainedPokemonDefenceTarget) TargetPokemonEffortC() value.EffortValue {
	return t.targetPokemonEffortC
}
