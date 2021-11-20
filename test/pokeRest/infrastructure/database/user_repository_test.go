package database

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/database"
	"github.com/Symthy/PokeRest/test/data"
	"github.com/Symthy/PokeRest/test/mock"
	"github.com/stretchr/testify/suite"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	dbClient   *orm.GormDbClient
	repository database.UserRepository
	mock       sqlmock.Sqlmock
}

// Before
func (suite *UserRepositoryTestSuite) SetupTest() {
	db, mock, _ := mock.GetGormDBMock()
	suite.mock = mock
	suite.dbClient = orm.NewGormDbClientForTesting(db)
	userRepository := database.NewUserRepository(suite.dbClient)
	suite.repository = *userRepository
}

// After
func (suite *UserRepositoryTestSuite) TearDownTest() {
	suite.dbClient.Close()
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}

func (suite *UserRepositoryTestSuite) TestFind() {
	suite.Run("find by id", func() {
		var id uint = 1
		dummyUser := data.DummyUser1()
		suite.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT 1`)).
			WithArgs(id).
			WillReturnRows(sqlmock.NewRows([]string{"ID", "CreatedAt", "UpdatedAt", "DeletedAt", "Name", "DisplayName", "Email", "Profile", "Role"}).
				AddRow(1, time.Now(), time.Now(), nil, "dummy_user", "dummy dummy user", "test@test.com", "detail\nprofile", []byte(enum.User)))

		expected := dummyUser.ConvertToDomain()
		actual, err := suite.repository.FindById(id)
		if err != nil {
			suite.Fail(err.Error())
		}
		if !reflect.DeepEqual(expected, actual) {
			suite.Fail("expected and actual is unmatched")
			fmt.Printf("expected:\n%#v\n", expected)
			fmt.Printf("actual:  \n%#v\n", actual)
		}
	})
}

func (suite *UserRepositoryTestSuite) TestCreate() {
	var id uint = 1
	dummyUser := data.DummyUser1()
	suite.mock.ExpectBegin()
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "users" ("created_at","updated_at","deleted_at","name","display_name","email","profile","role") ` +
			`VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING "id"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(id))
	suite.mock.ExpectCommit()

	user := dummyUser.ConvertToDomainNonId()
	actual, err := suite.repository.Create(&user)
	if err != nil {
		suite.Fail(err.Error())
	}
	if actual.Id() != id {
		suite.Fail("invalid id of created user record")
		fmt.Printf("expected:%v\n", id)
		fmt.Printf("actual:  %v\n", actual.Id())
	}
}
