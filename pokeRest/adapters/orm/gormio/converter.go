package gormio

import (
	"database/sql"

	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/model"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
)

func ConvertDomainToSchema(p model.Pokemon) schema.Pokemon {
	pokemon := schema.Pokemon{
		ID:               p.Id(),
		PokedexNo:        p.PokedexNo(),
		FormNo:           p.FormNo(),
		FormName:         p.FormName(),
		Name:             p.Name(),
		EnglishName:      p.EnglishName(),
		Generation:       p.Generation(),
		Type1:            enum.PokemonType(p.TypePrimary().EnglishName()),
		Type2:            enum.PokemonType(p.TypeSecondary().EnglishName()),
		AbilityId1:       convertOptionalIdToNullInt16(p.AbilityIdPrimary()),
		AbilityId2:       convertOptionalIdToNullInt16(p.AbilityIdSecondary()),
		HiddenAbilityId:  convertOptionalIdToNullInt16(p.HiddenAbilityId()),
		IsFinalEvolution: p.IsFinalEvolution(),
	}
	return pokemon
}

func convertOptionalIdToNullInt16(id value.OptionalId) sql.NullInt16 {
	value, _ := id.Get()
	nullInt := sql.NullInt16{}
	if value == nil {
		nullInt.Scan(nil)
	} else {
		nullInt.Scan(*value)
	}
	return nullInt
}
