package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure"
)

var _ repository.IBattleRecordRepository = (*BattleRecordRepository)(nil)

var emptyBattleRecordSchemaBuilder = func() schema.BattleRecord { return schema.BattleRecord{} }

type BattleRecordRepository struct {
	BaseWriteRepository[schema.BattleRecord, battles.BattleRecord]
	dbClient orm.IDbClient
}

func NewBattleRecordRepository(dbClient orm.IDbClient) *BattleRecordRepository {
	return &BattleRecordRepository{
		BaseWriteRepository: BaseWriteRepository[schema.BattleRecord, battles.BattleRecord]{
			dbClient:           dbClient,
			emptySchemaBuilder: emptyBattleRecordSchemaBuilder,
			schemaConverter:    infrastructure.ToSchemaBattleRecord,
		},
		dbClient: dbClient,
	}
}
