package database

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/database"
	"github.com/stretchr/testify/suite"
)

type PokemonRepositoryTestSuite struct {
	suite.Suite
	pokemonRepository database.PokemonRepository
	mock              sqlmock.Sqlmock
}
