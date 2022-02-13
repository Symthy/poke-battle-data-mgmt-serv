package controller

import (
	"strconv"

	"github.com/Symthy/PokeRest/pokeRest/adapters/rest/autogen/server"
	"github.com/Symthy/PokeRest/pokeRest/application/service/users"
	"github.com/Symthy/PokeRest/pokeRest/application/service/users/command"
	d_users "github.com/Symthy/PokeRest/pokeRest/domain/entity/users"
	"github.com/Symthy/PokeRest/pokeRest/presentation/response"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	service          users.UserReadService
	responseResolver response.ResponseResolver[d_users.User, server.User]
}

func NewUserController(service users.UserReadService) *UserController {
	return &UserController{
		service:          service,
		responseResolver: response.NewResponseResolver(response.ConvertUserToResponse),
	}
}

func (uc UserController) GetUserById(ctx echo.Context, id string) error {
	userId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		// Todo: error wrap
		return err
	}
	// Todo: change uint64
	user, err := uc.service.GetUserById(uint(userId))
	return uc.responseResolver.Resolve(ctx, user, err)
}

func (uc UserController) GetUser(ctx echo.Context, name string) error {
	cmd := command.NewGetUserCommand(name).SetFilterRequiredFields()
	user, err := uc.service.GetUser(*cmd)
	return uc.responseResolver.Resolve(ctx, user, err)
}
