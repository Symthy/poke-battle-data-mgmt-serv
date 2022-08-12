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
	schemas []*TS, toDomainConverter func(*TS) (*TD, error)) ([]*TD, error) {
	domains := make([]*TD, len(schemas), len(schemas))
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

func ToSchemaUser(u *users.User) *schema.User {
	email := u.Email().Value()
	user := &schema.User{
		Name:        u.Name().Value(),
		DisplayName: u.DisplayName(),
		Email:       &email,
		Profile:     u.Profile(),
		Role:        enum.Role(u.Role().String()),
	}
	user.ID = u.Id().Value()
	return user
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
