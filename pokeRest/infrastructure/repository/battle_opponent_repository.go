package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

var _ repository.IBattleOpponentPartyRepository = (*BattleOpponentPartyRepository)(nil)

type BattleOpponentPartyRepository struct {
	dbClient orm.IDbClient
}

func (repo BattleOpponentPartyRepository) Create(model battles.BattleOpponentParty) (*battles.BattleOpponentParty, error) {
	return nil, nil
}

func (repo BattleOpponentPartyRepository) Update(model battles.BattleOpponentParty) (*battles.BattleOpponentParty, error) {
	return nil, nil
}

func (repo BattleOpponentPartyRepository) Delete(model battles.BattleOpponentParty) (*battles.BattleOpponentParty, error) {
	return nil, nil
}
