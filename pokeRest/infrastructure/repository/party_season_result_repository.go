package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

var _ repository.IPartySeasonResultRepository = (*PartySeasonResultRepository)(nil)

type PartySeasonResultRepository struct {
	dbClient orm.IDbClient
}

func (s PartySeasonResultRepository) Save() {
}

func (s PartySeasonResultRepository) Update() {
}

func (s PartySeasonResultRepository) Delete() {
}
