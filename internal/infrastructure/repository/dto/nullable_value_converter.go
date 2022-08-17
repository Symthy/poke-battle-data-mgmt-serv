package dto

import (
	"database/sql"

	"github.com/Symthy/PokeRest/internal/infrastructure"
)

// Todo:
func convertIdToNullInt16[T uint16 | uint32 | uint64](id infrastructure.IValueId[T]) sql.NullInt16 {
	value := id.Value()
	var nullInt sql.NullInt16
	if value == 0 {
		nullInt = sql.NullInt16{
			Int16: 0,
			Valid: false,
		}
	} else {
		nullInt = sql.NullInt16{
			Int16: int16(value),
			Valid: true,
		}
	}
	return nullInt
}
