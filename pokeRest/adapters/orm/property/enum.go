package property

type Types string

const (
	Normal   Types = "Normal"
	Fighting Types = "Fighting"
	Flying   Types = "Flying"
	Poison   Types = "Poison"
	Ground   Types = "Ground"
	Rock     Types = "Rock"
	Bug      Types = "Bug"
	Ghost    Types = "Ghost"
	Steel    Types = "Steel"
	Fire     Types = "Fire"
	Water    Types = "Water"
	Grass    Types = "Grass"
	Electric Types = "Electric"
	Psychic  Types = "Psychic"
	Ice      Types = "Ice"
	Dragon   Types = "Dragon"
	Dark     Types = "Dark"
	None     Types = "None"
)

func (Types) Values() (values []string) {
	for _, s := range []Types{Normal, Fighting, Flying, Poison, Ground, Rock, Bug, Ghost, Steel, Fire, Water, Grass, Electric, Psychic, Ice, Dragon, Dark, None} {
		values = append(values, string(s))
	}
	return
}
