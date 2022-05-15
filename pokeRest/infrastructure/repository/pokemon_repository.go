package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/pokemons"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/repository/conv"
)

var _ repository.IPokemonRepository = (*PokemonRepository)(nil)

var (
	emptyPokemonBuilder  = func() schema.Pokemon { return schema.Pokemon{} }
	emptyPokemonsBuilder = func() []schema.Pokemon { return []schema.Pokemon{} }
)

type PokemonRepository struct {
	BaseReadRepository[schema.Pokemon, pokemons.Pokemon, pokemons.Pokemons, identifier.PokemonId, uint16]
	dbClient orm.IDbClient
}

func NewPokemonRepository(dbClient orm.IDbClient) *PokemonRepository {
	return &PokemonRepository{
		BaseReadRepository: NewBaseReadRepository[schema.Pokemon, pokemons.Pokemon, pokemons.Pokemons, identifier.PokemonId, uint16](
			dbClient, emptyPokemonBuilder, emptyPokemonsBuilder, pokemons.NewPokemons,
			conv.ToSchemaPokemon, conv.ToDomainPokemon,
		),
		dbClient: dbClient,
	}
}

func (rep PokemonRepository) FindById(id uint16) (*pokemons.Pokemon, error) {
	db := rep.dbClient.Db()
	var pokemon = schema.Pokemon{}
	tx := db.First(&pokemon, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return rep.toDomainConverter(pokemon)
}

// Todo: args is condition
func (rep PokemonRepository) FindAll(next uint32, pageSize uint16) (*pokemons.Pokemons, error) {
	db := rep.dbClient.Db()
	var schemas = []schema.Pokemon{}

	paginate := rep.dbClient.Paginate(next, pageSize)
	tx := db.Scopes(paginate).Find(&schemas)

	if tx.Error != nil {
		return nil, tx.Error
	}
	pokemonDomains, err := conv.ConvertToDomains[schema.Pokemon, pokemons.Pokemon, identifier.PokemonId, uint16](schemas, rep.toDomainConverter)
	if err != nil {
		return nil, err
	}
	pokemonList := pokemons.NewPokemons(pokemonDomains)
	return &pokemonList, nil
}

func (rep PokemonRepository) FindByMove(moveId uint16) (*pokemons.Pokemons, error) {
	return nil, nil
}
