package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/model"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/repository/field"
	"gorm.io/gorm"
)

type UserRepository struct {
	dbClient orm.IDbClient
}

func NewUserRepository(dbClient orm.IDbClient) *UserRepository {
	return &UserRepository{dbClient: dbClient}
}

func (rep UserRepository) FindById(id uint) (*model.User, error) {
	db := rep.dbClient.Db()
	var user = schema.User{}
	tx := db.First(&user, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	u := user.ConvertToDomain()
	return &u, nil
}

func (rep UserRepository) FindByName(targetName string, filterFields ...string) (*model.User, error) {
	db := rep.dbClient.Db()
	var user = schema.User{}
	var tx *gorm.DB
	selectedDbFields := field.ConvertToDbField(filterFields...)
	if len(selectedDbFields) > 0 {
		tx = db.Select(selectedDbFields).Find(&user, "name = ?", targetName)
	} else {
		tx = db.Find(&user, "name = ?", targetName)
	}
	if tx.Error != nil {
		return nil, tx.Error
	}
	u := user.ConvertToDomain()
	return &u, tx.Error
}

func (rep UserRepository) Create(user model.User) (*model.User, error) {
	schemaUser := orm.ConvertUserToSchema(user)
	db := rep.dbClient.Db()
	tx := db.Create(&schemaUser)
	if tx.Error != nil {
		return nil, tx.Error
	}
	created := schemaUser.ConvertToDomain()
	return &created, tx.Error
}
