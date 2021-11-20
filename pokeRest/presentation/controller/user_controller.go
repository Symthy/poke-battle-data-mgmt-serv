package controller

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/rest/autogen/server"
	"github.com/Symthy/PokeRest/pokeRest/application/service"
	"github.com/Symthy/PokeRest/pokeRest/presentation"
)

type UserController struct {
	service service.UserReadService
}

func NewUserController(service service.UserReadService) *UserController {
	return &UserController{service: service}
}

func (uc UserController) GetUser(id float32) (server.User, error) {
	user, err := uc.service.GetUser(uint(id))
	return presentation.ConvertUserToResponse(user), err
}
