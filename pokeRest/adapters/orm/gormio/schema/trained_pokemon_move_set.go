package schema

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/model/pokemons"
	"gorm.io/gorm"
)

type TrainedPokemonMoveSet struct {
	gorm.Model
	MoveId1 *uint // M:1 <- Move
	MoveId2 *uint // M:1 <- Move
	MoveId3 *uint // M:1 <- Move
	MoveId4 *uint // M:1 <- Move

}

func (TrainedPokemonMoveSet) TableName() string {
	return "trained_pokemon_move_sets"
}

func (t TrainedPokemonMoveSet) ConvertToDomain() pokemons.TrainedPokemonMoveSet {
	return pokemons.NewTrainedPokemonMoveSet(t.ID)
}
