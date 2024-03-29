package conv

import (
	"github.com/Symthy/PokeRest/internal/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/internal/domain/entity/abilities"
	"github.com/Symthy/PokeRest/internal/domain/entity/battles"
	"github.com/Symthy/PokeRest/internal/domain/entity/items"
	"github.com/Symthy/PokeRest/internal/domain/entity/moves"
	"github.com/Symthy/PokeRest/internal/domain/entity/parties"
	"github.com/Symthy/PokeRest/internal/domain/entity/pokemons"
	"github.com/Symthy/PokeRest/internal/domain/entity/trainings"
	"github.com/Symthy/PokeRest/internal/domain/entity/users"
	"github.com/Symthy/PokeRest/internal/infrastructure"
	"github.com/Symthy/PokeRest/internal/infrastructure/repository/dto"
)

func ConvertToDomains[TS infrastructure.ISchema[TD, K, I], TD infrastructure.IDomain[K, I], K infrastructure.IValueId[I], I uint16 | uint64](
	schemas []*TS, toDomainConverter func(*TS) (*TD, error)) ([]*TD, error) {
	domains := make([]*TD, len(schemas))
	for i, schema := range schemas {
		domain, err := toDomainConverter(schema)
		if err != nil {
			return nil, err
		}
		domains[i] = domain
	}
	return domains, nil
}

// Todo: autogen

func ToSchemaUser(domain *users.User) *schema.User {
	builder := dto.NewUserSchemaBuilder()
	domain.Notify(builder)
	return builder.Build()
}

func ToSchemaPokemon(domain *pokemons.Pokemon) *schema.Pokemon {
	builder := dto.NewPokemonSchemaBuilder()
	domain.Notify(builder)
	return builder.Build()
}

func ToSchemaAbility(domain *abilities.Ability) *schema.Ability {
	builder := dto.NewAbilitySchemaBuilder()
	domain.Notify(builder)
	return builder.Build()
}

func ToSchemaMove(domain *moves.Move) *schema.Move {
	builder := dto.NewMoveSchemaBuilder()
	domain.Notify(builder)
	return builder.Build()
}

func ToSchemaHeldItem(domain *items.HeldItem) *schema.HeldItem {
	builder := dto.NewHeldItemSchemaBuilder()
	domain.Notify(builder)
	return builder.Build()
}

func ToSchemaPartyTag(domain *parties.PartyTag) *schema.PartyTag {
	builder := dto.NewPartyTagSchemaBuilder()
	domain.Notify(builder)
	return builder.Build()
}

func ToSchemaTrainedPokemon(domain *trainings.TrainedPokemon) *schema.TrainedPokemon {
	builder := dto.NewTrainedPokemonSchemaBuilder()
	domain.Notify(builder)
	return builder.Build()
}

func ToSchemaTrainedPokemonAdjustment(domain *trainings.TrainedPokemonAdjustment) *schema.TrainedPokemonAdjustment {
	builder := dto.NewTrainedPokemonAdjustmentSchemaBuilder()
	domain.Notify(builder)
	return builder.Build()
}

func ToSchemaTrainedPokemonAttackTarget(domain *trainings.TrainedPokemonAttackTarget) *schema.TrainedPokemonAttackTarget {
	builder := dto.NewTrainedPokemonAttackTargetSchemaBuilder()
	domain.Notify(builder)
	return builder.Build()
}

func ToSchemaTrainedPokemonDefenceTarget(domain *trainings.TrainedPokemonDefenseTarget) *schema.TrainedPokemonDefenseTarget {
	builder := dto.NewTrainedPokemonDefenseTargetSchemaBuilder()
	domain.Notify(builder)
	return builder.Build()
}

func ToSchemaParty(domain *parties.Party) *schema.Party {
	builder := dto.NewPartySchemaBuilder()
	domain.Notify(builder)
	return builder.Build()
}

func ToSchemaPartySeasonResult(domain *parties.PartyBattleResult) *schema.PartySeasonResult {
	builder := dto.NewPartySeasonResultSchemaBuilder()
	domain.Notify(builder)
	return builder.Build()
}

func ToSchemaBattleRecord(domain *battles.BattleRecord) *schema.BattleRecord {
	builder := dto.NewBattleRecordBuilder()
	domain.Notify(builder)
	return builder.Build()
}

func ToSchemaBattleOpponentParty(domain *battles.BattleOpponentParty) *schema.BattleOpponentParty {
	builder := dto.NewBattleOpponentPartySchemaBuilder()
	domain.Notify(builder)
	return builder.Build()
}
