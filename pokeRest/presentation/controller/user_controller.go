package controller

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/rest/autogen/server"
	"github.com/Symthy/PokeRest/pokeRest/application/service/users"
	"github.com/Symthy/PokeRest/pokeRest/application/service/users/command"
	"github.com/Symthy/PokeRest/pokeRest/presentation"
)

type UserController struct {
	service users.UserReadService
}

func NewUserController(service users.UserReadService) *UserController {
	return &UserController{service: service}
}

func (uc UserController) GetUserById(id float32) (server.User, error) {
	user, err := uc.service.GetUserById(uint(id))
	return presentation.ConvertUserToResponse(user), err
}

func (uc UserController) GetUser(name string) (server.User, error) {
	command := command.NewGetUserCommand(name).SetFilterRequiredFields()
	user, err := uc.service.GetUser(*command)
	return presentation.ConvertUserToResponse(user), err
}
