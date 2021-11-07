package database

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/database"
	"github.com/Symthy/PokeRest/test/mock"
	"github.com/stretchr/testify/suite"
)

type PokemonRepositoryTestSuite struct {
	suite.Suite
	dbClient          *orm.GormDbClient
	pokemonRepository database.PokemonRepository
	mock              sqlmock.Sqlmock
}

// Before
func (suite *PokemonRepositoryTestSuite) SetupTest() {
	db, mock, _ := mock.GetGormDBMock()
	suite.mock = mock
	suite.dbClient = orm.NewGormDbClientForTesting(db)
	pokemonRepository := database.NewPokemonRepository(suite.dbClient)
	suite.pokemonRepository = *pokemonRepository
}

// After
func (suite *PokemonRepositoryTestSuite) TearDownTest() {
	suite.dbClient.Close()
}

func TestPokemonRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(PokemonRepositoryTestSuite))
}
