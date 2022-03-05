package factory

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type TrainedDefenceTargetInput struct {
	id                        uint
	trainedPokemonId          uint
	moveId                    uint
	targetPokemonId           uint
	targetPokemonNature       string
	targetPokemonEffortValueA int
	targetPokemonEffortValueC int
}

func NewTrainedDefenceTargetInput(
	id, trainedPokemonId, moveId, targetPokemonId uint, targetPokemonNature string,
	targetPokemonEffortValueA, trainedPokemonEffortValueC int,
) TrainedDefenceTargetInput {
	return TrainedDefenceTargetInput{
		id, trainedPokemonId, moveId, targetPokemonId, targetPokemonNature,
		targetPokemonEffortValueA, trainedPokemonEffortValueC,
	}
}

func (i TrainedDefenceTargetInput) BuildDomain() (*trainings.TrainedPokemonDefenceTarget, error) {
	id, err := identifier.NewTrainedDefenceTargetId(i.id)
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
