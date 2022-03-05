package trainings

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type TrainedPokemon struct {
	TrainedPokemonParam
	TrainedPokemonAdjustment
}

func NewTrainedPokemon(param TrainedPokemonParam, adjustment TrainedPokemonAdjustment) TrainedPokemon {
	entity := TrainedPokemon{}
	entity.TrainedPokemonParam = param
	entity.TrainedPokemonAdjustment = adjustment
	return entity
}

// Todo: factory
func NewTrainedPokemonOfUnregistered(
	nickname string, gender value.Gender, description string, isPrivate bool, userId identifier.UserId,
	pokemonId identifier.PokemonId, nature value.PokemonNature, abilityId identifier.AbilityId, heldItemId identifier.HeldItemId,
	effortValues value.EffortValues, moveSet value.PokemonMoveIdSet) TrainedPokemon {
	entity := TrainedPokemon{}
	entity.TrainedPokemonParam = TrainedPokemonParam{
		nickname:    nickname,
		gender:      gender,
		description: description,
		isPrivate:   isPrivate,
		userId:      userId,
	}
	entity.TrainedPokemonAdjustment = TrainedPokemonAdjustment{
		pokemonId:    pokemonId,
		nature:       nature,
		abilityId:    abilityId,
		heldItemId:   heldItemId,
		effortValues: effortValues,
		moveSet:      moveSet,
	}
	return entity

}

func (t TrainedPokemon) Id() uint {
	return t.TrainedPokemonParam.id.Value()
}
