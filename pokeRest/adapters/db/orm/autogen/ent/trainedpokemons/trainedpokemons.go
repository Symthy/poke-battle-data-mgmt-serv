// Code generated by entc, DO NOT EDIT.

package trainedpokemons

import (
	"time"
)

const (
	// Label holds the string label denoting the trainedpokemons type in the database.
	Label = "trained_pokemons"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldPokemonID holds the string denoting the pokemon_id field in the database.
	FieldPokemonID = "pokemon_id"
	// FieldCreateUserID holds the string denoting the create_user_id field in the database.
	FieldCreateUserID = "create_user_id"
	// FieldNature holds the string denoting the nature field in the database.
	FieldNature = "nature"
	// FieldEffortValueH holds the string denoting the effort_value_h field in the database.
	FieldEffortValueH = "effort_value_h"
	// FieldEffortValueA holds the string denoting the effort_value_a field in the database.
	FieldEffortValueA = "effort_value_a"
	// FieldEffortValueB holds the string denoting the effort_value_b field in the database.
	FieldEffortValueB = "effort_value_b"
	// FieldEffortValueC holds the string denoting the effort_value_c field in the database.
	FieldEffortValueC = "effort_value_c"
	// FieldEffortValueD holds the string denoting the effort_value_d field in the database.
	FieldEffortValueD = "effort_value_d"
	// FieldEffortValueS holds the string denoting the effort_value_s field in the database.
	FieldEffortValueS = "effort_value_s"
	// EdgeUsePokemon holds the string denoting the use_pokemon edge name in mutations.
	EdgeUsePokemon = "use_pokemon"
	// EdgeTrainingUser holds the string denoting the training_user edge name in mutations.
	EdgeTrainingUser = "training_user"
	// Table holds the table name of the trainedpokemons in the database.
	Table = "trained_pokemons"
	// UsePokemonTable is the table that holds the use_pokemon relation/edge.
	UsePokemonTable = "trained_pokemons"
	// UsePokemonInverseTable is the table name for the Pokemons entity.
	// It exists in this package in order to avoid circular dependency with the "pokemons" package.
	UsePokemonInverseTable = "pokemons"
	// UsePokemonColumn is the table column denoting the use_pokemon relation/edge.
	UsePokemonColumn = "pokemon_id"
	// TrainingUserTable is the table that holds the training_user relation/edge.
	TrainingUserTable = "trained_pokemons"
	// TrainingUserInverseTable is the table name for the Users entity.
	// It exists in this package in order to avoid circular dependency with the "users" package.
	TrainingUserInverseTable = "users"
	// TrainingUserColumn is the table column denoting the training_user relation/edge.
	TrainingUserColumn = "create_user_id"
)

// Columns holds all SQL columns for trainedpokemons fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldPokemonID,
	FieldCreateUserID,
	FieldNature,
	FieldEffortValueH,
	FieldEffortValueA,
	FieldEffortValueB,
	FieldEffortValueC,
	FieldEffortValueD,
	FieldEffortValueS,
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
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// PokemonIDValidator is a validator for the "pokemon_id" field. It is called by the builders before save.
	PokemonIDValidator func(int) error
	// CreateUserIDValidator is a validator for the "create_user_id" field. It is called by the builders before save.
	CreateUserIDValidator func(int) error
	// NatureValidator is a validator for the "nature" field. It is called by the builders before save.
	NatureValidator func(string) error
	// DefaultEffortValueH holds the default value on creation for the "effort_value_h" field.
	DefaultEffortValueH int
	// DefaultEffortValueA holds the default value on creation for the "effort_value_a" field.
	DefaultEffortValueA int
	// DefaultEffortValueB holds the default value on creation for the "effort_value_b" field.
	DefaultEffortValueB int
	// DefaultEffortValueC holds the default value on creation for the "effort_value_c" field.
	DefaultEffortValueC int
	// DefaultEffortValueD holds the default value on creation for the "effort_value_d" field.
	DefaultEffortValueD int
	// DefaultEffortValueS holds the default value on creation for the "effort_value_s" field.
	DefaultEffortValueS int
)