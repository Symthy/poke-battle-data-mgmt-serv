package repository

import (
	"time"

	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/repository/dto"
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
	domain := schema.ConvertToDomainSeason()
	return &domain, nil
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
	domains := battles.NewSeasonPeriods(dto.ConvertToDomains[schema.BattleSeason, battles.SeasonPeriod](schemas))
	return &domains, nil
}
