package damages

import (
	"github.com/Symthy/PokeRest/pokeRest/application/service/damages/command"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

type DamageCalculationService struct {
	pokemonRepo repository.IPokemonRepository
}

func Calculate(cmd command.DamageCalculation) {

	return
}
