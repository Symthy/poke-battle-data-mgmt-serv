package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/db"
	"github.com/Symthy/PokeRest/pokeRest/adapters/db/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/model"
	"github.com/thoas/go-funk"
)

type PokemonRepository struct {
	dbAccessor db.IDbAccessor
}

// Todo: args is condition
func (rep *PokemonRepository) FindAll() model.PokemonList {
	db := rep.dbAccessor.GetDb()
	var pokemons = []schema.Pokemon{}

	paginate := rep.dbAccessor.Paginate(1, 100)
	db.Scopes(paginate).Find(&pokemons)

	pokemonDomains := funk.Map(pokemons, func(p schema.Pokemon) model.Pokemon {
		return p.ConvertToDomain()
	}).([]model.Pokemon)
	return model.NewPokemonList(pokemonDomains)
}

func (rep *PokemonRepository) FindById(id uint) model.Pokemon {
	db := rep.dbAccessor.GetDb()
	var pokemon = schema.Pokemon{}
	db.First(&pokemon, id)
	return pokemon.ConvertToDomain()
}

// Todo: return domain
func (rep *PokemonRepository) Save(pokemon schema.Pokemon) schema.Pokemon {
	db := rep.dbAccessor.GetDb()
	db.Save(&pokemon)
	return pokemon
}

func (rep *PokemonRepository) Update(pokemon schema.Pokemon) schema.Pokemon {
	// Todo: args change
	db := rep.dbAccessor.GetDb()
	target := rep.FindById(pokemon.ID)
	db.Model(&target).Updates(pokemon)
	return pokemon
}

func (rep *PokemonRepository) Delete(id uint) schema.Pokemon {
	db := rep.dbAccessor.GetDb()
	var pokemon = schema.Pokemon{}
	db.Delete(&pokemon, id)
	return pokemon
}
