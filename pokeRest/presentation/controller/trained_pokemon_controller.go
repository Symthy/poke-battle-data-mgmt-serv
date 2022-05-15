package controller

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/rest/autogen/server"
	"github.com/Symthy/PokeRest/pokeRest/application/command"
	"github.com/Symthy/PokeRest/pokeRest/application/service/trainings"
	t_command "github.com/Symthy/PokeRest/pokeRest/application/service/trainings/command"
	d_trainings "github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
	"github.com/Symthy/PokeRest/pokeRest/domain/factory"
	"github.com/Symthy/PokeRest/pokeRest/presentation/auth"
	"github.com/Symthy/PokeRest/pokeRest/presentation/response"
	"github.com/labstack/echo/v4"
)

type TrainedPokemonController struct {
	writeServ        trainings.TrainedPokemonWriteService
	readServ         trainings.TrainedPokemonAdjustmentReadService
	userResolver     auth.AccessUserResolver
	responseResolver response.ResponseResolver[d_trainings.TrainedPokemon, server.TrainedPokemon]
}

func NewTrainedPokemonController(
	writeServ trainings.TrainedPokemonWriteService, userResolver auth.AccessUserResolver,
) *TrainedPokemonController {
	return &TrainedPokemonController{
		writeServ:        writeServ,
		userResolver:     userResolver,
		responseResolver: response.NewResponseResolver(response.ConvertTrainedPokemonToResponse),
	}
}

func (c TrainedPokemonController) FindTrainedPokemonAdjustments(ctx echo.Context, next uint64, pageSize uint64) {
	cmd := command.NewPaginationCommand(next, pageSize)
	c.readServ.FindAll(cmd)
}

func (c TrainedPokemonController) SaveTrainedPokemon(ctx echo.Context) error {
	userId, err := c.userResolver.ResolveId(ctx)
	if err != nil {
		return err
	}
	request := new(server.RequestCreateTrainedPokemon)
	if err = ctx.Bind(request); err != nil {
		return err
	}
	// Todo: accept value
	cmd := t_command.NewCreateTrainedPokemonCommand(request, userId)
	domain, error := c.writeServ.SaveTrainedPokemon(cmd)
	return c.responseResolver.Resolve(ctx, domain, error)
}

func (c TrainedPokemonController) UpdateTrainedPokemon(ctx echo.Context) error {
	userId, err := c.userResolver.ResolveId(ctx)
	if err != nil {
		return err
	}
	// Todo: accept value
	cmd := t_command.NewUpdateTrainedPokemonCommand(0, "", "", "", false, userId, factory.NewTrainedPokemonAdjustmentBuilder())
	domain, error := c.writeServ.UpdateTrainedPokemon(cmd)
	return c.responseResolver.Resolve(ctx, domain, error)
}

func (c TrainedPokemonController) DeleteTrainedPokemon(ctx echo.Context, id uint64) error {
	userId, err := c.userResolver.ResolveId(ctx)
	if err != nil {
		return err
	}
	cmd := t_command.NewDeleteTrainedPokemonCommand(id, userId)
	domain, error := c.writeServ.DeleteTrainedPokemon(cmd)
	return c.responseResolver.Resolve(ctx, domain, error)
}
