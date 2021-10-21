package controller

import "github.com/Symthy/PokeRest/pokeRest/application/service/item"

type ItemController struct {
	service item.ItemReadService
}
