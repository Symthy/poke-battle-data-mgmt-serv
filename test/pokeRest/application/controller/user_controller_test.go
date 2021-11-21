package controller

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Symthy/PokeRest/pokeRest/adapters/di"
	"github.com/Symthy/PokeRest/pokeRest/presentation"
	"github.com/Symthy/PokeRest/pokeRest/presentation/controller"
	"github.com/Symthy/PokeRest/test/data"
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
	actual, err := suite.controller.GetUserById(id)
	expected := presentation.ConvertUserToResponse(data.DummyUser1().ConvertToDomain())

	if err != nil {
		suite.Fail(err.Error())
	}
	if !reflect.DeepEqual(expected, actual) {
		suite.Fail("expected and actual is unmatched")
		fmt.Printf("expected:\n%#v\n", expected)
		fmt.Printf("actual:  \n%#v\n", actual)
	}
}

func (suite *UserControllerTestSuite) TestGetUserByName() {

}
