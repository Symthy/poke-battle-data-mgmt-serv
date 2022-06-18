package value

import "github.com/Symthy/PokeRest/pokeRest/common/lists"

type MoveSpecies string

const (
	MoveSpeciesPhysical MoveSpecies = "Physical" // 物理
	MoveSpeciesSpecial  MoveSpecies = "Special"  // 特殊
	MoveSpeciesStatus   MoveSpecies = "Status"   // 変化
	MoveSpeciesUnknown  MoveSpecies = ""
)

func allMoveSpecies() []MoveSpecies {
	return []MoveSpecies{MoveSpeciesPhysical, MoveSpeciesSpecial, MoveSpeciesStatus}
}

func NewMoveSpecies(species string) MoveSpecies {
	if lists.Contains(allMoveSpecies(), species) {
		return MoveSpecies(species)
	}
	return MoveSpeciesUnknown
}

func (m MoveSpecies) IsPhysical() bool {
	return m == MoveSpeciesPhysical
}

func (m MoveSpecies) IsSpecial() bool {
	return m == MoveSpeciesSpecial
}

func (m MoveSpecies) IsStatus() bool {
	return m == MoveSpeciesStatus
}

func (m MoveSpecies) String() string {
	return string(m)
}
