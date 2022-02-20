package trainings

import "github.com/Symthy/PokeRest/pokeRest/domain/value"

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

// Todo: builder
func NewTrainedPokemonOfUnregistered(
	nickname *string, gender string, description *string, isPrivate bool, userId *uint,
	pokemonId int, nature string, abilityId *int, heldItemId *int, effortValueH int, effortValueA int,
	effortValueB int, effortValueC int, effortValueD int, effortValueS int,
	moveId1 *int, moveId2 *int, moveId3 *int, moveId4 *int) TrainedPokemon {
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
	return t.TrainedPokemonParam.id
}
