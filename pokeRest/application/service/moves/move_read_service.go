package moves

import (
	"github.com/Symthy/PokeRest/pokeRest/application/service"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/moves"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type ms = moves.Moves
type m = moves.Move
type mi = identifier.MoveId

type MoveReadService struct {
	service.PokemonStatsFinder[ms, m, mi]
	service.EntityAllFinder[ms, m, mi]
	repo repository.IMoveRepository
}

func NewMoveReadService(repo repository.IMoveRepository) MoveReadService {
	serv := MoveReadService{
		repo: repo,
	}
	serv.PokemonStatsFinder = service.NewPokemonStatsFinder[ms, m, mi](repo)
	serv.EntityAllFinder = service.NewEntityAllFinder[ms, m, mi](repo)
	return serv
}

// UC: 技取得
func (s MoveReadService) FindMove(moveId uint) (*moves.Move, error) {
	return s.repo.FindById(moveId)
}

// UC: 技習得者取得
// FindOfPokemon <- PokemonStatsFinder

// UC: 技一覧取得
// FindAll <- EntityAllFinder
