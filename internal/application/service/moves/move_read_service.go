package moves

import (
	"github.com/Symthy/PokeRest/internal/application/service"
	"github.com/Symthy/PokeRest/internal/domain/entity/moves"
	"github.com/Symthy/PokeRest/internal/domain/repository"
	"github.com/Symthy/PokeRest/internal/domain/value/identifier"
)

type ms = moves.Moves
type m = moves.Move
type mi = identifier.MoveId

type MoveReadService struct {
	service.PokemonStatsFinder[ms, m, mi]
	service.EntityAllFinder[ms, m, mi, uint16]
	repo repository.IMoveRepository
}

func NewMoveReadService(repo repository.IMoveRepository) MoveReadService {
	serv := MoveReadService{
		repo: repo,
	}
	serv.PokemonStatsFinder = service.NewPokemonStatsFinder[ms, m, mi](repo)
	serv.EntityAllFinder = service.NewEntityAllFinder[ms, m, mi, uint16](repo)
	return serv
}

// UC: 技取得
func (s MoveReadService) FindMove(moveId uint64) (*moves.Move, error) {
	// Todo: validate id upper limit
	return s.repo.FindById(uint16(moveId))
}

// UC: 技習得者取得
// FindOfPokemon <- PokemonStatsFinder

// UC: 技一覧取得
// FindAll <- EntityAllFinder
