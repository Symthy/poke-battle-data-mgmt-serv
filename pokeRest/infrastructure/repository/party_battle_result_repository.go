package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/parties"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/repository/dto"
)

var _ repository.IPartyBattleResultRepository = (*PartyBattleResultRepository)(nil)

var emptyPartySeasonResultSchemaBuilder = func() schema.PartyBattleResult { return schema.PartyBattleResult{} }

type PartyBattleResultRepository struct {
	BaseWriteRepository[schema.PartyBattleResult, parties.PartyBattleResult]
	dbClient orm.IDbClient
}

func NewPartyBattleResultRepository(dbClient orm.IDbClient) *PartyBattleResultRepository {
	return &PartyBattleResultRepository{
		BaseWriteRepository: BaseWriteRepository[schema.PartyBattleResult, parties.PartyBattleResult]{
			dbClient:           dbClient,
			emptySchemaBuilder: emptyPartySeasonResultSchemaBuilder,
			schemaConverter:    dto.ToSchemaPartyBattleResult,
		},
		dbClient: dbClient,
	}
}
