package controller

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/rest/autogen/server"
	"github.com/Symthy/PokeRest/pokeRest/application/service/parties"
	"github.com/Symthy/PokeRest/pokeRest/application/service/parties/command"
	d_parties "github.com/Symthy/PokeRest/pokeRest/domain/entity/parties"
	"github.com/Symthy/PokeRest/pokeRest/presentation/response"
	"github.com/labstack/echo/v4"
)

type PartyController struct {
	writeServ        parties.PartyWriteService
	responseResolver response.ResponseResolver[d_parties.Party, server.Party]
}

func NewPartyController(writeServ parties.PartyWriteService) *PartyController {
	return &PartyController{
		writeServ:        writeServ,
		responseResolver: response.NewResponseResolver(response.ConvertPartyToResponse),
	}
}

func (c PartyController) SaveParty(ctx echo.Context) error {
	// Todo: accept value
	cmd := command.NewCreatePartyCommand("", "", false, 0, []uint{}, []uint{})
	domain, err := c.writeServ.SaveParty(cmd)
	return c.responseResolver.Resolve(ctx, domain, err)
}

func (c PartyController) UpdateParty(ctx echo.Context, id int) error {
	// Todo: accept value
	cmd := command.NewUpdatePartyCommand(uint(id), "", "", false, []uint{}, []uint{}, []uint{})
	domain, err := c.writeServ.UpdateParty(cmd)
	return c.responseResolver.Resolve(ctx, domain, err)
}

func (c PartyController) DeleteParty(ctx echo.Context, id int) error {
	domain, err := c.writeServ.DeleteParty(uint(id))
	return c.responseResolver.Resolve(ctx, domain, err)
}
