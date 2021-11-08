package database

import (
	"database/sql"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/database"
	"github.com/Symthy/PokeRest/test"
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

func (suite *PokemonRepositoryTestSuite) TestFindById() {
	suite.Run("find by id", func() {
		//abilityId1 := 1
		//hiddenAbilityId := 100
		isFinalEvolution := true
		dto := schema.PokemonDto{
			ID:               3,
			PokedexNo:        3,
			FormNo:           1,
			FormName:         "Standard",
			Name:             "フシギバナ",
			EnglishName:      "",
			Generation:       1,
			Type1:            enum.Grass,
			Type2:            enum.Poison,
			AbilityId1:       sql.NullInt16{Int16: 1, Valid: true},
			AbilityId2:       sql.NullInt16{Int16: 0, Valid: false},
			HiddenAbilityId:  sql.NullInt16{Int16: 100, Valid: true},
			IsFinalEvolution: isFinalEvolution,
		}

		os := test.ConvertToObjectStructure(dto)
		//fmt.Printf("%#v\n", os)

		suite.mock.ExpectQuery(regexp.QuoteMeta(
			`SELECT * FROM "pokemons" WHERE "pokemons"."id" = $1 ORDER BY "pokemons"."id" LIMIT 1`)).
			WithArgs(dto.ID).
			WillReturnRows(sqlmock.NewRows(os.Fields()).
				AddRow(3, 3, 1, "Standard", "フシギバナ", "", 1, []byte(enum.Grass), []byte(enum.Poison), 1, nil, 100, &isFinalEvolution))

		expected := dto.ConvertToDomain()
		actual := suite.pokemonRepository.FindById(dto.ID)

		if expected != actual {
			fmt.Printf("expected: %#v\n", expected)
			fmt.Printf("actual: %#v\n", actual)
			suite.Fail("期待値と異なる")
		}
	})
}
