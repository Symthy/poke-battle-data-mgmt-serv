package abilities

import (
	"github.com/Symthy/PokeRest/pokeRest/application/service"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/abilities"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type as = abilities.Abilities
type a = abilities.Ability
type ai = identifier.AbilityId

type AbilityReadService struct {
	service.PokemonStatsFinder[as, a, ai]
	service.EntityAllFinder[as, a, ai]
	repo repository.IAbilityRepository
}

func NewAbilityReadService(repo repository.IAbilityRepository) AbilityReadService {
	serv := AbilityReadService{repo: repo}
	serv.PokemonStatsFinder = service.NewPokemonStatsFinder[as, a, ai](repo)
	serv.EntityAllFinder = service.NewEntityAllFinder[as, a, ai](repo)
	return serv
}

// UC: 特性取得
func (s AbilityReadService) FindAbility(id uint) (*abilities.Ability, error) {
	return s.repo.FindById(id)
}

// UC: 特性保持者取得
// FindOfPokemon <- PokemonStatsFinder

// UC: 特性一覧取得
// FindAll <- EntityAllFinder
