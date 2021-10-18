package handler

import (
	"sync"

	"github.com/Symthy/PokeRest/pokeRest/autoGen/server"
	"github.com/labstack/echo/v4"
)

type PokemonHandler struct {
	Lock sync.Mutex
}

func NewPokemonHandler() *PokemonHandler {
	return &PokemonHandler{}
}

func (h *PokemonHandler) GetPokemonsId(ctx echo.Context, pokedexId string) error {
	h.Lock.Lock()
	defer h.Lock.Unlock()

	return nil
}

func (h *PokemonHandler) GetPokemons(ctx echo.Context, params server.GetPokemonsParams) error {
	h.Lock.Lock()
	defer h.Lock.Unlock()

	return nil
}
