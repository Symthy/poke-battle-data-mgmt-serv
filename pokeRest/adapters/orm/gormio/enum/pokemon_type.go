package enum

import (
	"database/sql/driver"

	"github.com/Symthy/PokeRest/pokeRest/domain/value"
)

// ポケモンタイプ
type PokemonType string

const (
	Normal   PokemonType = "Normal"
	Fighting PokemonType = "Fighting"
	Flying   PokemonType = "Flying"
	Poison   PokemonType = "Poison"
	Ground   PokemonType = "Ground"
	Rock     PokemonType = "Rock"
	Bug      PokemonType = "Bug"
	Ghost    PokemonType = "Ghost"
	Steel    PokemonType = "Steel"
	Fire     PokemonType = "Fire"
	Water    PokemonType = "Water"
	Grass    PokemonType = "Grass"
	Electric PokemonType = "Electric"
	Psychic  PokemonType = "Psychic"
	Ice      PokemonType = "Ice"
	Dragon   PokemonType = "Dragon"
	Dark     PokemonType = "Dark"
	Fairy    PokemonType = "Fairy"
	None     PokemonType = ""
)

func (ty *PokemonType) Scan(value interface{}) error {
	*ty = PokemonType(value.([]byte))
	return nil
}

func (ty PokemonType) Value() (driver.Value, error) {
	return string(ty), nil
}

func (ty PokemonType) String() string {
	return string(ty)
}

func (ty PokemonType) ConvertToDomain() value.PokemonType {
	return value.NewPokemonType(ty.String())
}

func Convert(typ value.PokemonType) PokemonType {
	return PokemonType(typ.ToString())
}
