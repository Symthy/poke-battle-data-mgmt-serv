package data

import (
	"database/sql"

	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
)

// test dummy data
func DummyPokemon3() schema.Pokemon {
	var abilityId1 = &sql.NullInt16{Int16: 3, Valid: true}
	var abilityId2 = &sql.NullInt16{Int16: 0, Valid: false}
	var abilityId3 = &sql.NullInt16{Int16: 100, Valid: true}
	return schema.Pokemon{
		ID:               3,
		PokedexNo:        3,
		FormNo:           1,
		FormName:         "Standard",
		Name:             "フシギバナ",
		EnglishName:      "Venusaur",
		Generation:       1,
		Type1:            enum.Grass,
		Type2:            enum.Poison,
		AbilityId1:       *abilityId1,
		AbilityId2:       *abilityId2,
		HiddenAbilityId:  *abilityId3,
		IsFinalEvolution: true,
	}
}
