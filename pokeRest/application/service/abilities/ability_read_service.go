package abilities

import (
	"github.com/Symthy/PokeRest/pokeRest/application/service"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/abilities"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

type as = abilities.Abilities
type a = abilities.Ability

type AbilityReadService struct {
	service.PokemonStatsFinder[as, a]
	service.EntityAllFinder[as, a]
	repo repository.IAbilityRepository
}

func (s AbilityReadService) NewAbilityReadService(repo repository.IAbilityRepository) AbilityReadService {
	serv := AbilityReadService{repo: repo}
	serv.PokemonStatsFinder = service.NewPokemonStatsFinder[as, a](repo)
	serv.EntityAllFinder = service.NewEntityAllFinder[as, a](repo)
	return serv
}

// UC: 特性保持者取得 PokemonStatsFinder

// UC: 特性一覧取得 EntityAllFinder
