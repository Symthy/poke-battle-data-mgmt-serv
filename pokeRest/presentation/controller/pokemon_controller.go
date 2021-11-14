package controller

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/rest/autogen/server"
	"github.com/Symthy/PokeRest/pokeRest/application/service"
	"github.com/Symthy/PokeRest/pokeRest/domain/model"
)

type PokemonController struct {
	service service.PokemonReadService
}

func NewPokemonController(service service.PokemonReadService) *PokemonController {
	return &PokemonController{service: service}
}

func (c PokemonController) GetPokemon(id float32) server.Pokemon {
	ret := server.Pokemon{}
	if pokemon, err := c.service.GetPokemon(uint(id)); err == nil {
		ret = ConvertToResponse(pokemon)
	}
	return ret
}

// visible for testing
func ConvertToResponse(domain model.Pokemon) server.Pokemon {
	ability1, _ := domain.AbilityIdPrimary().Get()
	var ability2 *float32 = nil
	if a2, _ := domain.AbilityIdSecondary().Get(); a2 != nil {
		*ability2 = float32(*a2)
	}
	var ability3 *float32 = nil
	if a3, _ := domain.AbilityIdSecondary().Get(); a3 != nil {
		*ability3 = float32(*a3)
	}
	type2 := domain.TypeSecondary().EnglishName()
	return server.Pokemon{
		Ability1:         float32(*ability1),
		Ability2:         ability2,
		BaseStatsH:       float32(domain.BaseStatsH()),
		BaseStatsA:       float32(domain.BaseStatsA()),
		BaseStatsB:       float32(domain.BaseStatsB()),
		BaseStatsC:       float32(domain.BaseStatsC()),
		BaseStatsD:       float32(domain.BaseStatsD()),
		BaseStatsS:       float32(domain.BaseStatsS()),
		EnglishName:      domain.EnglishName(),
		FormId:           float32(domain.FormNo()),
		FormName:         domain.FormName(),
		HiddenAbility:    ability3,
		Id:               float32(domain.Id()),
		IsFinalEvolution: domain.IsFinalEvolution(),
		Name:             domain.Name(),
		PokedexNo:        float32(domain.PokedexNo()),
		Type1:            domain.TypePrimary().EnglishName(),
		Type2:            &type2,
	}
}
