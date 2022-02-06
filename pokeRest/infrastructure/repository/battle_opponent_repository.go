package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

var _ repository.IBattleOpponentPartyRepository = (*BattleOpponentPartyRepository)(nil)

type BattleOpponentPartyRepository struct {
	dbClient orm.IDbClient
}

func (repo BattleOpponentPartyRepository) Save() {
}

func (repo BattleOpponentPartyRepository) Update() {
}

func (repo BattleOpponentPartyRepository) Delete() {
}
