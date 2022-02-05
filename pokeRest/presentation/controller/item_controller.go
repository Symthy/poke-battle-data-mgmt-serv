package controller

import "github.com/Symthy/PokeRest/pokeRest/application/service/items"

type ItemController struct {
	service items.ItemReadService
}
