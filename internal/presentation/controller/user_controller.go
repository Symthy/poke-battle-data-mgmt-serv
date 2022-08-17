package controller

import (
	"github.com/Symthy/PokeRest/internal/adapters/rest/autogen/server"
	"github.com/Symthy/PokeRest/internal/application/service/users"
	"github.com/Symthy/PokeRest/internal/application/service/users/command"
	d_users "github.com/Symthy/PokeRest/internal/domain/entity/users"
	"github.com/Symthy/PokeRest/internal/presentation/response"
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

func (uc UserController) GetUserById(ctx echo.Context, id uint64) error {
	// Todo: change uint64
	user, err := uc.service.GetUserById(id)
	return uc.responseResolver.Resolve(ctx, user, err)
}

func (uc UserController) GetUser(ctx echo.Context, name string) error {
	cmd := command.NewGetUserCommand(name).SetFilterRequiredFields()
	user, err := uc.service.GetUser(*cmd)
	return uc.responseResolver.Resolve(ctx, user, err)
}
