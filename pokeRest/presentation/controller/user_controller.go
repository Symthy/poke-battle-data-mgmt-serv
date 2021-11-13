package controller

import "github.com/Symthy/PokeRest/pokeRest/application/service"

type UserController struct {
	service service.UserReadService
}
