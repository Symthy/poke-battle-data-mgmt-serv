package orm

import (
	"database/sql"

	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/model"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/optional"
)

func ToSchemaPokemon(p model.Pokemon) *schema.Pokemon {
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
		AbilityId1:       ConvertOptionalIdToNullInt16(p.AbilityIdPrimary()),
		AbilityId2:       ConvertOptionalIdToNullInt16(p.AbilityIdSecondary()),
		HiddenAbilityId:  ConvertOptionalIdToNullInt16(p.HiddenAbilityId()),
		IsFinalEvolution: p.IsFinalEvolution(),
	}
	return &pokemon
}

func ConvertUserToSchema(u model.User) schema.User {
	email := u.Email().Value()
	user := schema.User{
		Name:        u.Name().Value(),
		DisplayName: u.DisplayName(),
		Email:       &email,
		Profile:     u.Profile(),
		Role:        enum.Role(u.Role().String()),
	}
	user.ID = u.Id()
	return user
}

func ConvertOptionalIdToNullInt16(id optional.OptionalId) sql.NullInt16 {
	value, _ := id.Get()
	nullInt := sql.NullInt16{}
	if value == nil {
		nullInt.Scan(nil)
	} else {
		nullInt.Scan(*value)
	}
	return nullInt
}
