package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/model"
	"github.com/thoas/go-funk"
)

type PokemonRepository struct {
	dbClient orm.IDbClient
}

func NewPokemonRepository(dbClient orm.IDbClient) *PokemonRepository {
	return &PokemonRepository{dbClient: dbClient}
}

func (rep PokemonRepository) FindById(id uint) (model.Pokemon, error) {
	db := rep.dbClient.Db()
	var pokemon = schema.Pokemon{}
	tx := db.First(&pokemon, id)
	return pokemon.ConvertToDomain(), tx.Error
}

// Todo: args is condition
func (rep PokemonRepository) FindAll() (model.PokemonList, error) {
	db := rep.dbClient.Db()
	var pokemons = []schema.Pokemon{}

	paginate := rep.dbClient.Paginate(1, 100)
	tx := db.Scopes(paginate).Find(&pokemons)

	if tx.Error != nil {
		return model.PokemonList{}, tx.Error
	}
	pokemonDomains := funk.Map(pokemons, func(p schema.Pokemon) model.Pokemon {
		return p.ConvertToDomain()
	}).([]model.Pokemon)
	pokemonList := model.NewPokemonList(pokemonDomains)
	return pokemonList, nil
}

func (rep PokemonRepository) Create(pokemon *model.Pokemon) (model.Pokemon, error) {
	schemaPokemon := orm.ToSchemaPokemon(*pokemon)
	db := rep.dbClient.Db()
	tx := db.Create(&schemaPokemon)
	created := schemaPokemon.ConvertToDomain()
	return created, tx.Error
}

func (rep PokemonRepository) Update(pokemon model.Pokemon) (model.Pokemon, error) {
	// Todo: args change
	db := rep.dbClient.Db()
	target, err := rep.FindById(pokemon.Id())
	if err != nil {
		return pokemon, err
	}
	db.Model(&target).Updates(pokemon)
	return pokemon, nil
}

func (rep PokemonRepository) Delete(id uint) (model.Pokemon, error) {
	db := rep.dbClient.Db()
	var pokemon = schema.Pokemon{}
	tx := db.Delete(&pokemon, id)
	return pokemon.ConvertToDomain(), tx.Error
}
