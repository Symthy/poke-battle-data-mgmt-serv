package factory

import (
	"github.com/Symthy/PokeRest/internal/domain/entity/trainings"
	"github.com/Symthy/PokeRest/internal/domain/value/identifier"
)

type TrainedAttackTargetInput struct {
	id                        uint64
	trainedPokemonId          uint64
	moveId                    uint64
	targetPokemonId           uint64
	targetPokemonNature       string
	targetPokemonAbilityId    uint64
	targetPokemonHeldItemId   uint64
	targetPokemonEffortValueH uint64
	targetPokemonEffortValueB uint64
	targetPokemonEffortValueD uint64
}

func NewTrainedAttackTargetBuilder() *TrainedAttackTargetInput {
	return &TrainedAttackTargetInput{}
}

func (i *TrainedAttackTargetInput) Id(id uint64) {
	i.id = id
}
func (i *TrainedAttackTargetInput) TrainedPokemonId(trainedPokemonId uint64) {
	i.trainedPokemonId = trainedPokemonId
}
func (i *TrainedAttackTargetInput) MoveId(moveId uint64) {
	i.moveId = moveId
}
func (i *TrainedAttackTargetInput) TargetPokemonId(targetPokemonId uint64) {
	i.targetPokemonId = targetPokemonId
}
func (i *TrainedAttackTargetInput) TargetPokemonNature(targetPokemonNature string) {
	i.targetPokemonNature = targetPokemonNature
}
func (i *TrainedAttackTargetInput) TargetPokemonAbilityId(targetPokemonAbilityId uint64) {
	i.targetPokemonAbilityId = targetPokemonAbilityId
}
func (i *TrainedAttackTargetInput) TargetPokemonHeldItemId(targetPokemonHeldItemId uint64) {
	i.targetPokemonHeldItemId = targetPokemonHeldItemId
}
func (i *TrainedAttackTargetInput) TargetPokemonEffortValueH(effortValueH uint64) {
	i.targetPokemonEffortValueH = effortValueH
}
func (i *TrainedAttackTargetInput) TargetPokemonEffortValueB(effortValueB uint64) {
	i.targetPokemonEffortValueB = effortValueB
}
func (i *TrainedAttackTargetInput) TargetPokemonEffortValueD(effortValueD uint64) {
	i.targetPokemonEffortValueD = effortValueD
}

func (i TrainedAttackTargetInput) BuildDomain() (*trainings.TrainedPokemonAttackTarget, error) {
	id, err := identifier.NewTrainedAttackTargetId(i.id)
	if err != nil {
		return nil, err
	}
	trainedPokemonId, err := identifier.NewTrainedPokemonId(i.trainedPokemonId)
	if err != nil {
		return nil, err
	}
	moveId, err := identifier.NewMoveId(i.moveId)
	if err != nil {
		return nil, err
	}
	pokemonId, err := identifier.NewPokemonId(i.targetPokemonId)
	if err != nil {
		return nil, err
	}
	return trainings.NewTrainedPokemonAttackTarget(
		*id, *trainedPokemonId, *moveId, *pokemonId, i.targetPokemonNature,
		i.targetPokemonEffortValueH, i.targetPokemonEffortValueB, i.targetPokemonEffortValueD)
}
