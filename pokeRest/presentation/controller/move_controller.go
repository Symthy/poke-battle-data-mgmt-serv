package controller

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/rest/autogen/server"
	"github.com/Symthy/PokeRest/pokeRest/application/command"
	"github.com/Symthy/PokeRest/pokeRest/application/service/moves"
	d_moves "github.com/Symthy/PokeRest/pokeRest/domain/entity/moves"
	"github.com/Symthy/PokeRest/pokeRest/presentation/response"
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

func (c MoveController) GetMoves(ctx echo.Context, next int, pageSize int) error {
	cmd := command.NewPaginationCommand(next, pageSize)
	domains, err := c.service.FindAll(cmd)
	return c.multiDataResolver.Resolve(ctx, domains, err)
}

func (c MoveController) GetMoveById(ctx echo.Context, id int) error {
	domain, err := c.service.FindMove(uint(id))
	return c.singleDataResolver.Resolve(ctx, domain, err)
}

func (c MoveController) GetMoveByPokemonId(ctx echo.Context, pokemonId int) error {
	domains, err := c.service.FindOfPokemon(uint(pokemonId))
	return c.multiDataResolver.Resolve(ctx, domains, err)
}
