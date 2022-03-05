package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/users"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/repository/conv"
)

var _ repository.IUserRepository = (*UserRepository)(nil)

var (
	emptyUserBuilder  = func() schema.User { return schema.User{} }
	emptyUsersBuilder = func() []schema.User { return []schema.User{} }
)

type UserRepository struct {
	BaseReadRepository[schema.User, users.User, users.Users, identifier.UserId]
	BaseWriteRepository[schema.User, users.User, identifier.UserId]
	dbClient orm.IDbClient
}

func NewUserRepository(dbClient orm.IDbClient) *UserRepository {
	return &UserRepository{
		BaseReadRepository: BaseReadRepository[schema.User, users.User, users.Users, identifier.UserId]{
			dbClient:            dbClient,
			emptySchemaBuilder:  emptyUserBuilder,
			emptySchemasBuilder: emptyUsersBuilder,
			domainsConstructor:  users.NewUsers,
			toSchemaConverter:   conv.ToSchemaUser,
		},
		BaseWriteRepository: BaseWriteRepository[schema.User, users.User, identifier.UserId]{
			dbClient:           dbClient,
			emptySchemaBuilder: emptyUserBuilder,
			toSchemaConverter:  conv.ToSchemaUser,
			toDomainConverter:  conv.ToDomainUser,
		},
		dbClient: dbClient,
	}
}

func (rep UserRepository) FindByName(targetName string, filterFields ...string) (*users.User, error) {
	users, err := rep.FindByField("Name", targetName, filterFields...)
	if err != nil {
		// Todo: err wrap
		return nil, err
	}
	user := users.First()
	return &user, err
}
