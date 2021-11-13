package controller

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Symthy/PokeRest/pokeRest/adapters/di"
	"github.com/Symthy/PokeRest/pokeRest/presentation/controller"
	"github.com/Symthy/PokeRest/test/data"
	"github.com/stretchr/testify/suite"
)

type PokemonControllerTestSuite struct {
	suite.Suite
	controller controller.PokemonController
}

// Before
func (suite *PokemonControllerTestSuite) SetupTest() {
	suite.controller = *di.InitPokemonControllerByRepoMock()
}

// After
func (suite *PokemonControllerTestSuite) TearDownTest() {
}

func TestPokemonControllerTestSuite(t *testing.T) {
	suite.Run(t, new(PokemonControllerTestSuite))
}

func (suite *PokemonControllerTestSuite) TestGetPokemon() {
	var id float32 = 3
	actual := suite.controller.GetPokemon(id)
	expected := controller.ConvertToResponse(data.DummyPokemon3().ConvertToDomain())

	if !reflect.DeepEqual(expected, actual) {
		suite.Fail("expected and actual is unmatched")
		fmt.Printf("expected:\n%#v\n", expected)
		fmt.Printf("actual:  \n%#v\n", actual)
	}
}
