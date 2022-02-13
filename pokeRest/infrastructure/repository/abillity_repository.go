package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/abilities"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/repository/dto"
)

var _ repository.IAbilityRepository = (*AbilityRepository)(nil)

var (
	emptyAbilityBuilder   = func() schema.Ability { return schema.Ability{} }
	emptyAbilitiesBuilder = func() []schema.Ability { return []schema.Ability{} }
)

type AbilityRepository struct {
	dbClient orm.IDbClient
	BaseReadRepository[schema.Ability, abilities.Ability, abilities.Abilities]
}

func NewAbilityRepository(dbClient orm.IDbClient) *AbilityRepository {
	return &AbilityRepository{
		dbClient: dbClient,
		BaseReadRepository: BaseReadRepository[schema.Ability, abilities.Ability, abilities.Abilities]{
			dbClient:            dbClient,
			emptySchemaBuilder:  emptyAbilityBuilder,
			emptySchemasBuilder: emptyAbilitiesBuilder,
			domainsConstructor:  abilities.NewAbilities,
			schemaConverter:     dto.ToSchemaAbility,
		},
	}
}

// FindById

func (r AbilityRepository) FindOfPokemon(pokemonId uint) (*abilities.Abilities, error) {
	return r.FindByField("PokemonId", string(rune(pokemonId)))
}

func (r AbilityRepository) FindAll(page int, pageSize int) (*abilities.Abilities, error) {
	return nil, nil
}
