package value

import "github.com/Symthy/PokeRest/pokeRest/common/lists"

type MoveSpecies string

const (
	PhysicalMove MoveSpecies = "Physical" // 物理
	SpecialMove  MoveSpecies = "Special"  // 特殊
	StatusMove   MoveSpecies = "Status"   // 変化
	UnknownMove  MoveSpecies = ""
)

func allMoveSpecies() []MoveSpecies {
	return []MoveSpecies{PhysicalMove, SpecialMove, StatusMove}
}

func NewMoveSpecies(species string) MoveSpecies {
	if lists.Contains(allMoveSpecies(), species) {
		return MoveSpecies(species)
	}
	return UnknownMove
}

func (m MoveSpecies) IsPhysical() bool {
	return m == PhysicalMove
}

func (m MoveSpecies) IsSpecial() bool {
	return m == SpecialMove
}
