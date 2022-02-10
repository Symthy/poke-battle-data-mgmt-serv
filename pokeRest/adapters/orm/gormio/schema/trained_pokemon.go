package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/trainings"
	"gorm.io/gorm"
)

type TrainedPokemon struct {
	gorm.Model
	PokemonId    uint
	Pokemon      Pokemon // belongs to
	Nickname     *string
	Gender       *enum.Gender `sql:"type:gender"`
	Description  *string
	Nature       enum.Nature
	AbilityId    uint  // M:1 <- Ability
	HeldItemId   *uint // M:1 <- HeldItem
	EffortValueH int   `gorm:"default:0"`
	EffortValueA int   `gorm:"default:0"`
	EffortValueB int   `gorm:"default:0"`
	EffortValueC int   `gorm:"default:0"`
	EffortValueD int   `gorm:"default:0"`
	EffortValueS int   `gorm:"default:0"`
	MoveId1      *uint // M:1 <- Move
	MoveId2      *uint // M:1 <- Move
	MoveId3      *uint // M:1 <- Move
	MoveId4      *uint // M:1 <- Move
	IsPrivate    bool  `gorm:"default:false"`
	// belong to
	CreateUserId uint // M:1 from User

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

func (t TrainedPokemon) ConvertToDomain() trainings.TrainedPokemon {
	return trainings.NewTrainedPokemon(t.ID)
}
