package factory

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type TrainedPokemonInput struct {
	id          uint
	gender      string
	nickname    string
	description string
	isPrivate   bool
	userId      uint
	adjustment  TrainedPokemonAdjustmentInput
}

func NewTrainedPokemonInput(
	id uint, gender string, nickname string, description string, adjustmentId uint, isPrivate bool,
	userId uint, adjustment TrainedPokemonAdjustmentInput) TrainedPokemonInput {
	return TrainedPokemonInput{
		id:          id,
		gender:      gender,
		nickname:    nickname,
		description: description,
		isPrivate:   isPrivate,
		userId:      userId,
	}
}

func (i TrainedPokemonInput) BuildDomain() (*trainings.TrainedPokemon, error) {
	id, err := identifier.NewTrainedPokemonId(i.id)
	if err != nil {
		return nil, err
	}
	gender := value.Gender(i.gender)
	userId, err := identifier.NewUserId(i.userId)
	if err != nil {
		return nil, err
	}
	adjustment, err := i.adjustment.BuildDomain()
	if err != nil {
		return nil, err
	}
	domain := trainings.NewTrainedPokemon(*id, gender, i.nickname, i.description,
		i.isPrivate, *userId, *adjustment)
	return &domain, nil
}
