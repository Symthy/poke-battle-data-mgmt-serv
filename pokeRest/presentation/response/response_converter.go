package response

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/rest/autogen/server"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/abilities"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/items"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/moves"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/pokemons"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/types"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/users"
)

func ConvertPokemonToResponse(domain pokemons.Pokemon) server.Pokemon {
	ability1, _ := domain.AbilityIdPrimary().Get()
	var ability2 *float32 = nil
	if a2, _ := domain.AbilityIdSecondary().Get(); a2 != nil {
		*ability2 = float32(*a2)
	}
	var ability3 *float32 = nil
	if a3, _ := domain.AbilityIdSecondary().Get(); a3 != nil {
		*ability3 = float32(*a3)
	}
	type2 := domain.TypeSecondary().NameEN()
	return server.Pokemon{
		Ability1:      float32(*ability1),
		Ability2:      ability2,
		BaseStatsH:    float32(domain.BaseStatsH()),
		BaseStatsA:    float32(domain.BaseStatsA()),
		BaseStatsB:    float32(domain.BaseStatsB()),
		BaseStatsC:    float32(domain.BaseStatsC()),
		BaseStatsD:    float32(domain.BaseStatsD()),
		BaseStatsS:    float32(domain.BaseStatsS()),
		EnglishName:   domain.EnglishName(),
		FormNo:        float32(domain.FormNo()),
		FormName:      domain.FormName(),
		HiddenAbility: ability3,
		Id:            float32(domain.Id()),
		Name:          domain.Name(),
		PokedexNo:     float32(domain.PokedexNo()),
		Type1:         domain.TypePrimary().NameEN(),
		Type2:         &type2,
	}
}

func ConvertAbilitiesToResponse(domains abilities.Abilities) server.Abilities {
	// Todo: implementation
	return server.Abilities{}
}

func ConvertAbilityToResponse(domain abilities.Ability) server.Ability {
	// Todo: implementation
	return server.Ability{}
}

func ConvertMovesToResponse(domains moves.Moves) server.Moves {
	// Todo: implementation
	return server.Moves{}
}

func ConvertMoveToResponse(domains moves.Move) server.Move {
	// Todo: implementation
	return server.Move{}
}

func ConvertHeldItemsToResponse(domains items.HeldItems) server.HeldItems {
	// Todo: implementation
	return server.HeldItems{}
}

func ConvertHeldItemToResponse(domains items.HeldItem) server.HeldItem {
	// Todo: implementation
	return server.HeldItem{}
}

func ConvertUserToResponse(domain users.User) server.User {
	id := float32(domain.Id())
	name := domain.Name().Value()
	email := domain.Email().Value()
	return server.User{
		Id:          &id,
		Name:        &name,
		DisplayName: domain.DisplayName(),
		Email:       &email,
		Profile:     domain.Profile(),
	}
}

func ConvertTypeTableToResponse(domain types.TypeCompatibilityTable, lang string) server.TypeCompatibilityTable {
	return server.TypeCompatibilityTable{
		CompatibilityTable: domain.GenerateTypeDamageRateTable(),
		TypeOrder:          domain.GenerateTypeNames(lang),
	}
}

func ConvertTypesToResponse(domain types.TypeCompatibility, lang string) server.TypeCompatibility {
	pType := domain.ResolveTargetType(lang)
	return server.TypeCompatibility{
		TargetType:      &pType,
		Compatibilities: domain.GenerateTypeDamageRates(),
		TypeOrder:       domain.GenerateTypeNames(lang),
	}
}
