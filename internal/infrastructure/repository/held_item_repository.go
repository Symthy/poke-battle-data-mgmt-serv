package repository

import (
	"github.com/Symthy/PokeRest/internal/adapters/orm"
	"github.com/Symthy/PokeRest/internal/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/internal/domain/entity/items"
	"github.com/Symthy/PokeRest/internal/domain/repository"
	"github.com/Symthy/PokeRest/internal/domain/value/identifier"
	"github.com/Symthy/PokeRest/internal/infrastructure/repository/conv"
)

var _ repository.IHeldItemRepository = (*HeldItemRepository)(nil)

var (
	emptyHeldItemBuilder  = func() *schema.HeldItem { return &schema.HeldItem{} }
	emptyHeldItemsBuilder = func() []*schema.HeldItem { return []*schema.HeldItem{} }
)

type HeldItemRepository struct {
	dbClient orm.IDbClient
	BaseReadRepository[schema.HeldItem, items.HeldItem, items.HeldItems, identifier.HeldItemId, uint16]
}

func NewHeldItemRepository(dbClient orm.IDbClient) *HeldItemRepository {
	return &HeldItemRepository{
		dbClient: dbClient,
		BaseReadRepository: NewBaseReadRepository[schema.HeldItem, items.HeldItem, items.HeldItems, identifier.HeldItemId, uint16](
			dbClient,
			emptyHeldItemBuilder,
			emptyHeldItemsBuilder,
			items.NewHeldItems,
			conv.ToSchemaHeldItem,
			conv.ToDomainHeldItem,
		),
	}
}

// FindById <- BaseReadRepository

// FindAll <- BaseReadRepository

func (r HeldItemRepository) FindByName(name string) (*items.HeldItems, error) {
	return r.FindByField("Name", name)
}
