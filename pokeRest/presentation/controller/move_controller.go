package controller

import "github.com/Symthy/PokeRest/pokeRest/application/service/move"

type MoveController struct {
	service move.MoveReadService
}
