package dto

import (
	"database/sql"

	"github.com/Symthy/PokeRest/pokeRest/infrastructure"
)

func convertIdToNullInt16(id infrastructure.IValueId) sql.NullInt16 {
	value := id.Value()
	nullInt := sql.NullInt16{}
	if value == 0 {
		nullInt.Scan(nil)
	} else {
		nullInt.Scan(value)
	}
	return nullInt
}
