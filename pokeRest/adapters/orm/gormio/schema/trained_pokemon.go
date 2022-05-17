package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/mixin"
)

type TrainedPokemon struct {
	ID                       uint64      `gorm:"primaryKey;autoIncrement:true"`
	Gender                   enum.Gender `sql:"type:gender"`
	Nickname                 *string
	Description              *string
	AdjustmentID             uint64
	TrainedPokemonAdjustment TrainedPokemonAdjustment `gorm:"foreignKey:AdjustmentID;references:id;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
	IsPrivate                bool                     `gorm:"default:false"`
	CreateUserId             *uint64                  // M:1 from User
	mixin.UpdateTimes

	// relation
	BattleRecord1 []BattleRecord `gorm:"foreignKey:SelfTrainedPokemonId1;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 1:M -> BattleRecord
	BattleRecord2 []BattleRecord `gorm:"foreignKey:SelfTrainedPokemonId2;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 1:M -> BattleRecord
	BattleRecord3 []BattleRecord `gorm:"foreignKey:SelfTrainedPokemonId3;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 1:M -> BattleRecord
	BattleRecord4 []BattleRecord `gorm:"foreignKey:SelfTrainedPokemonId4;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 1:M -> BattleRecord
	Party         []*Party       `gorm:"many2many:trained_pokemons_parties;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`              // M:M
}

func (TrainedPokemon) TableName() string {
	return "trained_pokemons"
}
