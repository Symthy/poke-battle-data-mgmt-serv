package controller

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/rest/autogen/server"
	"github.com/Symthy/PokeRest/pokeRest/application/service/battles"
	"github.com/Symthy/PokeRest/pokeRest/application/service/battles/command"
	d_battles "github.com/Symthy/PokeRest/pokeRest/domain/entity/battles"
	"github.com/Symthy/PokeRest/pokeRest/presentation/response"
	"github.com/labstack/echo/v4"
)

type BattleRecordController struct {
	service            battles.BattleRecordWriteService
	singleDataResolver response.ResponseResolver[d_battles.BattleRecord, server.Ability]
}

func NewBattleRecordController(service battles.BattleRecordWriteService) *BattleRecordController {
	return &BattleRecordController{service: service}
}

func (c BattleRecordController) SaveBattleRecord(ctx echo.Context) error {
	// Todo: accept value
	cmd := command.NewAddBattleRecordCommand(0, "", "", 0, 0, 0, []int{}, []uint{}, []int{}, []int{})
	domain, err := c.service.AddBattleRecord(cmd)
	return c.singleDataResolver.Resolve(ctx, domain, err)
}

func (c BattleRecordController) UpdateBattleRecord(ctx echo.Context) error {
	// Todo: accept value
	cmd := command.NewEditBattleRecordCommand(0, 0, "", "", 0, 0, 0, []int{}, []uint{}, []int{})
	domain, err := c.service.EditBattleRecord(cmd)
	return c.singleDataResolver.Resolve(ctx, domain, err)
}

func (c BattleRecordController) DeleteBattleRecord(ctx echo.Context, id uint) error {
	_, err := c.service.DeleteBattleRecord(id)
	// Todo: delete success response
	return err
}
