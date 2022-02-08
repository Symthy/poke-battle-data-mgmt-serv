package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/pokemons"
	"gorm.io/gorm"
)

type TrainedPokemon struct {
	gorm.Model
	PokemonId            uint
	Pokemon              Pokemon // belongs to
	Nickname             *string
	Gender               *enum.Gender `sql:"type:gender"`
	Description          *string
	TrainedPokemonBaseId uint
	TrainedPokemonBase   TrainedPokemonBase `gorm:"constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
	// belong to
	CreateUserId uint // M:1 from User

	// relation
	BattleRecord1 []BattleRecord `gorm:"foreignKey:SelfTrainedPokemonId1;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 1:M -> BattleRecord
	BattleRecord2 []BattleRecord `gorm:"foreignKey:SelfTrainedPokemonId2;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 1:M -> BattleRecord
	BattleRecord3 []BattleRecord `gorm:"foreignKey:SelfTrainedPokemonId3;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 1:M -> BattleRecord
	BattleRecord4 []BattleRecord `gorm:"foreignKey:SelfTrainedPokemonId4;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 1:M -> BattleRecord
}

func (TrainedPokemon) TableName() string {
	return "trained_pokemons"
}

func (t TrainedPokemon) ConvertToDomain() pokemons.TrainedPokemon {
	return pokemons.NewTrainedPokemon(t.ID)
}
