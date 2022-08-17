package sqltype

import (
	"database/sql"
	"log"
)

func ResolveNullInt16(nullInt sql.NullInt16) *int {
	var ret *int = nil
	value, err := nullInt.Value()
	if err != nil {
		log.Printf("Warning: " + err.Error())
	}

	if value != nil {
		convertVal := int(value.(int64))
		ret = &convertVal
	}
	return ret
}
