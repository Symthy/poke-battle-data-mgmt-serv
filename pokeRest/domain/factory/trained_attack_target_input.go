package factory

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type TrainedAttackTargetInput struct {
	id                        uint
	trainedPokemonId          uint
	moveId                    uint
	targetPokemonId           uint
	targetPokemonNature       string
	targetPokemonEffortValueH int
	targetPokemonEffortValueB int
	targetPokemonEffortValueD int
}

func NewTrainedAttackTargetInput(
	id, trainedPokemonId, moveId, targetPokemonId uint, targetPokemonNature string,
	targetPokemonEffortValueH, trainedPokemonEffortValueB, trainedPokemonEffortValueD int,
) TrainedAttackTargetInput {
	return TrainedAttackTargetInput{
		id, trainedPokemonId, moveId, targetPokemonId, targetPokemonNature,
		targetPokemonEffortValueH, trainedPokemonEffortValueB, trainedPokemonEffortValueD,
	}
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
	return trainings.NewTrainedPokemonAttackTarget(
		*id, *trainedPokemonId, *moveId, *pokemonId, i.targetPokemonNature,
		i.targetPokemonEffortValueH, i.targetPokemonEffortValueB, i.targetPokemonEffortValueD)
}
