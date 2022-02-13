package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/pokemons"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/repository/dto"
)

var _ repository.IPokemonRepository = (*PokemonRepository)(nil)

var emptyPokemonBuilder = func() schema.Pokemon { return schema.Pokemon{} }

type PokemonRepository struct {
	BaseWriteRepository[schema.Pokemon, pokemons.Pokemon]
	dbClient orm.IDbClient
}

func NewPokemonRepository(dbClient orm.IDbClient) *PokemonRepository {
	return &PokemonRepository{
		BaseWriteRepository: BaseWriteRepository[schema.Pokemon, pokemons.Pokemon]{
			dbClient:           dbClient,
			emptySchemaBuilder: emptyPokemonBuilder,
			schemaConverter:    dto.ToSchemaPokemon,
		},
		dbClient: dbClient,
	}
}

func (rep PokemonRepository) FindById(id uint) (*pokemons.Pokemon, error) {
	db := rep.dbClient.Db()
	var pokemon = schema.Pokemon{}
	tx := db.First(&pokemon, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	p := pokemon.ConvertToDomain()
	return &p, tx.Error
}

// Todo: args is condition
func (rep PokemonRepository) FindAll(next int, pageSize int) (*pokemons.Pokemons, error) {
	db := rep.dbClient.Db()
	var schemas = []schema.Pokemon{}

	paginate := rep.dbClient.Paginate(next, pageSize)
	tx := db.Scopes(paginate).Find(&schemas)

	if tx.Error != nil {
		return nil, tx.Error
	}
	pokemonDomains := dto.ConvertToDomains[schema.Pokemon, pokemons.Pokemon](schemas)
	pokemonList := pokemons.NewPokemons(pokemonDomains)
	return &pokemonList, nil
}

func (rep PokemonRepository) FindByMove(moveId uint) (*pokemons.Pokemons, error) {
	return nil, nil
}
