package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/moves"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

var _ repository.IMoveRepository = (*MoveRepository)(nil)

type MoveRepository struct {
	dbClient orm.IDbClient
}

func NewMoveRepository(dbClient orm.IDbClient) *MoveRepository {
	return &MoveRepository{dbClient: dbClient}
}

func (r MoveRepository) FindOfPokemon(pokemonId uint) (*moves.Moves, error) {
	return nil, nil
}

func (r MoveRepository) FindAll(page int, pageSize int) (*moves.Moves, error) {
	return nil, nil
}
