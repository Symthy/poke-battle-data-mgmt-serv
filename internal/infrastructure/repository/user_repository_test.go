package repository

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Symthy/PokeRest/internal/adapters/orm"
	"github.com/Symthy/PokeRest/internal/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/internal/infrastructure/repository/conv"
	"github.com/Symthy/PokeRest/test/data"
	"github.com/Symthy/PokeRest/test/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	dbClient   *orm.GormDbClient
	repository UserRepository
	mock       sqlmock.Sqlmock
}

// Before
func (suite *UserRepositoryTestSuite) SetupTest() {
	db, mock, _ := mock.GetGormDBMock()
	suite.mock = mock
	suite.dbClient = orm.NewGormDbClientForTesting(db)
	userRepository := NewUserRepository(suite.dbClient)
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
		var id uint64 = 1
		dummyUser := data.DummyUser1()
		suite.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT 1`)).
			WithArgs(id).
			WillReturnRows(sqlmock.NewRows([]string{"ID", "CreatedAt", "UpdatedAt", "DeletedAt", "Name", "TwitterID", "DisplayName", "Email", "Profile", "Role"}).
				AddRow(1, time.Now(), time.Now(), nil, "dummy_user", "twitterId", "dummy dummy user", "test@test.com", "detail\nprofile", []byte(enum.User)))

		expected, _ := conv.ToDomainUser(dummyUser)
		actual, err := suite.repository.FindById(id)
		if err != nil {
			suite.Fail(err.Error())
		}
		assert.EqualValues(suite.T(), *expected, *actual)
	})

	suite.Run("find by name when non filter", func() {
		dummyUser := data.DummyUser1()
		suite.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE Name = $1 AND "users"."deleted_at" IS NULL`)).
			WithArgs("dummy_user").
			WillReturnRows(sqlmock.NewRows([]string{"ID", "CreatedAt", "UpdatedAt", "DeletedAt", "Name", "TwitterID", "DisplayName", "Email", "Profile", "Role"}).
				AddRow(1, time.Now(), time.Now(), nil, "dummy_user", "twitterId", "dummy dummy user", "test@test.com", "detail\nprofile", []byte(enum.User)))

		expected, _ := conv.ToDomainUser(dummyUser)
		actual, err := suite.repository.FindByName("dummy_user")
		if err != nil {
			suite.Fail(err.Error())
		}
		assert.EqualValues(suite.T(), *expected, *actual)
	})

	suite.Run("find by name when exist filter", func() {
		filterFields := []string{"id", "name", "displayName", "role"}
		dummyUser := data.DummyUser1(filterFields...)
		suite.mock.ExpectQuery(regexp.QuoteMeta(`SELECT "id","name","display_name","role" FROM "users" WHERE Name = $1 AND "users"."deleted_at" IS NULL`)).
			WithArgs("dummy_user").
			WillReturnRows(sqlmock.NewRows([]string{"ID", "Name", "DisplayName", "Role"}).
				AddRow(1, "dummy_user", "dummy dummy user", []byte(enum.User)))

		expected, _ := conv.ToDomainUser(dummyUser)
		actual, err := suite.repository.FindByName("dummy_user", filterFields...)
		if err != nil {
			suite.Fail(err.Error())
		}
		assert.EqualValues(suite.T(), *expected, *actual)
	})
}

func (suite *UserRepositoryTestSuite) TestCreate() {
	var id uint64 = 1
	dummyUser := data.DummyUser1()
	suite.mock.ExpectBegin()
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "users" ("name","twitter_id","display_name","email","profile","role","created_at","updated_at","deleted_at","id") ` +
			`VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING "id"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(id))
	suite.mock.ExpectCommit()

	user, _ := conv.ToDomainUser(dummyUser)
	actual, err := suite.repository.Create(user)
	if err != nil {
		suite.Fail(err.Error())
	}
	assert.Equal(suite.T(), id, actual.Id().Value())
}
