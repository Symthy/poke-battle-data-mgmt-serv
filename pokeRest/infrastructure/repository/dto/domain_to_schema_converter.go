package dto

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
	"github.com/Symthy/PokeRest/pokeRest/infrastructure"
)

func ConvertToDomains[TS infrastructure.ISchema[TD, K], TD infrastructure.IDomain[K], K infrastructure.IValueId](schemas []TS) []TD {
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
	user.ID = u.Id().Value()
	return user
}

func ToSchemaPokemon(domain pokemons.Pokemon) schema.Pokemon {
	builder := PokemonSchemaBuilder{}
	domain.Notify(&builder)
	return builder.Build()
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

func ToSchemaPartyTag(domain parties.PartyTag) schema.PartyTag {
	builder := PartyTagSchemaBuilder{}
	domain.Notify(&builder)
	return builder.Build()
}

func ToSchemaTrainedPokemon(t trainings.TrainedPokemonParam) schema.TrainedPokemon {
	schema := schema.TrainedPokemon{}
	return schema
}

func ToSchemaTrainedPokemonAdjustment(t trainings.TrainedPokemonAdjustment) schema.TrainedPokemonAdjustment {
	schema := schema.TrainedPokemonAdjustment{}
	return schema
}

func ToSchemaTrainedPokemonAttackTarget(t trainings.TrainedPokemonAttackTarget) schema.TrainedPokemonAttackTarget {
	schema := schema.TrainedPokemonAttackTarget{}
	return schema
}

func ToSchemaTrainedPokemonDeffenceTarget(t trainings.TrainedPokemonDefenceTarget) schema.TrainedPokemonDeffenceTarget {
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

func ToSchemaBattleRecord(domain battles.BattleRecord) schema.BattleRecord {
	builder := NewBattleRecordBuilder()
	domain.Notify(&builder)
	return builder.Build()
}

func ToSchemaBattleOpponentParty(b battles.BattleOpponentParty) schema.BattleOpponentParty {
	battleParty := schema.BattleOpponentParty{}
	battleParty.ID = b.Id().Value()
	return battleParty
}

func convertIdToNullInt16(id infrastructure.IValueId) sql.NullInt16 {
	value := id.Value()
	nullInt := sql.NullInt16{}
	if value == 0 {
		nullInt.Scan(nil)
	} else {
		nullInt.Scan(value)
	}
	return nullInt
}
