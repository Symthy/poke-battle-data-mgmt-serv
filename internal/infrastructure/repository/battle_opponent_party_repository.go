package repository

import (
	"github.com/Symthy/PokeRest/internal/adapters/orm"
	"github.com/Symthy/PokeRest/internal/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/internal/domain/entity/battles"
	"github.com/Symthy/PokeRest/internal/domain/repository"
	"github.com/Symthy/PokeRest/internal/domain/value/identifier"
	"github.com/Symthy/PokeRest/internal/infrastructure/repository/conv"
	"gorm.io/gorm"
)

var _ repository.IBattleOpponentPartyRepository = (*BattleOpponentPartyRepository)(nil)

var emptyBattleOpponentPartySchemaBuilder = func() *schema.BattleOpponentParty {
	return &schema.BattleOpponentParty{}
}

type BattleOpponentPartyRepository struct {
	BaseWriteRepository[schema.BattleOpponentParty, battles.BattleOpponentParty, identifier.BattleOpponentPartyId, uint64]
	dbClient orm.IDbClient
}

func NewBattleOpponentPartyRepository(dbClient orm.IDbClient) *BattleOpponentPartyRepository {
	return &BattleOpponentPartyRepository{
		BaseWriteRepository: BaseWriteRepository[schema.BattleOpponentParty, battles.BattleOpponentParty, identifier.BattleOpponentPartyId, uint64]{
			dbClient:           dbClient,
			emptySchemaBuilder: emptyBattleOpponentPartySchemaBuilder,
			toSchemaConverter:  conv.ToSchemaBattleOpponentParty,
			toDomainConverter:  conv.ToDomainBattleOpponentParty,
		},
		dbClient: dbClient,
	}
}

// Create <- BaseWriteRepository
// Update <- BaseWriteRepository
// Delete <- BaseWriteRepository

func (rep BattleOpponentPartyRepository) FindParty(opponentParty *battles.BattleOpponentParty) (*battles.BattleOpponentParty, error) {
	db := rep.dbClient.Db()

	searchParty := conv.ToSchemaBattleOpponentParty(opponentParty)
	ret := &schema.BattleOpponentParty{}
	tx := db.Where(searchParty).First(ret)
	if tx.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if tx != nil {
		return nil, tx.Error
	}
	return rep.toDomainConverter(ret)
}
