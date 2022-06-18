package factory

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type TrainedPokemonAdjustmentInput struct {
	id           uint64
	pokemonId    uint64
	nature       string
	abilityId    uint64
	heldItemId   uint64
	effortValueH uint64
	effortValueA uint64
	effortValueB uint64
	effortValueC uint64
	effortValueD uint64
	effortValueS uint64
	moveId1      uint64
	moveId2      uint64
	moveId3      uint64
	moveId4      uint64
}

func NewTrainedPokemonAdjustmentInput(
	id, pokemonId uint64, nature string, abilityId, heldItemId uint64,
	effortValueH, effortValueA, effortValueB, effortValueC, effortValueD, effortValueS uint64,
	moveId1, moveId2, moveId3, moveId4 uint64) TrainedPokemonAdjustmentInput {
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
func (i *TrainedPokemonAdjustmentInput) Id(id uint64) {
	i.id = id
}
func (i *TrainedPokemonAdjustmentInput) PokemonId(pokemonId uint64) {
	i.pokemonId = pokemonId
}
func (i *TrainedPokemonAdjustmentInput) Nature(nature string) {
	i.nature = nature
}
func (i *TrainedPokemonAdjustmentInput) AbilityId(abilityId uint64) {
	i.abilityId = abilityId
}
func (i *TrainedPokemonAdjustmentInput) HeldItemId(itemId uint64) {
	i.heldItemId = itemId
}
func (i *TrainedPokemonAdjustmentInput) EffortValueH(value uint64) {
	i.effortValueH = value
}
func (i *TrainedPokemonAdjustmentInput) EffortValueA(value uint64) {
	i.effortValueA = value
}
func (i *TrainedPokemonAdjustmentInput) EffortValueB(value uint64) {
	i.effortValueB = value
}
func (i *TrainedPokemonAdjustmentInput) EffortValueC(value uint64) {
	i.effortValueC = value
}
func (i *TrainedPokemonAdjustmentInput) EffortValueD(value uint64) {
	i.effortValueD = value
}
func (i *TrainedPokemonAdjustmentInput) EffortValueS(value uint64) {
	i.effortValueS = value
}
func (i *TrainedPokemonAdjustmentInput) MoveIdFirst(moveId uint64) {
	i.moveId1 = moveId
}
func (i *TrainedPokemonAdjustmentInput) MoveIdSecond(moveId uint64) {
	i.moveId2 = moveId
}
func (i *TrainedPokemonAdjustmentInput) MoveIdThird(moveId uint64) {
	i.moveId3 = moveId
}
func (i *TrainedPokemonAdjustmentInput) MoveIdFourth(moveId uint64) {
	i.moveId4 = moveId
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
