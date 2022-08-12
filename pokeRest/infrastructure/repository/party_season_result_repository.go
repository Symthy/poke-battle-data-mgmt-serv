package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/parties"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/repository/conv"
)

var _ repository.IPartyBattleResultRepository = (*PartySeasonResultRepository)(nil)

var emptyPartySeasonResultSchemaBuilder = func() *schema.PartySeasonResult { return &schema.PartySeasonResult{} }

type PartySeasonResultRepository struct {
	BaseWriteRepository[schema.PartySeasonResult, parties.PartyBattleResult, identifier.PartyBattleResultId, uint64]
	dbClient orm.IDbClient
}

func NewPartyBattleResultRepository(dbClient orm.IDbClient) *PartySeasonResultRepository {
	return &PartySeasonResultRepository{
		BaseWriteRepository: BaseWriteRepository[schema.PartySeasonResult, parties.PartyBattleResult, identifier.PartyBattleResultId, uint64]{
			dbClient:           dbClient,
			emptySchemaBuilder: emptyPartySeasonResultSchemaBuilder,
			toSchemaConverter:  conv.ToSchemaPartySeasonResult,
			toDomainConverter:  conv.ToDomainPartyBattleResult,
		},
		dbClient: dbClient,
	}
}

// Create <- BaseWriteRepository
// Update <- BaseWriteRepository
// Delete <- BaseWriteRepository
