package controller

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/rest/autogen/server"
	"github.com/Symthy/PokeRest/pokeRest/application/service/battles"
	"github.com/Symthy/PokeRest/pokeRest/application/service/battles/command"
	d_battles "github.com/Symthy/PokeRest/pokeRest/domain/entity/battles"
	"github.com/Symthy/PokeRest/pokeRest/presentation/auth"
	"github.com/Symthy/PokeRest/pokeRest/presentation/response"
	"github.com/labstack/echo/v4"
)

type BattleRecordController struct {
	service            battles.BattleRecordWriteService
	userResolver       auth.AccessUserResolver
	singleDataResolver response.ResponseResolver[d_battles.BattleRecord, server.Ability]
}

func NewBattleRecordController(
	service battles.BattleRecordWriteService, userResolver auth.AccessUserResolver) *BattleRecordController {
	return &BattleRecordController{
		service:      service,
		userResolver: userResolver,
	}
}

func (c BattleRecordController) SaveBattleRecord(ctx echo.Context) error {
	userId, err := c.userResolver.ResolveId(ctx)
	if err != nil {
		return err
	}
	// Todo: accept value
	cmd := command.NewAddBattleRecordCommand(0, userId, 0, 0, 0, "", []uint64{}, []uint64{}, []uint64{}, []uint64{})
	domain, err := c.service.AddBattleRecord(cmd)
	return c.singleDataResolver.Resolve(ctx, domain, err)
}

func (c BattleRecordController) UpdateBattleRecord(ctx echo.Context) error {
	userId, err := c.userResolver.ResolveId(ctx)
	if err != nil {
		return err
	}
	// Todo: accept value
	cmd := command.NewEditBattleRecordCommand(0, 0, userId, 0, 0, 0, "", []uint64{}, []uint64{}, []uint64{}, []uint64{})
	domain, err := c.service.EditBattleRecord(cmd)
	return c.singleDataResolver.Resolve(ctx, domain, err)
}

func (c BattleRecordController) DeleteBattleRecord(ctx echo.Context, id uint64) error {
	userId, err := c.userResolver.ResolveId(ctx)
	if err != nil {
		return err
	}
	cmd := command.NewDeleteBattleRecord(id, userId)
	_, err = c.service.DeleteBattleRecord(cmd)
	// Todo: delete success response
	return err
}
