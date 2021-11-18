package database

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/database"
	"github.com/Symthy/PokeRest/test/data"
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

func (suite *PokemonRepositoryTestSuite) TestFind() {
	suite.Run("find by id", func() {
		var id uint = 3
		dummyPokemon := data.DummyPokemon3()
		suite.mock.ExpectQuery(regexp.QuoteMeta(
			`SELECT * FROM "pokemons" WHERE "pokemons"."id" = $1 ORDER BY "pokemons"."id" LIMIT 1`)).
			WithArgs(id).
			WillReturnRows(sqlmock.NewRows([]string{"ID", "PokedexNo", "FormNo", "FormName", "Name",
				"EnglishName", "Generation", "Type1", "Type2", "AbilityId1", "AbilityId2",
				"HiddenAbilityId", "BaseStatsH", "BaseStatsA", "BaseStatsB", "BaseStatsC",
				"BaseStatsD", "BaseStatsS", "IsFinalEvolution"}).
				AddRow(3, 3, 1, "Standard", "フシギバナ", "Venusaur", 1, []byte(enum.Grass),
					[]byte(enum.Poison), dummyPokemon.AbilityId1, dummyPokemon.AbilityId2,
					dummyPokemon.HiddenAbilityId, 80, 82, 83, 100, 100, 80, true))

		expected := dummyPokemon.ConvertToDomain()
		actual, err := suite.repository.FindById(id)
		if err != nil {
			suite.Fail(err.Error())
		}
		if !reflect.DeepEqual(expected, actual) {
			suite.Fail("expected and actual is unmatched")
			fmt.Printf("expected:\n%#v\n", expected)
			fmt.Printf("actual:  \n%#v\n", actual)
		}
	})
}

func (suite *PokemonRepositoryTestSuite) TestCreate() {
	var id uint = 3
	dummyPokemon := data.DummyPokemon3()
	suite.mock.ExpectBegin()
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "pokemons" ("pokedex_no","form_no","form_name","name","english_name",` +
			`"generation","type1","type2","ability_id1","ability_id2","hidden_ability_id",` +
			`"base_stats_h","base_stats_a","base_stats_b","base_stats_c","base_stats_d","base_stats_s",` +
			`"is_final_evolution") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18)` +
			` RETURNING "id"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(id))
	suite.mock.ExpectCommit()

	pokemon := dummyPokemon.ConvertToDomainNonId()
	created, err := suite.repository.Create(&pokemon)

	if err != nil {
		suite.Fail(err.Error())
	}
	if created.Id() != id {
		suite.Fail("invalid id of created pokemon record")
		fmt.Printf("expected:%v\n", id)
		fmt.Printf("actual:  %v\n", created.Id())
	}
}

// fmt.Printf("%#v\n", *v)
