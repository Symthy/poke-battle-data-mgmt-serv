package factory

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type TrainedPokemonAdjustmentInput struct {
	id           uint
	pokemonId    uint
	nature       string
	abilityId    uint
	heldItemId   uint
	effortValueH int
	effortValueA int
	effortValueB int
	effortValueC int
	effortValueD int
	effortValueS int
	moveId1      uint
	moveId2      uint
	moveId3      uint
	moveId4      uint
}

func NewTrainedPokemonAdjustmentInput(
	id uint, pokemonId uint, nature string, abilityId uint, heldItemId uint, effortValueH int,
	effortValueA int, effortValueB int, effortValueC int, effortValueD int, effortValueS int,
	moveId1 uint, moveId2 uint, moveId3 uint, moveId4 uint) TrainedPokemonAdjustmentInput {
	return TrainedPokemonAdjustmentInput{
		id:           id,
		pokemonId:    pokemonId,
		nature:       nature,
		abilityId:    abilityId,
		heldItemId:   heldItemId,
		effortValueH: effortValueH,
		effortValueA: effortValueA,
		effortValueB: effortValueB,
		effortValueC: effortValueC,
		effortValueD: effortValueD,
		effortValueS: effortValueS,
		moveId1:      moveId1,
		moveId2:      moveId2,
		moveId3:      moveId3,
		moveId4:      moveId4,
	}
}

func NewTrainedPokemonAdjustmentBuilder() TrainedPokemonAdjustmentInput {
	return TrainedPokemonAdjustmentInput{}
}

func (i TrainedPokemonAdjustmentInput) BuildDomain() (*trainings.TrainedPokemonAdjustment, error) {
	id, err := identifier.NewTrainedAdjustmentId(i.id)
	if err != nil {
		return nil, err
	}
	pokemonId, err := identifier.NewPokemonId(i.pokemonId)
	if err != nil {
		return nil, err
	}
	nature := value.NewPokemonNature(i.nature)
	abilityId, err := identifier.NewAbilityId(i.abilityId)
	if err != nil {
		return nil, err
	}
	heldItemId, err := identifier.NewHeldItemId(i.heldItemId)
	if err != nil {
		return nil, err
	}
	effortValues, err := value.NewEffortValues(i.effortValueH, i.effortValueA, i.effortValueB, i.effortValueC, i.effortValueD, i.effortValueS)
	if err != nil {
		return nil, err
	}
	moveSet, err := value.NewPokemonMoveIdSet(i.moveId1, i.moveId2, i.moveId3, i.moveId4)
	if err != nil {
		return nil, err
	}
	domain := trainings.NewTrainedPokemonAdjustment(*id, *pokemonId, nature, *abilityId, *heldItemId, *effortValues, *moveSet)
	return &domain, nil
}
