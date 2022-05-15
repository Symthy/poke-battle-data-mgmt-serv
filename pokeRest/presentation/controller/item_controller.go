package controller

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/rest/autogen/server"
	"github.com/Symthy/PokeRest/pokeRest/application/command"
	"github.com/Symthy/PokeRest/pokeRest/application/service/items"
	d_items "github.com/Symthy/PokeRest/pokeRest/domain/entity/items"
	"github.com/Symthy/PokeRest/pokeRest/presentation/response"
	"github.com/labstack/echo/v4"
)

type ItemController struct {
	service            items.ItemReadService
	singleDataResolver response.ResponseResolver[d_items.HeldItem, server.HeldItem]
	multiDataResolver  response.ResponseResolver[d_items.HeldItems, server.HeldItems]
}

func NewItemController(service items.ItemReadService) *ItemController {
	return &ItemController{
		service:            service,
		singleDataResolver: response.NewResponseResolver(response.ConvertHeldItemToResponse),
		multiDataResolver:  response.NewResponseResolver(response.ConvertHeldItemsToResponse),
	}
}

func (c ItemController) GetHeldItemById(ctx echo.Context, id uint64) error {
	domain, err := c.service.FindHeldItem(id)
	return c.singleDataResolver.Resolve(ctx, domain, err)
}
func (c ItemController) GetHeldItems(ctx echo.Context, next uint64, pageSize uint64) error {
	cmd := command.NewPaginationCommand(next, pageSize)
	domains, err := c.service.FindAll(cmd)
	return c.multiDataResolver.Resolve(ctx, domains, err)
}
