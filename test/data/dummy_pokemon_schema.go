package data

import (
	"database/sql"

	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
)

// test dummy data
func DummyPokemon3() *schema.Pokemon {
	var abilityId1 = &sql.NullInt16{Int16: 3, Valid: true}
	var abilityId2 = &sql.NullInt16{Int16: 0, Valid: false}
	var abilityId3 = &sql.NullInt16{Int16: 100, Valid: true}
	return &schema.Pokemon{
		PokemonSchema: schema.PokemonSchema{
			ID:               3,
			PokedexNo:        3,
			FormNo:           1,
			FormName:         "Standard",
			Name:             "フシギバナ",
			EnglishName:      "Venusaur",
			Generation:       1,
			Type1:            enum.Grass,
			Type2:            enum.Poison,
			AbilityID1:       *abilityId1,
			AbilityID2:       *abilityId2,
			HiddenAbilityID:  *abilityId3,
			BaseStatsH:       80,
			BaseStatsA:       82,
			BaseStatsB:       83,
			BaseStatsC:       100,
			BaseStatsD:       100,
			BaseStatsS:       80,
			IsFinalEvolution: true,
		},
	}
}
