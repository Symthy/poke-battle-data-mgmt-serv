package damages

import (
	"github.com/Symthy/PokeRest/internal/application/service/damages/command"
	"github.com/Symthy/PokeRest/internal/domain/repository"
)

type DamageCalculationService struct {
	pokemonRepo repository.IPokemonRepository
}

func Calculate(cmd command.DamageCalculation) {
}
