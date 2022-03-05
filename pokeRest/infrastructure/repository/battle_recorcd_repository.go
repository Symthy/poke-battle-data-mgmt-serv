package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/repository/conv"
)

var _ repository.IBattleRecordRepository = (*BattleRecordRepository)(nil)

var emptyBattleRecordSchemaBuilder = func() schema.BattleRecord { return schema.BattleRecord{} }

type BattleRecordRepository struct {
	BaseSingleReadRepository[schema.BattleRecord, battles.BattleRecord, identifier.BattleRecordId]
	BaseWriteRepository[schema.BattleRecord, battles.BattleRecord, identifier.BattleRecordId]
	dbClient orm.IDbClient
}

func NewBattleRecordRepository(dbClient orm.IDbClient) *BattleRecordRepository {
	return &BattleRecordRepository{
		BaseSingleReadRepository: BaseSingleReadRepository[schema.BattleRecord, battles.BattleRecord, identifier.BattleRecordId]{
			dbClient:           dbClient,
			emptySchemaBuilder: emptyBattleRecordSchemaBuilder,
			toSchemaConverter:  conv.ToSchemaBattleRecord,
			toDomainConverter:  conv.ToDomainBattleRecord,
		},
		BaseWriteRepository: BaseWriteRepository[schema.BattleRecord, battles.BattleRecord, identifier.BattleRecordId]{
			dbClient:           dbClient,
			emptySchemaBuilder: emptyBattleRecordSchemaBuilder,
			toSchemaConverter:  conv.ToSchemaBattleRecord,
			toDomainConverter:  conv.ToDomainBattleRecord,
		},
		dbClient: dbClient,
	}
}

// FindById <- BaseSingleReadRepository

// Create <- BaseWriteRepository
// Update <- BaseWriteRepository
// Delete <- BaseWriteRepository
