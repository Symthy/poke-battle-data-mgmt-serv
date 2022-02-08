package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure"
)

var _ repository.IBattleOpponentPartyRepository = (*BattleOpponentPartyRepository)(nil)

var emptyBattleOpponentPartySchemaBuilder = func() schema.BattleOpponentParty {
	return schema.BattleOpponentParty{}
}

type BattleOpponentPartyRepository struct {
	BaseWriteRepository[schema.BattleOpponentParty, battles.BattleOpponentParty]
	dbClient orm.IDbClient
}

func NewBattleOpponentParty(dbClient orm.IDbClient) *BattleOpponentPartyRepository {
	return &BattleOpponentPartyRepository{
		BaseWriteRepository: BaseWriteRepository[schema.BattleOpponentParty, battles.BattleOpponentParty]{
			dbClient:           dbClient,
			emptySchemaBuilder: emptyBattleOpponentPartySchemaBuilder,
			schemaConverter:    infrastructure.ToSchemaBattleOpponentParty,
		},
		dbClient: dbClient,
	}
}
