package repository

import (
	"github.com/Symthy/PokeRest/internal/adapters/orm"
	"github.com/Symthy/PokeRest/internal/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/internal/domain/entity/users"
	"github.com/Symthy/PokeRest/internal/domain/repository"
	"github.com/Symthy/PokeRest/internal/domain/value/identifier"
	"github.com/Symthy/PokeRest/internal/infrastructure/repository/conv"
)

var _ repository.IUserRepository = (*UserRepository)(nil)

var (
	emptyUserBuilder  = func() *schema.User { return &schema.User{} }
	emptyUsersBuilder = func() []*schema.User { return []*schema.User{} }
)

type UserRepository struct {
	BaseReadRepository[schema.User, users.User, users.Users, identifier.UserId, uint64]
	BaseWriteRepository[schema.User, users.User, identifier.UserId, uint64]
	dbClient orm.IDbClient
}

func NewUserRepository(dbClient orm.IDbClient) *UserRepository {
	return &UserRepository{
		BaseReadRepository: NewBaseReadRepository[schema.User, users.User, users.Users, identifier.UserId, uint64](
			dbClient,
			emptyUserBuilder,
			emptyUsersBuilder,
			users.NewUsers,
			conv.ToSchemaUser,
			conv.ToDomainUser,
		),
		BaseWriteRepository: BaseWriteRepository[schema.User, users.User, identifier.UserId, uint64]{
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
	return user, err
}
