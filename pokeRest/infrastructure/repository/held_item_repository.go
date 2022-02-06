package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/items"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

var _ repository.IItemRepository = (*HeldItemRepository)(nil)

type HeldItemRepository struct {
	dbClient orm.IDbClient
}

func (repo HeldItemRepository) FindAll(page int, pageSize int) (*items.HeldItems, error) {
	return nil, nil
}
