package moves

import (
	"github.com/Symthy/PokeRest/pokeRest/application/service"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/moves"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

type ms = moves.Moves
type m = moves.Move

type MoveReadService struct {
	service.PokemonStatsFinder[ms, m]
	service.EntityAllFinder[ms, m]
	repo repository.IMoveRepository
}

func NewMoveReadService(repo repository.IMoveRepository) MoveReadService {
	serv := MoveReadService{
		repo: repo,
	}
	serv.PokemonStatsFinder = service.NewPokemonStatsFinder[ms, m](repo)
	serv.EntityAllFinder = service.NewEntityAllFinder[ms, m](repo)
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
