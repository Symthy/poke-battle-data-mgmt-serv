package moves

import (
	"github.com/Symthy/PokeRest/pokeRest/application/service"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/moves"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

type ms = moves.Moves
type m = moves.Move

type MoveReadService struct {
	service.StatsOfPokemonFinder[ms, m]
	service.EntityAllFinder[ms, m]
	repo repository.IMoveRepository
}

func NewMoveReadService(repo repository.IMoveRepository) MoveReadService {
	serv := MoveReadService{
		repo: repo,
	}
	serv.StatsOfPokemonFinder = service.NewStatsOfPokemonFinder[ms, m](repo)
	serv.EntityAllFinder = service.NewEntityAllFinder[ms, m](repo)
	return serv
}

// UC: 技取得 StatsOfPokemonFinder

// UC: 技一覧取得 EntityAllFinder
