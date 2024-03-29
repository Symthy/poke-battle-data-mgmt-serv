package repository

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Symthy/PokeRest/internal/adapters/orm"
	"github.com/Symthy/PokeRest/internal/adapters/orm/gormio/enum"
	"github.com/Symthy/PokeRest/internal/infrastructure/repository/conv"
	"github.com/Symthy/PokeRest/test/mock"
	"github.com/Symthy/PokeRest/test/mock/data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PokemonRepositoryTestSuite struct {
	suite.Suite
	dbClient   *orm.GormDbClient
	repository PokemonRepository
	mock       sqlmock.Sqlmock
}

// Before
func (suite *PokemonRepositoryTestSuite) SetupTest() {
	db, mock, _ := mock.GetGormDBMock()
	suite.mock = mock
	suite.dbClient = orm.NewGormDbClientForTesting(db)
	pokemonRepository := NewPokemonRepository(suite.dbClient)
	suite.repository = *pokemonRepository
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
		var id uint16 = 3
		dummyPokemon := data.DummyPokemon3()
		suite.mock.ExpectQuery(regexp.QuoteMeta(
			`SELECT * FROM "pokemons" WHERE "pokemons"."id" = $1 ORDER BY "pokemons"."id" LIMIT 1`)).
			WithArgs(id).
			WillReturnRows(sqlmock.NewRows([]string{"ID", "PokedexNo", "FormNo", "FormName", "Name",
				"EnglishName", "Generation", "Type1", "Type2", "AbilityID1", "AbilityID2",
				"HiddenAbilityID", "BaseStatsH", "BaseStatsA", "BaseStatsB", "BaseStatsC",
				"BaseStatsD", "BaseStatsS", "IsFinalEvolution"}).
				AddRow(3, 3, 1, "Standard", "フシギバナ", "Venusaur", 1, []byte(enum.Grass),
					[]byte(enum.Poison), dummyPokemon.AbilityID1, dummyPokemon.AbilityID2,
					dummyPokemon.HiddenAbilityID, 80, 82, 83, 100, 100, 80, true))

		expected, _ := conv.ToDomainPokemon(dummyPokemon)
		actual, err := suite.repository.FindById(id)
		if err != nil {
			suite.Fail(err.Error())
		}
		assert.EqualValues(suite.T(), *expected, *actual)
	})
}

// func (suite *PokemonRepositoryTestSuite) TestCreate() {
// 	var id uint = 3
// 	dummyPokemon := data.DummyPokemon3()
// 	suite.mock.ExpectBegin()
// 	suite.mock.ExpectQuery(regexp.QuoteMeta(
// 		`INSERT INTO "pokemons" ("pokedex_no","form_no","form_name","name","english_name",` +
// 			`"generation","type1","type2","ability_id1","ability_id2","hidden_ability_id",` +
// 			`"base_stats_h","base_stats_a","base_stats_b","base_stats_c","base_stats_d","base_stats_s",` +
// 			`"is_final_evolution") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18)` +
// 			` RETURNING "id"`)).
// 		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(id))
// 	suite.mock.ExpectCommit()

// 	dummyPokemon.ID = 0 // non id
// 	pokemon, _ := conv.ToDomainPokemon(dummyPokemon)
// 	created, err := suite.repository.Create(*pokemon)

// 	if err != nil {
// 		suite.Fail(err.Error())
// 	}
// 	assert.Equal(suite.T(), id, created.Id().Value())
// }

// fmt.Printf("%#v\n", *v)
