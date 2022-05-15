package controller

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/rest/autogen/server"
	"github.com/Symthy/PokeRest/pokeRest/application/service/parties"
	d_parties "github.com/Symthy/PokeRest/pokeRest/domain/entity/parties"
	"github.com/Symthy/PokeRest/pokeRest/presentation/response"
	"github.com/labstack/echo/v4"
)

type PartyTagController struct {
	readServ         parties.PartyTagReadService
	writeServ        parties.PartyTagWriteService
	responseResolver response.ResponseResolver[d_parties.PartyTag, server.PartyTag]
}

func NewPartyTagController(readServ parties.PartyTagReadService, writeServ parties.PartyTagWriteService) *PartyTagController {
	return &PartyTagController{
		readServ:         readServ,
		writeServ:        writeServ,
		responseResolver: response.NewResponseResolver(response.ConvertPartyTagToResponse),
	}
}

func (c PartyTagController) GetPartyTag(ctx echo.Context, id uint64) error {
	domain, err := c.readServ.FindPartyTag(id)
	return c.responseResolver.Resolve(ctx, domain, err)
}

func (c PartyTagController) SavePartyTag(ctx echo.Context, name string) error {
	domain, err := c.writeServ.SavePartyTag(name)
	return c.responseResolver.Resolve(ctx, domain, err)
}

func (c PartyTagController) DeletePartyTag(ctx echo.Context, id uint64) error {
	domain, err := c.writeServ.DeletePartyTag(id)
	return c.responseResolver.Resolve(ctx, domain, err)
}
