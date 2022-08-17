package repository

import (
	"time"

	"github.com/Symthy/PokeRest/internal/adapters/orm"
	"github.com/Symthy/PokeRest/internal/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/internal/domain/entity/battles"
	"github.com/Symthy/PokeRest/internal/domain/repository"
	"github.com/Symthy/PokeRest/internal/infrastructure/repository/conv"
)

var _ repository.IBattleSeasonRepository = (*BattleSeasonRepository)(nil)

type BattleSeasonRepository struct {
	dbClient orm.IDbClient
}

func NewBattleSeasonRepository(dbClient orm.IDbClient) *BattleSeasonRepository {
	return &BattleSeasonRepository{
		dbClient: dbClient,
	}
}

func (rep BattleSeasonRepository) FindCurrent() (*battles.Season, error) {
	db := rep.dbClient.Db()
	schema := schema.BattleSeason{}
	tx := db.Where("? BETWEEN start_date_time AND end_date_time", time.Now()).First(&schema)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return conv.ToDomainSeason(schema)
}

func (rep BattleSeasonRepository) Find(season battles.Season) (*battles.SeasonPeriods, error) {
	return nil, nil
}

func (rep BattleSeasonRepository) FindAll() (*battles.SeasonPeriods, error) {
	db := rep.dbClient.Db()
	schemas := []schema.BattleSeason{}
	tx := db.Find(schemas)

	if tx.Error != nil {
		return nil, tx.Error
	}
	seasonPeriods, err := convertToDomains(schemas)
	if err != nil {
		return nil, err
	}
	domains := battles.NewSeasonPeriods(seasonPeriods)
	return &domains, nil
}

// non use common func. season is non id...
func convertToDomains(schemas []schema.BattleSeason) ([]battles.SeasonPeriod, error) {
	domains := make([]battles.SeasonPeriod, len(schemas))
	for i, s := range schemas {
		d, err := conv.ToDomainSeasonPeriod(s)
		if err != nil {
			return nil, err
		}
		domains[i] = *d
	}
	return domains, nil
}
