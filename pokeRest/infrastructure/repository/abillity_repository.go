package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/abilities"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

var _ repository.IAbilityRepository = (*AbilityRepository)(nil)

type AbilityRepository struct {
	dbClient orm.IDbClient
}

func NewAbilityRepository(dbClient orm.IDbClient) *AbilityRepository {
	return &AbilityRepository{dbClient: dbClient}
}

func (r AbilityRepository) FindOfPokemon(pokemonId uint) (*abilities.Abilities, error) {
	return nil, nil
}

func (r AbilityRepository) FindAll(page int, pageSize int) (*abilities.Abilities, error) {
	return nil, nil
}
