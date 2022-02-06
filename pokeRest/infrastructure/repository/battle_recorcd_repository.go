package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

var _ repository.IBattleRecordRepository = (*BattleRecordRepository)(nil)

type BattleRecordRepository struct {
	dbClient orm.IDbClient
}

func (repo BattleRecordRepository) Save() {
}

func (repo BattleRecordRepository) Update() {
}

func (repo BattleRecordRepository) Delete() {
}
