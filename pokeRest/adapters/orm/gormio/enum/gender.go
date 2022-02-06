package enum

import "database/sql/driver"

// 性別
type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
)

func (ge *Gender) Scan(value interface{}) error {
	*ge = Gender(value.([]byte))
	return nil
}

func (ge Gender) Value() (driver.Value, error) {
	return string(ge), nil
}
