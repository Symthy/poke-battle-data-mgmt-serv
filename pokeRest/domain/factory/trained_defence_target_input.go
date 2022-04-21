package factory

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type TrainedDefenseTargetInput struct {
	id                        uint
	trainedPokemonId          uint
	moveId                    uint
	targetPokemonId           uint
	targetPokemonAbilityId    uint
	targetPokemonNature       string
	targetPokemonEffortValueA int
	targetPokemonEffortValueC int
}

func NewTrainedDefenseTargetInput(
	id, trainedPokemonId, moveId, targetPokemonId uint, targetPokemonNature string,
	targetPokemonAbilityId uint, targetPokemonEffortValueA, trainedPokemonEffortValueC int,
) *TrainedDefenseTargetInput {
	return &TrainedDefenseTargetInput{
		id:                        id,
		trainedPokemonId:          trainedPokemonId,
		moveId:                    moveId,
		targetPokemonId:           targetPokemonId,
		targetPokemonNature:       targetPokemonNature,
		targetPokemonAbilityId:    targetPokemonAbilityId,
		targetPokemonEffortValueA: targetPokemonEffortValueA,
		targetPokemonEffortValueC: trainedPokemonEffortValueC,
	}
}

func NewTrainedDefenseTargetBuilder() *TrainedDefenseTargetInput {
	return &TrainedDefenseTargetInput{}
}

func (i *TrainedDefenseTargetInput) Id(id uint) {
	i.id = id
}
func (i *TrainedDefenseTargetInput) TrainedPokemonId(trainedPokemonId uint) {
	i.trainedPokemonId = trainedPokemonId
}
func (i *TrainedDefenseTargetInput) MoveId(moveId uint) {
	i.moveId = moveId
}
func (i *TrainedDefenseTargetInput) TargetPokemonAbilityId(targetPokemonAbilityId uint) {
	i.targetPokemonAbilityId = targetPokemonAbilityId
}
func (i *TrainedDefenseTargetInput) TargetPokemonNature(targetPokemonNature string) {
	i.targetPokemonNature = targetPokemonNature
}
func (i *TrainedDefenseTargetInput) TargetPokemonEffortValueA(effortValueA int) {
	i.targetPokemonEffortValueA = effortValueA
}
func (i *TrainedDefenseTargetInput) TargetPokemonEffortValueC(effortValueC int) {
	i.targetPokemonEffortValueC = effortValueC
}

func (i TrainedDefenseTargetInput) BuildDomain() (*trainings.TrainedPokemonDefenceTarget, error) {
	id, err := identifier.NewTrainedDefenseTargetId(i.id)
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
	return trainings.NewTrainedPokemonDefenceTarget(
		*id, *trainedPokemonId, *moveId, *pokemonId, i.targetPokemonNature,
		i.targetPokemonEffortValueA, i.targetPokemonEffortValueC)
}
