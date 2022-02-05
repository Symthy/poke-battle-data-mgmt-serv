package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/common/collections"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/pokemons"
)

type PokemonRepository struct {
	dbClient orm.IDbClient
}

func NewPokemonRepository(dbClient orm.IDbClient) *PokemonRepository {
	return &PokemonRepository{dbClient: dbClient}
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
func (rep PokemonRepository) FindAll() (pokemons.PokemonList, error) {
	db := rep.dbClient.Db()
	var schemas = []schema.Pokemon{}

	paginate := rep.dbClient.Paginate(1, 100)
	tx := db.Scopes(paginate).Find(&schemas)

	if tx.Error != nil {
		return pokemons.PokemonList{}, tx.Error
	}
	pokemonDomains := collections.ListMap[schema.Pokemon, pokemons.Pokemon](schemas)
	pokemonList := pokemons.NewPokemonList(pokemonDomains)
	return pokemonList, nil
}

func (rep PokemonRepository) Create(pokemon *pokemons.Pokemon) (*pokemons.Pokemon, error) {
	schemaPokemon := orm.ToSchemaPokemon(*pokemon)
	db := rep.dbClient.Db()
	tx := db.Create(&schemaPokemon)
	if tx.Error != nil {
		return nil, tx.Error
	}
	created := schemaPokemon.ConvertToDomain()
	return &created, tx.Error
}

func (rep PokemonRepository) Update(pokemon pokemons.Pokemon) (*pokemons.Pokemon, error) {
	// Todo: args change
	db := rep.dbClient.Db()
	target, err := rep.FindById(pokemon.Id())
	if err != nil {
		return nil, err
	}
	schemaPokemon := schema.Pokemon{}
	tx := db.Model(&target).Updates(&schemaPokemon)
	if tx.Error != nil {
		return nil, tx.Error
	}
	p := schemaPokemon.ConvertToDomain()
	return &p, nil
}

func (rep PokemonRepository) Delete(id uint) (*pokemons.Pokemon, error) {
	db := rep.dbClient.Db()
	var pokemon = schema.Pokemon{}
	tx := db.Delete(&pokemon, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	s := pokemon.ConvertToDomain()
	return &s, nil
}
