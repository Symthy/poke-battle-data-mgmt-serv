package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

var _ repository.IBattleRecordRepository = (*BattleRecordRepository)(nil)

type BattleRecordRepository struct {
	dbClient orm.IDbClient
}

func (repo BattleRecordRepository) Create(model battles.BattleRecord) (*battles.BattleRecord, error) {
	return nil, nil
}

func (repo BattleRecordRepository) Update(model battles.BattleRecord) (*battles.BattleRecord, error) {
	return nil, nil
}

func (repo BattleRecordRepository) Delete(model battles.BattleRecord) (*battles.BattleRecord, error) {
	return nil, nil
}
