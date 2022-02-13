package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/moves"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/repository/dto"
)

var _ repository.IMoveRepository = (*MoveRepository)(nil)

var (
	emptyMoveBuilder  = func() schema.Move { return schema.Move{} }
	emptyMovesBuilder = func() []schema.Move { return []schema.Move{} }
)

type MoveRepository struct {
	dbClient orm.IDbClient
	BaseReadRepository[schema.Move, moves.Move, moves.Moves]
}

func NewMoveRepository(dbClient orm.IDbClient) *MoveRepository {
	return &MoveRepository{
		dbClient: dbClient,
		BaseReadRepository: BaseReadRepository[schema.Move, moves.Move, moves.Moves]{
			dbClient:            dbClient,
			emptySchemaBuilder:  emptyMoveBuilder,
			emptySchemasBuilder: emptyMovesBuilder,
			domainsConstructor:  moves.NewMoves,
			schemaConverter:     dto.ToSchemaMove,
		},
	}
}

func (r MoveRepository) FindOfPokemon(pokemonId uint) (*moves.Moves, error) {
	return r.FindByField("PokemonId", string(rune(pokemonId)))
}

func (r MoveRepository) FindAll(page int, pageSize int) (*moves.Moves, error) {
	return nil, nil
}
