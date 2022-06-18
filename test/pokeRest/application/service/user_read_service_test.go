package service

import (
	"testing"

	"github.com/Symthy/PokeRest/pokeRest/application/service/users"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/repository/conv"
	"github.com/Symthy/PokeRest/test/data"
	"github.com/Symthy/PokeRest/test/mock/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	suite.Suite
	serv users.UserReadService
}

// Before
func (suite *UserServiceTestSuite) SetupTest() {
	suite.serv = users.NewUserReadService(repository.NewUserRepositoryMock())
}

// After
func (suite *UserServiceTestSuite) TearDownTest() {
}

func TestUserControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}

func (suite *UserServiceTestSuite) TestGetUser() {
	var id uint64 = 1
	expected, _ := conv.ToDomainUser(data.DummyUser1())

	actual, err := suite.serv.GetUserById(id)
	if err != nil {
		suite.Fail(err.Error())
	}
	assert.EqualValues(suite.T(), expected, actual)
}

func (suite *UserServiceTestSuite) TestGetUserByName() {

}
