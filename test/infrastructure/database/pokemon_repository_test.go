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
	dbClient   *orm.GormDbClient
	repository database.PokemonRepository
	mock       sqlmock.Sqlmock
}

// Before
func (suite *PokemonRepositoryTestSuite) SetupTest() {
	db, mock, _ := mock.GetGormDBMock()
	suite.mock = mock
	suite.dbClient = orm.NewGormDbClientForTesting(db)
	pokemonRepository := database.NewPokemonRepository(suite.dbClient)
	suite.repository = *pokemonRepository
}

// After
func (suite *PokemonRepositoryTestSuite) TearDownTest() {
	suite.dbClient.Close()
}

func TestPokemonRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(PokemonRepositoryTestSuite))
}

// test dummy data
var abilityId1 = &sql.NullInt16{Int16: 3, Valid: true}
var abilityId2 = &sql.NullInt16{Int16: 0, Valid: false}
var abilityId3 = &sql.NullInt16{Int16: 100, Valid: true}
var dummyPokemonDto = schema.PokemonDto{
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
	IsFinalEvolution: true,
}

func (suite *PokemonRepositoryTestSuite) TestFind() {
	suite.Run("find by id", func() {
		os := test.Convert(dummyPokemonDto)
		suite.mock.ExpectQuery(regexp.QuoteMeta(
			`SELECT * FROM "pokemons" WHERE "pokemons"."id" = $1 ORDER BY "pokemons"."id" LIMIT 1`)).
			WithArgs(dummyPokemonDto.ID).
			WillReturnRows(sqlmock.NewRows(os.Fields()).
				AddRow(3, 3, 1, "Standard", "フシギバナ", "", 1, []byte(enum.Grass), []byte(enum.Poison),
					*abilityId1, *abilityId2, *abilityId3, true))

		expected := dummyPokemonDto.ConvertToDomain()
		actual := suite.repository.FindById(dummyPokemonDto.ID)

		if !reflect.DeepEqual(expected, actual) {
			fmt.Printf("expected:\n%#v\n", expected)
			fmt.Printf("actual:\n%#v\n", actual)
			suite.Fail("expected and actual is unmatched")
		}
	})
}

func (suite *PokemonRepositoryTestSuite) TestSave() {
	newId := 3
	suite.mock.ExpectBegin()
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "pokemons" ("pokedex_no","form_no","form_name","name","english_name",` + `"generation","type1","type2","ability_id1","ability_id2","hidden_ability_id",` + `"is_final_evolution") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) RETURNING "id"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(newId))
	suite.mock.ExpectCommit()

	pokemon := dummyPokemonDto.ConvertToDomainNonId()
	err := suite.repository.Create(&pokemon)

	if err != nil {
		suite.Fail("Error")
	}
	if pokemon.Id() != int(newId) {
		suite.Fail("invalid id of resisted pokemon")
	}
}

// fmt.Printf("%#v\n", *v)
