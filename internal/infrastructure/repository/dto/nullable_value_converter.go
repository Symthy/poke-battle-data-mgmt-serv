package dto

import (
	"database/sql"

	"github.com/Symthy/PokeRest/internal/infrastructure"
)

func convertIdToNullInt16[T uint16 | uint32 | uint64](id infrastructure.IValueId[T]) sql.NullInt16 {
	value := id.Value()
	nullInt := sql.NullInt16{}
	if value == 0 {
		nullInt.Scan(nil)
	} else {
		nullInt.Scan(value)
	}
	return nullInt
}
