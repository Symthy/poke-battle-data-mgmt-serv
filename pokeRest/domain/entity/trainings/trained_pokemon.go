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
	nickname *string, gender string, description *string, isPrivate bool, userId identifier.UserId,
	pokemonId identifier.PokemonId, nature string, abilityId identifier.AbilityId, heldItemId identifier.HeldItemId,
	effortValueH, effortValueA, effortValueB, effortValueC, effortValueD, effortValueS int,
	moveId1, moveId2, moveId3, moveId4 identifier.MoveId) TrainedPokemon {
	entity := TrainedPokemon{}
	entity.TrainedPokemonParam = TrainedPokemonParam{
		nickname:    nickname,
		gender:      value.Gender(gender),
		description: description,
		isPrivate:   isPrivate,
		userId:      userId,
	}
	entity.TrainedPokemonAdjustment = TrainedPokemonAdjustment{
		pokemonId:    pokemonId,
		nature:       value.NewPokemonNature(nature),
		abilityId:    abilityId,
		heldItemId:   heldItemId,
		effortValueH: value.NewEffortValue(effortValueH),
		effortValueA: value.NewEffortValue(effortValueA),
		effortValueB: value.NewEffortValue(effortValueB),
		effortValueC: value.NewEffortValue(effortValueC),
		effortValueD: value.NewEffortValue(effortValueD),
		effortValueS: value.NewEffortValue(effortValueS),
		moveId1:      moveId1,
		moveId2:      moveId2,
		moveId3:      moveId3,
		moveId4:      moveId4,
	}
	return entity

}

func (t TrainedPokemon) Id() uint {
	return t.TrainedPokemonParam.id.Value()
}
