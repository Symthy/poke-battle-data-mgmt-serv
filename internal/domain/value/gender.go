package value

const (
	Male       Gender = "Male"
	Female     Gender = "Female"
	NoneGender Gender = ""
)

type Gender string

func (g Gender) String() string {
	return string(g)
}
