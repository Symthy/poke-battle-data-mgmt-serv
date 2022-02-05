package controller

import (
	"github.com/Symthy/PokeRest/pokeRest/application/service/abilities"
)

type AbilityController struct {
	service abilities.AbilityReadService
}
