package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/parties"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

var _ repository.IPartyRepository = (*PartyRepository)(nil)

type PartyRepository struct {
	dbClient orm.IDbClient
}

func (repo PartyRepository) Create(model parties.Party) (*parties.Party, error) {
	return nil, nil
}

func (repo PartyRepository) Update(model parties.Party) (*parties.Party, error) {
	return nil, nil
}

func (repo PartyRepository) Delete(model parties.Party) (*parties.Party, error) {
	return nil, nil
}
