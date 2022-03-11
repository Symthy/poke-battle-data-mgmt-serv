package controller

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/rest/autogen/server"
	"github.com/Symthy/PokeRest/pokeRest/application/service/parties"
	"github.com/Symthy/PokeRest/pokeRest/application/service/parties/command"
	d_parties "github.com/Symthy/PokeRest/pokeRest/domain/entity/parties"
	"github.com/Symthy/PokeRest/pokeRest/presentation/auth"
	"github.com/Symthy/PokeRest/pokeRest/presentation/response"
	"github.com/labstack/echo/v4"
)

type PartyController struct {
	writeServ        parties.PartyWriteService
	userResolver     auth.AccessUserResolver
	responseResolver response.ResponseResolver[d_parties.Party, server.Party]
}

func NewPartyController(
	writeServ parties.PartyWriteService, userResolver auth.AccessUserResolver) *PartyController {
	return &PartyController{
		writeServ:        writeServ,
		userResolver:     userResolver,
		responseResolver: response.NewResponseResolver(response.ConvertPartyToResponse),
	}
}

func (c PartyController) SaveParty(ctx echo.Context) error {
	userId, err := c.userResolver.ResolveId(ctx)
	if err != nil {
		return err
	}
	// Todo: accept value
	cmd := command.NewCreatePartyCommand("", "", false, userId, []uint{}, []uint{})
	domain, err := c.writeServ.SaveParty(cmd)
	return c.responseResolver.Resolve(ctx, domain, err)
}

func (c PartyController) UpdateParty(ctx echo.Context, id int) error {
	userId, err := c.userResolver.ResolveId(ctx)
	if err != nil {
		return err
	}
	// Todo: accept value
	cmd := command.NewUpdatePartyCommand(uint(id), "", "", false, userId, []uint{}, []uint{}, []uint{})
	domain, err := c.writeServ.UpdateParty(cmd)
	return c.responseResolver.Resolve(ctx, domain, err)
}

func (c PartyController) DeleteParty(ctx echo.Context, id uint) error {
	userId, err := c.userResolver.ResolveId(ctx)
	if err != nil {
		return err
	}
	cmd := command.NewDeletePartyCommand(id, userId)
	domain, err := c.writeServ.DeleteParty(cmd)
	return c.responseResolver.Resolve(ctx, domain, err)
}
