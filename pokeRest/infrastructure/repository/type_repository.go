package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/types"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

var _ repository.ITypeRepository = (*TypeRepository)(nil)

type TypeRepository struct {
	dbClient orm.IDbClient
}

func NewTypeRepository(dbClient orm.IDbClient) *TypeRepository {
	return &TypeRepository{dbClient: dbClient}
}

func (repo TypeRepository) BuildTypeCompatibility() (*types.TypeCompatibility, error) {
	return nil, nil
}
