package controller

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/rest/autogen/server"
	"github.com/Symthy/PokeRest/pokeRest/application/command"
	"github.com/Symthy/PokeRest/pokeRest/application/service/trainings"
	t_command "github.com/Symthy/PokeRest/pokeRest/application/service/trainings/command"
	d_trainings "github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
	"github.com/Symthy/PokeRest/pokeRest/presentation/response"
	"github.com/labstack/echo/v4"
)

type TrainedPokemonController struct {
	readServ         trainings.TrainedPokemonAdjustmentReadService
	writeServ        trainings.TrainedPokemonWriteService
	responseResolver response.ResponseResolver[d_trainings.TrainedPokemon, server.TrainedPokemon]
}

func NewTrainedPokemonController(
	writeServ trainings.TrainedPokemonWriteService) *TrainedPokemonController {
	return &TrainedPokemonController{
		writeServ:        writeServ,
		responseResolver: response.NewResponseResolver(response.ConvertTrainedPokemonToResponse),
	}
}

func (c TrainedPokemonController) FindTrainedPokemonAdjustments(ctx echo.Context, next int, pageSize int) {
	cmd := command.NewPaginationCommand(next, pageSize)
	c.readServ.FindAll(cmd)
}

func (c TrainedPokemonController) SaveTrainedPokemon(ctx echo.Context) error {
	// Todo: accept value
	cmd := t_command.NewCreateTrainedPokemonCommand()
	domain, error := c.writeServ.SaveTrainedPokemon(cmd)
	return c.responseResolver.Resolve(ctx, domain, error)
}

func (c TrainedPokemonController) UpdateTrainedPokemon(ctx echo.Context) error {
	// Todo: accept value
	cmd := t_command.NewUpdateTrainedPokemonCommand()
	domain, error := c.writeServ.UpdateTrainedPokemon(cmd)
	return c.responseResolver.Resolve(ctx, domain, error)
}

func (c TrainedPokemonController) DeleteTrainedPokemon(ctx echo.Context, id uint) error {
	domain, error := c.writeServ.DeleteTrainedPokemon(id)
	return c.responseResolver.Resolve(ctx, domain, error)
}
