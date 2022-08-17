package enum

import "database/sql/driver"

// 性別
type Gender string

const (
	Male       Gender = "Male"
	Female     Gender = "Female"
	NoneGender Gender = ""
)

func (ge *Gender) Scan(value interface{}) error {
	*ge = Gender(value.([]byte))
	return nil
}

func (ge Gender) Value() (driver.Value, error) {
	return string(ge), nil
}

func (ge Gender) String() string {
	return string(ge)
}
