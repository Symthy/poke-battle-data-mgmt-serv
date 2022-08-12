package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/moves"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/repository/conv"
)

var _ repository.IMoveRepository = (*MoveRepository)(nil)

var (
	emptyMoveBuilder  = func() *schema.Move { return &schema.Move{} }
	emptyMovesBuilder = func() []*schema.Move { return []*schema.Move{} }
)

type MoveRepository struct {
	dbClient orm.IDbClient
	BaseReadRepository[schema.Move, moves.Move, moves.Moves, identifier.MoveId, uint16]
}

func NewMoveRepository(dbClient orm.IDbClient) *MoveRepository {
	return &MoveRepository{
		dbClient: dbClient,
		BaseReadRepository: NewBaseReadRepository[schema.Move, moves.Move, moves.Moves, identifier.MoveId, uint16](
			dbClient,
			emptyMoveBuilder,
			emptyMovesBuilder,
			moves.NewMoves,
			conv.ToSchemaMove,
			conv.ToDomainMove,
		),
	}
}

// FindById <- BaseReadRepository

// FindAll <- BaseReadRepository

func (r MoveRepository) FindOfPokemon(pokemonId uint16) (*moves.Moves, error) {
	// Todo: implementation
	return nil, nil
}

func (r MoveRepository) FindByName(name string) (*moves.Moves, error) {
	return r.FindByField("Name", name)
}
