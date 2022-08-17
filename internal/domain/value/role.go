package value

type Role string

const (
	User                Role = "User"
	ReferenceableMaster Role = "ReferenceableMaster"
	Administrator       Role = "Administrator"
)

func (ro Role) String() string {
	return string(ro)
}
