package trainings

import (
	"github.com/Symthy/PokeRest/internal/domain/entity"
	"github.com/Symthy/PokeRest/internal/domain/value"
	"github.com/Symthy/PokeRest/internal/domain/value/identifier"
)

var _ entity.IDomain[identifier.TrainedDefenseTargetId, uint64] = (*TrainedPokemonDefenseTarget)(nil)

type TrainedPokemonDefenseTarget struct {
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
	targetPokemonEffortA, targetPokemonEffortC uint64,
) (*TrainedPokemonDefenseTarget, error) {
	effortValueA, err := value.NewEffortValue(targetPokemonEffortA, value.A)
	if err != nil {
		return nil, err
	}
	effortValueC, err := value.NewEffortValue(targetPokemonEffortC, value.C)
	if err != nil {
		return nil, err
	}
	return &TrainedPokemonDefenseTarget{
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
func (t TrainedPokemonDefenseTarget) Id() identifier.TrainedDefenseTargetId {
	return t.id
}

func (t TrainedPokemonDefenseTarget) TrainedPokemonId() identifier.TrainedPokemonId {
	return t.trainedPokemonId
}

func (t TrainedPokemonDefenseTarget) Notify(note ITrainedPokemonDefenseNotification) {
	note.SetId(t.id)
	note.SetTargetPokemonId(t.targetPokemonId)
	note.SetMoveId(t.moveId)
	note.SetTargetPokemonNature(t.targetPokemonNature)
	note.SetTargetPokemonId(t.targetPokemonId)
	note.SetTargetPokemonEffortA(t.targetPokemonEffortA)
	note.SetTargetPokemonEffortC(t.targetPokemonEffortC)
}
