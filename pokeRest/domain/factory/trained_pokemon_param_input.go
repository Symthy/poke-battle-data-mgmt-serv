package factory

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type TrainedPokemonParamInput struct {
	id           uint
	gender       string
	nickname     string
	description  string
	adjustmentId uint
	isPrivate    bool
	userId       uint
}

func NewTrainedPokemonParamInput(
	id uint, gender string, nickname string, description string, adjustmentId uint, isPrivate bool,
	userId uint) TrainedPokemonParamInput {
	return TrainedPokemonParamInput{
		id:           id,
		gender:       gender,
		nickname:     nickname,
		description:  description,
		adjustmentId: adjustmentId,
		isPrivate:    isPrivate,
		userId:       userId,
	}
}

func (i TrainedPokemonParamInput) BuildDomain() (*trainings.TrainedPokemonParam, error) {
	id, err := identifier.NewTrainedPokemonId(i.id)
	if err != nil {
		return nil, err
	}
	gender := value.Gender(i.gender)
	adjustmentId, err := identifier.NewTrainedAdjustmentId(i.adjustmentId)
	if err != nil {
		return nil, err
	}
	userId, err := identifier.NewUserId(i.userId)
	if err != nil {
		return nil, err
	}

	domain := trainings.NewTrainedPokemonParam(*id, gender, i.nickname, i.description, *adjustmentId,
		i.isPrivate, *userId)
	return &domain, nil
}
