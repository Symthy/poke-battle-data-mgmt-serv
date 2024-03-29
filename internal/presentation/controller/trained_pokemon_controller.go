package controller

import (
	"github.com/Symthy/PokeRest/internal/adapters/rest/autogen/server"
	"github.com/Symthy/PokeRest/internal/application/service/trainings"
	t_command "github.com/Symthy/PokeRest/internal/application/service/trainings/command"
	d_trainings "github.com/Symthy/PokeRest/internal/domain/entity/trainings"
	"github.com/Symthy/PokeRest/internal/domain/factory"
	"github.com/Symthy/PokeRest/internal/presentation/auth"
	"github.com/Symthy/PokeRest/internal/presentation/response"
	"github.com/labstack/echo/v4"
)

type TrainedPokemonController struct {
	writeServ                  trainings.TrainedPokemonWriteService
	readServ                   trainings.TrainedPokemonAdjustmentReadService
	userResolver               auth.AccessUserResolver
	responseSingleDataResolver response.ResponseResolver[d_trainings.TrainedPokemon, server.TrainedPokemon]
}

func NewTrainedPokemonController(
	writeServ trainings.TrainedPokemonWriteService, userResolver auth.AccessUserResolver,
) *TrainedPokemonController {
	return &TrainedPokemonController{
		writeServ:                  writeServ,
		userResolver:               userResolver,
		responseSingleDataResolver: response.NewResponseResolver(response.ConvertTrainedPokemonToResponse),
	}
}

func (c TrainedPokemonController) FindTrainedPokemonAdjustments(ctx echo.Context, next uint64, pageSize uint64) {
	// cmd := command.NewPaginationCommand(next, pageSize)
	// c.readServ.FindAll(cmd)
	// Todo: implements response resolver
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
	return c.responseSingleDataResolver.Resolve(ctx, domain, error)
}

func (c TrainedPokemonController) UpdateTrainedPokemon(ctx echo.Context) error {
	userId, err := c.userResolver.ResolveId(ctx)
	if err != nil {
		return err
	}
	// Todo: accept value
	cmd := t_command.NewUpdateTrainedPokemonCommand(0, "", "", "", false, userId, factory.NewTrainedPokemonAdjustmentBuilder())
	domain, error := c.writeServ.UpdateTrainedPokemon(cmd)
	return c.responseSingleDataResolver.Resolve(ctx, domain, error)
}

func (c TrainedPokemonController) DeleteTrainedPokemon(ctx echo.Context, id uint64) error {
	userId, err := c.userResolver.ResolveId(ctx)
	if err != nil {
		return err
	}
	cmd := t_command.NewDeleteTrainedPokemonCommand(id, userId)
	domain, error := c.writeServ.DeleteTrainedPokemon(cmd)
	return c.responseSingleDataResolver.Resolve(ctx, domain, error)
}
