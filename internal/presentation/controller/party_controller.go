package controller

import (
	"github.com/Symthy/PokeRest/internal/adapters/rest/autogen/server"
	"github.com/Symthy/PokeRest/internal/application/service/parties"
	"github.com/Symthy/PokeRest/internal/application/service/parties/command"
	d_parties "github.com/Symthy/PokeRest/internal/domain/entity/parties"
	"github.com/Symthy/PokeRest/internal/presentation/auth"
	"github.com/Symthy/PokeRest/internal/presentation/response"
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
	cmd := command.NewCreatePartyCommand("", "", false, userId, []uint64{}, []uint64{})
	domain, err := c.writeServ.SaveParty(cmd)
	return c.responseResolver.Resolve(ctx, domain, err)
}

func (c PartyController) UpdateParty(ctx echo.Context, id uint64) error {
	userId, err := c.userResolver.ResolveId(ctx)
	if err != nil {
		return err
	}
	// Todo: accept value
	cmd := command.NewUpdatePartyCommand(id, "", "", false, userId, []uint64{}, []uint64{}, []uint64{})
	domain, err := c.writeServ.UpdateParty(cmd)
	return c.responseResolver.Resolve(ctx, domain, err)
}

func (c PartyController) DeleteParty(ctx echo.Context, id uint64) error {
	userId, err := c.userResolver.ResolveId(ctx)
	if err != nil {
		return err
	}
	cmd := command.NewDeletePartyCommand(id, userId)
	domain, err := c.writeServ.DeleteParty(cmd)
	return c.responseResolver.Resolve(ctx, domain, err)
}
