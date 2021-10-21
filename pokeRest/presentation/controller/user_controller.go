package controller

import "github.com/Symthy/PokeRest/pokeRest/application/service/user"

type UserController struct {
	service user.UserReadService
}
