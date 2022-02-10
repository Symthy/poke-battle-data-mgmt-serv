package controller

import (
	"testing"

	"github.com/Symthy/PokeRest/pokeRest/adapters/di"
	"github.com/Symthy/PokeRest/pokeRest/presentation"
	"github.com/Symthy/PokeRest/pokeRest/presentation/controller"
	"github.com/Symthy/PokeRest/test/data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserControllerTestSuite struct {
	suite.Suite
	controller controller.UserController
}

// Before
func (suite *UserControllerTestSuite) SetupTest() {
	suite.controller = *di.InitUserControllerByRepoMock()
}

// After
func (suite *UserControllerTestSuite) TearDownTest() {
}

func TestUserControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}

func (suite *UserControllerTestSuite) TestGetUser() {
	var id float32 = 1
	expected := presentation.ConvertUserToResponse(data.DummyUser1().ConvertToDomain())

	actual, err := suite.controller.GetUserById(id)
	if err != nil {
		suite.Fail(err.Error())
	}
	assert.EqualValues(suite.T(), expected, actual)
}

func (suite *UserControllerTestSuite) TestGetUserByName() {

}
