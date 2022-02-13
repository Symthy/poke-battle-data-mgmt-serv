package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/items"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/repository/dto"
)

var _ repository.IHeldItemRepository = (*HeldItemRepository)(nil)

var (
	emptyHeldItemBuilder  = func() schema.HeldItem { return schema.HeldItem{} }
	emptyHeldItemsBuilder = func() []schema.HeldItem { return []schema.HeldItem{} }
)

type HeldItemRepository struct {
	dbClient orm.IDbClient
	BaseReadRepository[schema.HeldItem, items.HeldItem, items.HeldItems]
}

func NewHeldItemRepository(dbClient orm.IDbClient) *HeldItemRepository {
	return &HeldItemRepository{
		dbClient: dbClient,
		BaseReadRepository: BaseReadRepository[schema.HeldItem, items.HeldItem, items.HeldItems]{
			dbClient:            dbClient,
			emptySchemaBuilder:  emptyHeldItemBuilder,
			emptySchemasBuilder: emptyHeldItemsBuilder,
			domainsConstructor:  items.NewHeldItems,
			schemaConverter:     dto.ToSchemaHeldItem,
		},
	}
}

// FindById <- BaseReadRepository

// FindAll <- BaseReadRepository

func (r HeldItemRepository) FindByName(name string) (*items.HeldItems, error) {
	return r.FindByField("Name", name)
}
