package controller

import (
	"github.com/Symthy/PokeRest/internal/adapters/rest/autogen/server"
	"github.com/Symthy/PokeRest/internal/application/command"
	"github.com/Symthy/PokeRest/internal/application/service/moves"
	d_moves "github.com/Symthy/PokeRest/internal/domain/entity/moves"
	"github.com/Symthy/PokeRest/internal/presentation/response"
	"github.com/labstack/echo/v4"
)

type MoveController struct {
	service            moves.MoveReadService
	singleDataResolver response.ResponseResolver[d_moves.Move, server.Move]
	multiDataResolver  response.ResponseResolver[d_moves.Moves, server.Moves]
}

func NewMoveController(service moves.MoveReadService) *MoveController {
	return &MoveController{
		service:            service,
		singleDataResolver: response.NewResponseResolver(response.ConvertMoveToResponse),
		multiDataResolver:  response.NewResponseResolver(response.ConvertMovesToResponse),
	}
}

func (c MoveController) GetMoveById(ctx echo.Context, id uint64) error {
	domain, err := c.service.FindMove(id)
	return c.singleDataResolver.Resolve(ctx, domain, err)
}

func (c MoveController) GetMoves(ctx echo.Context, next uint64, pageSize uint64) error {
	cmd := command.NewPaginationCommand(next, pageSize)
	domains, err := c.service.FindAll(cmd)
	return c.multiDataResolver.Resolve(ctx, domains, err)
}

func (c MoveController) GetMoveByPokemonId(ctx echo.Context, pokemonId uint64) error {
	domains, err := c.service.FindOfPokemon(pokemonId)
	return c.multiDataResolver.Resolve(ctx, domains, err)
}
