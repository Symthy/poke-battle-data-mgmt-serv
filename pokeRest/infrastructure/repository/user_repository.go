package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/users"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure"
)

var _ repository.IUserRepository = (*UserRepository)(nil)

var emptyUserBuilder = func() schema.User { return schema.User{} }

type UserRepository struct {
	BaseReadRepository[schema.User, users.User]
	BaseWriteRepository[schema.User, users.User]
	dbClient orm.IDbClient
}

func NewUserRepository(dbClient orm.IDbClient) *UserRepository {
	return &UserRepository{
		BaseReadRepository: BaseReadRepository[schema.User, users.User]{
			dbClient:           dbClient,
			emptySchemaBuilder: emptyUserBuilder,
			schemaConverter:    infrastructure.ToSchemaUser,
		},
		BaseWriteRepository: BaseWriteRepository[schema.User, users.User]{
			dbClient:           dbClient,
			emptySchemaBuilder: emptyUserBuilder,
			schemaConverter:    infrastructure.ToSchemaUser,
		},
		dbClient: dbClient,
	}
}

func (rep UserRepository) FindByName(targetName string, filterFields ...string) (*users.User, error) {
	return rep.FindByField("name", targetName, filterFields...)
}
