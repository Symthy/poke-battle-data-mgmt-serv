package moves

import (
	"github.com/Symthy/PokeRest/pokeRest/application/command"
	"github.com/Symthy/PokeRest/pokeRest/application/service"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/moves"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

type MoveReadService struct {
	service.StatsOfPokemonFinder[moves.Moves, moves.Move]
	repo repository.IMoveRepository
}

func NewMoveReadService(repo repository.IMoveRepository) MoveReadService {
	serv := MoveReadService{
		repo: repo,
	}
	serv.StatsOfPokemonFinder = service.NewStatsOfPokemonFinder[moves.Moves, moves.Move](repo)
	return serv
}

// UC: 技取得

// UC: 技一覧取得
func (s MoveReadService) GetAllMove(cmd command.GetAllEntityCommand) ([]interface{}, error) {
	return nil, nil
}
