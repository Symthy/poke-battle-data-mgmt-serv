// Code generated by entc, DO NOT EDIT.

package party

const (
	// Label holds the string label denoting the party type in the database.
	Label = "party"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldBattleFormat holds the string denoting the battle_format field in the database.
	FieldBattleFormat = "battle_format"
	// EdgePartyBattleRecord holds the string denoting the party_battle_record edge name in mutations.
	EdgePartyBattleRecord = "party_battle_record"
	// EdgePartyToTag holds the string denoting the party_to_tag edge name in mutations.
	EdgePartyToTag = "party_to_tag"
	// EdgeResultRecord holds the string denoting the result_record edge name in mutations.
	EdgeResultRecord = "result_record"
	// Table holds the table name of the party in the database.
	Table = "parties"
	// PartyBattleRecordTable is the table that holds the party_battle_record relation/edge.
	PartyBattleRecordTable = "battle_records"
	// PartyBattleRecordInverseTable is the table name for the BattleRecords entity.
	// It exists in this package in order to avoid circular dependency with the "battlerecords" package.
	PartyBattleRecordInverseTable = "battle_records"
	// PartyBattleRecordColumn is the table column denoting the party_battle_record relation/edge.
	PartyBattleRecordColumn = "party_id"
	// PartyToTagTable is the table that holds the party_to_tag relation/edge. The primary key declared below.
	PartyToTagTable = "tags_tag_to_party"
	// PartyToTagInverseTable is the table name for the Tags entity.
	// It exists in this package in order to avoid circular dependency with the "tags" package.
	PartyToTagInverseTable = "tags"
	// ResultRecordTable is the table that holds the result_record relation/edge.
	ResultRecordTable = "party_result_records"
	// ResultRecordInverseTable is the table name for the PartyResultRecord entity.
	// It exists in this package in order to avoid circular dependency with the "partyresultrecord" package.
	ResultRecordInverseTable = "party_result_records"
	// ResultRecordColumn is the table column denoting the result_record relation/edge.
	ResultRecordColumn = "party_id"
)

// Columns holds all SQL columns for party fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldName,
	FieldBattleFormat,
}

var (
	// PartyToTagPrimaryKey and PartyToTagColumn2 are the table columns denoting the
	// primary key for the party_to_tag relation (M2M).
	PartyToTagPrimaryKey = []string{"tags_id", "party_id"}
)

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
	// UserIDValidator is a validator for the "user_id" field. It is called by the builders before save.
	UserIDValidator func(int) error
)