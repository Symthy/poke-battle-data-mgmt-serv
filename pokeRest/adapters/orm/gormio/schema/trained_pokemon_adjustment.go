package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
	"gorm.io/gorm"
)

type TrainedPokemonAdjustment struct {
	gorm.Model
	PokemonId    uint
	Pokemon      Pokemon // belongs to
	Nature       enum.Nature
	AbilityId    *int // M:1 <- Ability
	HeldItemId   *int // M:1 <- HeldItem
	EffortValueH int  `gorm:"default:0"`
	EffortValueA int  `gorm:"default:0"`
	EffortValueB int  `gorm:"default:0"`
	EffortValueC int  `gorm:"default:0"`
	EffortValueD int  `gorm:"default:0"`
	EffortValueS int  `gorm:"default:0"`
	MoveId1      *int // M:1 <- Move
	MoveId2      *int // M:1 <- Move
	MoveId3      *int // M:1 <- Move
	MoveId4      *int // M:1 <- Move
}

func (TrainedPokemonAdjustment) TableName() string {
	return "trained_pokemon_adjustments"
}

func (t TrainedPokemonAdjustment) ConvertToDomain() trainings.TrainedPokemonAdjustment {
	return trainings.NewTrainedPokemonAdjustment(
		t.ID, int(t.PokemonId), t.Nature.String(), t.AbilityId, t.HeldItemId, t.EffortValueH,
		t.EffortValueA, t.EffortValueB, t.EffortValueC, t.EffortValueD, t.EffortValueS,
		t.MoveId1, t.MoveId2, t.MoveId3, t.MoveId4)
}
