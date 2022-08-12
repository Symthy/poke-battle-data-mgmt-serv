package battles

import "github.com/Symthy/PokeRest/pokeRest/common/lists"

type ConditionEntry string

func (c ConditionEntry) ToString() string {
	return string(c)
}

const (
	ConditionPokemonType       ConditionEntry = "pokemonType"
	ConditionTypeCompatibility ConditionEntry = "typeCompatibility"
	ConditionTypeMatch         ConditionEntry = "isTypeMatch"
	ConditionGender            ConditionEntry = "gender"
	ConditionNotHaveItem       ConditionEntry = "notHaveItem"
	ConditionUnknown           ConditionEntry = ""
)

func allConditionEntry() []ConditionEntry {
	return []ConditionEntry{ConditionPokemonType, ConditionTypeCompatibility, ConditionTypeMatch, ConditionGender}
}

func NewConditionEntry(entry string) ConditionEntry {
	if lists.Contains(allConditionEntry(), entry) {
		return ConditionEntry(entry)
	}
	return ConditionUnknown
}

func (c ConditionEntry) String() string {
	return string(c)
}
