package database

import (
	"database/sql"
	"fmt"
	"reflect"
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
		abilityId1 := &sql.NullInt16{}
		abilityId1.Scan(3)
		abilityId2 := &sql.NullInt16{}
		abilityId2.Scan(nil)
		abilityId3 := &sql.NullInt16{}
		abilityId3.Scan(100)

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
			AbilityId1:       *abilityId1,
			AbilityId2:       *abilityId2,
			HiddenAbilityId:  *abilityId3,
			IsFinalEvolution: isFinalEvolution,
		}

		os := test.ConvertToObjectStructure(dto)
		suite.mock.ExpectQuery(regexp.QuoteMeta(
			`SELECT * FROM "pokemons" WHERE "pokemons"."id" = $1 ORDER BY "pokemons"."id" LIMIT 1`)).
			WithArgs(dto.ID).
			WillReturnRows(sqlmock.NewRows(os.Fields()).
				AddRow(3, 3, 1, "Standard", "フシギバナ", "", 1, []byte(enum.Grass), []byte(enum.Poison),
					*abilityId1, *abilityId2, *abilityId3, isFinalEvolution))

		expected := dto.ConvertToDomain()
		actual := suite.pokemonRepository.FindById(dto.ID)

		if !reflect.DeepEqual(expected, actual) {
			fmt.Printf("expected:\n%#v\n", expected)
			fmt.Printf("actual:\n%#v\n", actual)
			// v, _ := actual.AbilityIdPrimary().Get()
			// fmt.Printf("%#v\n", *v)
			suite.Fail("期待値と異なる")
		}
	})
}
