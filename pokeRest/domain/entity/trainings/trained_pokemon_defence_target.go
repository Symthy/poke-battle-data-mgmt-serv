package trainings

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

var _ entity.IDomain[identifier.TrainedDefenseTargetId] = (*TrainedPokemonDefenceTarget)(nil)

type TrainedPokemonDefenceTarget struct {
	id                   identifier.TrainedDefenseTargetId
	trainedPokemonId     identifier.TrainedPokemonId
	moveId               identifier.MoveId
	targetPokemonId      identifier.PokemonId
	targetPokemonNature  value.PokemonNature
	targetPokemonEffortA value.EffortValue
	targetPokemonEffortC value.EffortValue
}

func NewTrainedPokemonDefenceTarget(
	id identifier.TrainedDefenseTargetId, trainedPokemonId identifier.TrainedPokemonId,
	moveId identifier.MoveId, targetPokemonId identifier.PokemonId, targetPokemonNature string,
	targetPokemonEffortA int, targetPokemonEffortC int,
) (*TrainedPokemonDefenceTarget, error) {
	effortValueA, err := value.NewEffortValue(targetPokemonEffortA, value.A)
	if err != nil {
		return nil, err
	}
	effortValueC, err := value.NewEffortValue(targetPokemonEffortC, value.C)
	if err != nil {
		return nil, err
	}
	return &TrainedPokemonDefenceTarget{
		id:                   id,
		trainedPokemonId:     trainedPokemonId,
		moveId:               moveId,
		targetPokemonId:      targetPokemonId,
		targetPokemonNature:  value.NewPokemonNature(targetPokemonNature),
		targetPokemonEffortA: *effortValueA,
		targetPokemonEffortC: *effortValueC,
	}, nil
}

// Todo: refactor Notification
func (t TrainedPokemonDefenceTarget) Id() identifier.TrainedDefenseTargetId {
	return t.id
}

func (t TrainedPokemonDefenceTarget) TrainedPokemonId() identifier.TrainedPokemonId {
	return t.trainedPokemonId
}

func (t TrainedPokemonDefenceTarget) Notify(note ITrainedPokemonDefenceNotification) {
	note.SetId(t.id)
	note.SetTargetPokemonId(t.targetPokemonId)
	note.SetMoveId(t.moveId)
	note.SetTargetPokemonNature(t.targetPokemonNature)
	note.SetTargetPokemonId(t.targetPokemonId)
	note.SetTargetPokemonEffortA(t.targetPokemonEffortA)
	note.SetTargetPokemonEffortC(t.targetPokemonEffortC)
}
