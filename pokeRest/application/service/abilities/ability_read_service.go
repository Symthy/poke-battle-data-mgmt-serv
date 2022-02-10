package abilities

import (
	"github.com/Symthy/PokeRest/pokeRest/application/service"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/abilities"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

type as = abilities.Abilities
type a = abilities.Ability

type AbilityReadService struct {
	service.StatsOfPokemonFinder[as, a]
	service.EntityAllFinder[as, a]
	repo repository.IAbilityRepository
}

func (s AbilityReadService) NewAbilityReadService(repo repository.IAbilityRepository) AbilityReadService {
	serv := AbilityReadService{repo: repo}
	serv.StatsOfPokemonFinder = service.NewStatsOfPokemonFinder[as, a](repo)
	serv.EntityAllFinder = service.NewEntityAllFinder[as, a](repo)
	return serv
}

// UC: 特性取得 StatsOfPokemonFinder

// UC: 特性一覧取得 EntityAllFinder