package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

var _ repository.IPartyRepository = (*PartyRepository)(nil)

type PartyRepository struct {
	dbClient orm.IDbClient
}

func (repo PartyRepository) Save() {
}

func (repo PartyRepository) Update() {
}

func (repo PartyRepository) Delete() {
}
