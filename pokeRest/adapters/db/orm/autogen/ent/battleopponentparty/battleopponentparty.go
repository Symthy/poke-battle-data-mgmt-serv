// Code generated by entc, DO NOT EDIT.

package battleopponentparty

const (
	// Label holds the string label denoting the battleopponentparty type in the database.
	Label = "battle_opponent_party"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOpponentPokemonId1 holds the string denoting the opponent_pokemon_id1 field in the database.
	FieldOpponentPokemonId1 = "opponent_pokemon_id1"
	// FieldOpponentPokemonId2 holds the string denoting the opponent_pokemon_id2 field in the database.
	FieldOpponentPokemonId2 = "opponent_pokemon_id2"
	// FieldOpponentPokemonId3 holds the string denoting the opponent_pokemon_id3 field in the database.
	FieldOpponentPokemonId3 = "opponent_pokemon_id3"
	// FieldOpponentPokemonId4 holds the string denoting the opponent_pokemon_id4 field in the database.
	FieldOpponentPokemonId4 = "opponent_pokemon_id4"
	// FieldOpponentPokemonId5 holds the string denoting the opponent_pokemon_id5 field in the database.
	FieldOpponentPokemonId5 = "opponent_pokemon_id5"
	// FieldOpponentPokemonId6 holds the string denoting the opponent_pokemon_id6 field in the database.
	FieldOpponentPokemonId6 = "opponent_pokemon_id6"
	// EdgeBattleContent holds the string denoting the battle_content edge name in mutations.
	EdgeBattleContent = "battle_content"
	// Table holds the table name of the battleopponentparty in the database.
	Table = "battle_opponent_parties"
	// BattleContentTable is the table that holds the battle_content relation/edge.
	BattleContentTable = "battle_records"
	// BattleContentInverseTable is the table name for the BattleRecords entity.
	// It exists in this package in order to avoid circular dependency with the "battlerecords" package.
	BattleContentInverseTable = "battle_records"
	// BattleContentColumn is the table column denoting the battle_content relation/edge.
	BattleContentColumn = "battle_opponent_party_id"
)

// Columns holds all SQL columns for battleopponentparty fields.
var Columns = []string{
	FieldID,
	FieldOpponentPokemonId1,
	FieldOpponentPokemonId2,
	FieldOpponentPokemonId3,
	FieldOpponentPokemonId4,
	FieldOpponentPokemonId5,
	FieldOpponentPokemonId6,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// OpponentPokemonId1Validator is a validator for the "opponent_pokemon_id1" field. It is called by the builders before save.
	OpponentPokemonId1Validator func(int) error
	// OpponentPokemonId2Validator is a validator for the "opponent_pokemon_id2" field. It is called by the builders before save.
	OpponentPokemonId2Validator func(int) error
	// OpponentPokemonId3Validator is a validator for the "opponent_pokemon_id3" field. It is called by the builders before save.
	OpponentPokemonId3Validator func(int) error
	// OpponentPokemonId4Validator is a validator for the "opponent_pokemon_id4" field. It is called by the builders before save.
	OpponentPokemonId4Validator func(int) error
	// OpponentPokemonId5Validator is a validator for the "opponent_pokemon_id5" field. It is called by the builders before save.
	OpponentPokemonId5Validator func(int) error
	// OpponentPokemonId6Validator is a validator for the "opponent_pokemon_id6" field. It is called by the builders before save.
	OpponentPokemonId6Validator func(int) error
)
