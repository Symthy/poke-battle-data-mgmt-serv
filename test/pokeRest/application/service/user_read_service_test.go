package service

import (
	"testing"

	"github.com/Symthy/PokeRest/pokeRest/application/service/users"
	"github.com/Symthy/PokeRest/test/data"
	"github.com/Symthy/PokeRest/test/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	suite.Suite
	serv users.UserReadService
}

// Before
func (suite *UserServiceTestSuite) SetupTest() {
	suite.serv = users.NewUserReadService(mock.NewUserRepositoryMock())
}

// After
func (suite *UserServiceTestSuite) TearDownTest() {
}

func TestUserControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}

func (suite *UserServiceTestSuite) TestGetUser() {
	var id uint = 1
	expected := data.DummyUser1().ConvertToDomain()

	actual, err := suite.serv.GetUserById(id)
	if err != nil {
		suite.Fail(err.Error())
	}
	assert.EqualValues(suite.T(), expected, actual)
}

func (suite *UserServiceTestSuite) TestGetUserByName() {

}
