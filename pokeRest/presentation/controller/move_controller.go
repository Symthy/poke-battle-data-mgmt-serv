package controller

import "github.com/Symthy/PokeRest/pokeRest/application/service/moves"

type MoveController struct {
	service moves.MoveReadService
}
