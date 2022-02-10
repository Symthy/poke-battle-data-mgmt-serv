package infrastructure

import (
	"database/sql"

	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/abilities"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/items"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/moves"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/parties"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/pokemons"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/users"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/optional"
)

func ConvertToDomains[TS ISchema[TD], TD IDomain](schemas []TS) []TD {
	domains := make([]TD, len(schemas), len(schemas))
	for i, s := range schemas {
		domains[i] = s.ConvertToDomain()
	}
	return domains
}

// Todo: autogen

func ToSchemaUser(u users.User) schema.User {
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

func ToSchemaPokemon(p pokemons.Pokemon) schema.Pokemon {
	pokemon := schema.Pokemon{
		ID:               p.Id(),
		PokedexNo:        p.PokedexNo(),
		FormNo:           p.FormNo(),
		FormName:         p.FormName(),
		Name:             p.Name(),
		EnglishName:      p.EnglishName(),
		Generation:       p.Generation(),
		Type1:            enum.PokemonType(p.TypePrimary().NameEN()),
		Type2:            enum.PokemonType(p.TypeSecondary().NameEN()),
		AbilityId1:       ConvertOptionalIdToNullInt16(p.AbilityIdPrimary()),
		AbilityId2:       ConvertOptionalIdToNullInt16(p.AbilityIdSecondary()),
		HiddenAbilityId:  ConvertOptionalIdToNullInt16(p.HiddenAbilityId()),
		IsFinalEvolution: p.IsFinalEvolution(),
	}
	return pokemon
}

func ToSchemaAbility(a abilities.Ability) schema.Ability {
	schema := schema.Ability{}
	return schema
}

func ToSchemaMove(a moves.Move) schema.Move {
	schema := schema.Move{}
	return schema
}

func ToSchemaHeldItem(a items.HeldItem) schema.HeldItem {
	schema := schema.HeldItem{}
	return schema
}

func ToSchemaPartyTag(t parties.PartyTag) schema.PartyTag {
	isGeneration := t.IsGenerationTag()
	isSeason := t.IsSeasonTag()
	return schema.PartyTag{
		Name:            t.Name(),
		IsGenerationTag: &isGeneration,
		IsSeasonTag:     &isSeason,
	}
}

func ToSchemaTrainedPokemon(t trainings.TrainedPokemon) schema.TrainedPokemon {
	schema := schema.TrainedPokemon{}
	return schema
}

func ToSchemaTrainedPokemonAttackTarget(t trainings.TrainedPokemonAttackTarget) schema.TrainedPokemonAttackTarget {
	schema := schema.TrainedPokemonAttackTarget{}
	return schema
}

func ToSchemaTrainedPokemonDeffenceTarget(t trainings.TrainedPokemonDeffenceTarget) schema.TrainedPokemonDeffenceTarget {
	schema := schema.TrainedPokemonDeffenceTarget{}
	return schema
}

func ToSchemaParty(p parties.Party) schema.Party {
	schema := schema.Party{}
	return schema
}

func ToSchemaPartyBattleResult(p parties.PartyBattleResult) schema.PartyBattleResult {
	schema := schema.PartyBattleResult{}
	return schema
}

func ToSchemaBattleRecord(b battles.BattleRecord) schema.BattleRecord {
	schema := schema.BattleRecord{}
	return schema
}

func ToSchemaBattleOpponentParty(b battles.BattleOpponentParty) schema.BattleOpponentParty {
	battleParty := schema.BattleOpponentParty{}
	battleParty.ID = b.Id()
	return battleParty
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
