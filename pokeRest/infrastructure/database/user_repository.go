package database

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/model"
)

type UserRepository struct {
	dbClient orm.IDbClient
}

func NewUserRepository(dbClient orm.IDbClient) *PokemonRepository {
	return &PokemonRepository{dbClient: dbClient}
}

func (rep UserRepository) FindById(id uint) (model.User, error) {
	db := rep.dbClient.Db()
	var user = schema.User{}
	tx := db.First(&user, id)
	return user.ConvertToDomain(), tx.Error
}

func (rep UserRepository) Create(user *model.User) (model.User, error) {
	schemaUser := gormio.ConvertUserToSchema(*user)
	db := rep.dbClient.Db()
	tx := db.Create(&schemaUser)
	created := schemaUser.ConvertToDomain()
	return created, tx.Error
}
