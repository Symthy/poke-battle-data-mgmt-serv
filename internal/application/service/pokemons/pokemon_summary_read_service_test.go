package pokemons

import (
	"testing"

	"github.com/Symthy/PokeRest/internal/infrastructure/repository/conv"
	"github.com/Symthy/PokeRest/test/mock/data"
	"github.com/Symthy/PokeRest/test/mock/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PokemonControllerTestSuite struct {
	suite.Suite
	serv PokemonSummaryReadService
}

// Before
func (suite *PokemonControllerTestSuite) SetupTest() {
	suite.serv = NewPokemonSummaryReadService(repository.NewPokemonRepositoryMock())
}

// After
func (suite *PokemonControllerTestSuite) TearDownTest() {
}

func TestPokemonControllerTestSuite(t *testing.T) {
	suite.Run(t, new(PokemonControllerTestSuite))
}

func (suite *PokemonControllerTestSuite) TestGetPokemon() {
	var id uint64 = 3
	actual, err := suite.serv.FindPokemon(id)
	expected, _ := conv.ToDomainPokemon(data.DummyPokemon3())

	if err != nil {
		suite.Fail(err.Error())
	}
	assert.EqualValues(suite.T(), expected, actual)
}
