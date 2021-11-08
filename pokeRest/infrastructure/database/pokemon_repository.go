package database

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

// Todo: args is condition
func (rep PokemonRepository) FindAll() model.PokemonList {
	db := rep.dbClient.Db()
	var pokemons = []schema.PokemonDto{}

	paginate := rep.dbClient.Paginate(1, 100)
	db.Scopes(paginate).Find(&pokemons)

	pokemonDomains := funk.Map(pokemons, func(p schema.PokemonDto) model.Pokemon {
		return p.ConvertToDomain()
	}).([]model.Pokemon)
	return model.NewPokemonList(pokemonDomains)
}

func (rep PokemonRepository) FindById(id int) model.Pokemon {
	db := rep.dbClient.Db()
	var pokemon = schema.Pokemon{}
	db.First(&pokemon, id)
	return pokemon.ConvertToDomain()
}

// Todo: return domain
func (rep PokemonRepository) Save(pokemon model.Pokemon) model.Pokemon {
	db := rep.dbClient.Db()
	db.Save(&pokemon)
	return pokemon
}

func (rep PokemonRepository) Update(pokemon model.Pokemon) model.Pokemon {
	// Todo: args change
	db := rep.dbClient.Db()
	target := rep.FindById(pokemon.Id())
	db.Model(&target).Updates(pokemon)
	return pokemon
}

func (rep PokemonRepository) Delete(id uint) model.Pokemon {
	db := rep.dbClient.Db()
	var pokemon = model.Pokemon{}
	db.Delete(&pokemon, id)
	return pokemon
}
