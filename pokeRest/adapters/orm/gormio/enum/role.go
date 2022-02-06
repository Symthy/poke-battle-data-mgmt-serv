package enum

import "database/sql/driver"

// ロール
type Role string

const (
	User                Role = "User"
	ReferenceableMaster Role = "ReferenceableMaster"
	Administrator       Role = "Administrator"
)

func (ro *Role) Scan(value interface{}) error {
	*ro = Role(value.([]byte))
	return nil
}

func (ro Role) Value() (driver.Value, error) {
	return string(ro), nil
}

func (ro Role) String() string {
	return string(ro)
}
