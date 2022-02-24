package service

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

type TrainedPokemonResolver struct {
	adjustmentRepo repository.ITrainedPokemonAdjustmentRepository
}

func NewTrainedPokemonResolver(adjustmentRepo repository.ITrainedPokemonAdjustmentRepository) TrainedPokemonResolver {
	return TrainedPokemonResolver{adjustmentRepo: adjustmentRepo}
}

func (r TrainedPokemonResolver) Resolve(input trainings.TrainedPokemon) (*trainings.TrainedPokemon, error) {
	trainedPokemon := input
	adjustment, err := r.adjustmentRepo.Find(trainedPokemon.TrainedPokemonAdjustment)
	if err != nil {
		return nil, err
	}

	if adjustment == nil {
		createdAdjustment, err := r.adjustmentRepo.Create(trainedPokemon.TrainedPokemonAdjustment)
		if err != nil {
			return nil, err
		}
		trainedPokemon.TrainedPokemonParam.ApplyAdjustmentId(createdAdjustment.Id())
		adjustment = createdAdjustment
	}
	return &trainedPokemon, nil
}
