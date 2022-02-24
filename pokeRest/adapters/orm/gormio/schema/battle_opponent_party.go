package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/common/collections"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
)

type BattleOpponentParty struct {
	ID                 uint `gorm:"primaryKey;autoIncrement:true"`
	OpponentPokemonId1 *int // M:1 <- Pokemon
	OpponentPokemonId2 *int
	OpponentPokemonId3 *int
	OpponentPokemonId4 *int
	OpponentPokemonId5 *int
	OpponentPokemonId6 *int
	// relation
	BattleRecord []BattleRecord `gorm:"constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"` // 1:M -> BattleRecord
}

func (BattleOpponentParty) TableName() string {
	return "battle_opponent_party"
}

func (b BattleOpponentParty) ConvertToDomain() battles.BattleOpponentParty {
	pokemonIds := []int{}
	collections.AddsToList(pokemonIds, b.OpponentPokemonId1, b.OpponentPokemonId2,
		b.OpponentPokemonId3, b.OpponentPokemonId4, b.OpponentPokemonId5, b.OpponentPokemonId6)
	return battles.NewBattleOpponentParty(b.ID, value.NewPartyPokemonIds(pokemonIds...))
}
