package trainings

import (
	"github.com/Symthy/PokeRest/internal/domain/entity"
	"github.com/Symthy/PokeRest/internal/domain/value"
	"github.com/Symthy/PokeRest/internal/domain/value/identifier"
)

var _ entity.IDomain[identifier.TrainedAttackTargetId, uint64] = (*TrainedPokemonAttackTarget)(nil)

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
	targetPokemonEffortH, targetPokemonEffortB, targetPokemonEffortD uint64,
) (*TrainedPokemonAttackTarget, error) {
	effortValueH, err := value.NewEffortValue(targetPokemonEffortH, value.H)
	if err != nil {
		return nil, err
	}
	effortValueB, err := value.NewEffortValue(targetPokemonEffortH, value.B)
	if err != nil {
		return nil, err
	}
	effortValueD, err := value.NewEffortValue(targetPokemonEffortH, value.D)
	if err != nil {
		return nil, err
	}
	return &TrainedPokemonAttackTarget{
		id:                   id,
		trainedPokemonId:     trainedPokemonId,
		moveId:               moveId,
		targetPokemonId:      targetPokemonId,
		targetPokemonNature:  value.NewPokemonNature(targetPokemonNature),
		targetPokemonEffortH: *effortValueH,
		targetPokemonEffortB: *effortValueB,
		targetPokemonEffortD: *effortValueD,
	}, nil
}

// Todo: refactor Notification
func (t TrainedPokemonAttackTarget) Id() identifier.TrainedAttackTargetId {
	return t.id
}

func (t TrainedPokemonAttackTarget) TrainedPokemonId() identifier.TrainedPokemonId {
	return t.trainedPokemonId
}

func (t TrainedPokemonAttackTarget) Notify(note ITrainedPokemonAttackNotification) {
	note.SetId(t.id)
	note.SetTargetPokemonId(t.targetPokemonId)
	note.SetMoveId(t.moveId)
	note.SetTargetPokemonNature(t.targetPokemonNature)
	note.SetTargetPokemonId(t.targetPokemonId)
	note.SetTargetPokemonEffortH(t.targetPokemonEffortH)
	note.SetTargetPokemonEffortB(t.targetPokemonEffortB)
	note.SetTargetPokemonEffortD(t.targetPokemonEffortD)
}
