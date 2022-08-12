package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/repository/conv"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ repository.IBattleRecordRepository = (*BattleRecordRepository)(nil)

var emptyBattleRecordSchemaBuilder = func() *schema.BattleRecord { return &schema.BattleRecord{} }

type BattleRecordRepository struct {
	BaseSingleReadRepository[schema.BattleRecord, battles.BattleRecord, identifier.BattleRecordId, uint64]
	seasonRepo        BattleSeasonRepository
	opponentPartyRepo BattleOpponentPartyRepository
	dbClient          orm.IDbClient
}

func NewBattleRecordRepository(dbClient orm.IDbClient) *BattleRecordRepository {
	return &BattleRecordRepository{
		BaseSingleReadRepository: BaseSingleReadRepository[schema.BattleRecord, battles.BattleRecord, identifier.BattleRecordId, uint64]{
			dbClient:           dbClient,
			emptySchemaBuilder: emptyBattleRecordSchemaBuilder,
			toSchemaConverter:  conv.ToSchemaBattleRecord,
			toDomainConverter:  conv.ToDomainBattleRecord,
		},
		dbClient: dbClient,
	}
}

// FindById <- BaseSingleReadRepository

func (rep BattleRecordRepository) Create(domain *battles.BattleRecord) (*battles.BattleRecord, error) {
	db := rep.dbClient.Db()
	if domain.UnsolvedSeason() {
		currentSeason, err := rep.seasonRepo.FindCurrent()
		if err != nil {
			return nil, err
		}
		domain.ApplySeason(*currentSeason)
	}

	battleRecordSchema := conv.ToSchemaBattleRecord(domain)
	db.Transaction(func(tx *gorm.DB) error {
		opponentParty, err := rep.opponentPartyRepo.FindParty(domain.OpponentParty())
		if err != nil {
			return err
		}

		if opponentParty == nil {
			opponentPartySchema := conv.ToSchemaBattleOpponentParty(domain.OpponentParty())
			battleRecordSchema.BattleOpponentParty = *opponentPartySchema
			tx.Omit("Season").Create(&battleRecordSchema)
		} else {
			battleRecordSchema.ID = opponentParty.Id().Value()
			tx.Omit(clause.Associations).Create(&battleRecordSchema)
		}
		return nil
	})
	return conv.ToDomainBattleRecord(battleRecordSchema)
}

func (rep BattleRecordRepository) Update(domain *battles.BattleRecord) (*battles.BattleRecord, error) {
	db := rep.dbClient.Db()
	if domain.UnsolvedSeason() {
		currentSeason, err := rep.seasonRepo.FindCurrent()
		if err != nil {
			return nil, err
		}
		domain.ApplySeason(*currentSeason)
	}

	ret := emptyBattleRecordSchemaBuilder()
	battleRecordSchema := conv.ToSchemaBattleRecord(domain)
	db.Transaction(func(tx *gorm.DB) error {
		opponentParty, err := rep.opponentPartyRepo.FindParty(domain.OpponentParty())
		if err != nil {
			return err
		}

		if opponentParty == nil {
			createdOpponentParty, err := rep.opponentPartyRepo.Create(domain.OpponentParty())
			if err != nil {
				return err
			}
			opponentParty = createdOpponentParty
		}
		battleRecordSchema.ID = opponentParty.Id().Value()
		tx.Model(&ret).Omit(clause.Associations).Updates(&battleRecordSchema)
		return nil
	})
	return conv.ToDomainBattleRecord(ret)
}

func (rep BattleRecordRepository) Delete(id uint64) (*battles.BattleRecord, error) {
	db := rep.dbClient.Db()
	deleted := emptyBattleRecordSchemaBuilder()
	db.Transaction(func(tx *gorm.DB) error {
		if ret := tx.First(schema.BattleRecord{}, id); tx.Error != nil {
			return ret.Error
		}
		ret := tx.Delete(&deleted, id)
		if ret.Error != nil {
			return ret.Error
		}
		return nil
	})
	// Todo:
	// delete used opponent party if ref num is zero
	isRefZero := false
	if isRefZero {
		_, err := rep.opponentPartyRepo.Delete(deleted.BattleOpponentPartyId)
		if err != nil {
			return nil, err
		}
	}
	return conv.ToDomainBattleRecord(deleted)
}
