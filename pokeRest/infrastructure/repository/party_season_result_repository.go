package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/parties"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

var _ repository.IPartySeasonResultRepository = (*PartySeasonResultRepository)(nil)

type PartySeasonResultRepository struct {
	dbClient orm.IDbClient
}

func (s PartySeasonResultRepository) Create(model parties.PartySeasonResult) (*parties.PartySeasonResult, error) {
	return nil, nil
}

func (s PartySeasonResultRepository) Update(model parties.PartySeasonResult) (*parties.PartySeasonResult, error) {
	return nil, nil
}

func (s PartySeasonResultRepository) Delete(model parties.PartySeasonResult) (*parties.PartySeasonResult, error) {
	return nil, nil
}
