package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
	"gorm.io/gorm"
)

type TrainedPokemon struct {
	gorm.Model
	TrainedPokemonAdjustmentId uint
	TrainedPokemonAdjustment   TrainedPokemonAdjustment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
	Nickname                   *string
	Gender                     *enum.Gender `sql:"type:gender"`
	Description                *string
	IsPrivate                  bool `gorm:"default:false"`
	CreateUserId               uint // M:1 from User

	// relation
	BattleRecord1 []BattleRecord `gorm:"foreignKey:SelfTrainedPokemonId1;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 1:M -> BattleRecord
	BattleRecord2 []BattleRecord `gorm:"foreignKey:SelfTrainedPokemonId2;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 1:M -> BattleRecord
	BattleRecord3 []BattleRecord `gorm:"foreignKey:SelfTrainedPokemonId3;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 1:M -> BattleRecord
	BattleRecord4 []BattleRecord `gorm:"foreignKey:SelfTrainedPokemonId4;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 1:M -> BattleRecord
	Party         []*Party       `gorm:"many2many:trained_pokemon_parties;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`               // M:M
}

func (TrainedPokemon) TableName() string {
	return "trained_pokemons"
}

func (t TrainedPokemon) ConvertToDomain() trainings.TrainedPokemonParam {
	return trainings.NewTrainedPokemonParam(t.ID)
}
