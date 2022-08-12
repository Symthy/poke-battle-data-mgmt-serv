package factory

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type TrainedDefenseTargetInput struct {
	id                        uint64
	trainedPokemonId          uint64
	moveId                    uint64
	targetPokemonId           uint64
	targetPokemonNature       string
	targetPokemonAbilityId    uint64
	targetPokemonHeldItemId   uint64
	targetPokemonEffortValueA uint64
	targetPokemonEffortValueC uint64
}

func NewTrainedDefenseTargetBuilder() *TrainedDefenseTargetInput {
	return &TrainedDefenseTargetInput{}
}

func (i *TrainedDefenseTargetInput) Id(id uint64) {
	i.id = id
}
func (i *TrainedDefenseTargetInput) TrainedPokemonId(trainedPokemonId uint64) {
	i.trainedPokemonId = trainedPokemonId
}
func (i *TrainedDefenseTargetInput) MoveId(moveId uint64) {
	i.moveId = moveId
}
func (i *TrainedDefenseTargetInput) TargetPokemonId(targetPokemonId uint64) {
	i.targetPokemonId = targetPokemonId
}
func (i *TrainedDefenseTargetInput) TargetPokemonNature(targetPokemonNature string) {
	i.targetPokemonNature = targetPokemonNature
}
func (i *TrainedDefenseTargetInput) TargetPokemonAbilityId(targetPokemonAbilityId uint64) {
	i.targetPokemonAbilityId = targetPokemonAbilityId
}
func (i *TrainedDefenseTargetInput) TargetPokemonHeldItemId(targetPokemonHeldItemId uint64) {
	i.targetPokemonHeldItemId = targetPokemonHeldItemId
}
func (i *TrainedDefenseTargetInput) TargetPokemonEffortValueA(effortValueA uint64) {
	i.targetPokemonEffortValueA = effortValueA
}
func (i *TrainedDefenseTargetInput) TargetPokemonEffortValueC(effortValueC uint64) {
	i.targetPokemonEffortValueC = effortValueC
}

func (i TrainedDefenseTargetInput) BuildDomain() (*trainings.TrainedPokemonDefenseTarget, error) {
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
