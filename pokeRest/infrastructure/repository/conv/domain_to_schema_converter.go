package conv

import (
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
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/repository/dto"
)

func ConvertToDomains[TS infrastructure.ISchema[TD, K, I], TD infrastructure.IDomain[K, I], K infrastructure.IValueId[I], I uint16 | uint64](
	schemas []TS, toDomainConverter func(TS) (*TD, error)) ([]TD, error) {
	domains := make([]TD, len(schemas), len(schemas))
	for i, schema := range schemas {
		domain, err := toDomainConverter(schema)
		if err != nil {
			return nil, err
		}
		domains[i] = *domain
	}
	return domains, nil
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
	builder := dto.PokemonSchemaBuilder{}
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
	builder := dto.PartyTagSchemaBuilder{}
	domain.Notify(&builder)
	return builder.Build()
}

func ToSchemaTrainedPokemon(t trainings.TrainedPokemon) schema.TrainedPokemon {
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

func ToSchemaTrainedPokemonDefenceTarget(t trainings.TrainedPokemonDefenceTarget) schema.TrainedPokemonDefenceTarget {
	schema := schema.TrainedPokemonDefenceTarget{}
	return schema
}

func ToSchemaParty(p parties.Party) schema.Party {
	schema := schema.Party{}
	return schema
}

func ToSchemaPartySeasonResult(p parties.PartyBattleResult) schema.PartySeasonResult {
	schema := schema.PartySeasonResult{}
	return schema
}

func ToSchemaBattleRecord(domain battles.BattleRecord) schema.BattleRecord {
	builder := dto.BattleRecordSchemaBuilder{}
	domain.Notify(&builder)
	return builder.Build()
}

func ToSchemaBattleOpponentParty(domain battles.BattleOpponentParty) schema.BattleOpponentParty {
	builder := dto.BattleOpponentPartySchemaBuilder{}
	domain.Notify(&builder)
	return builder.Build()
}
