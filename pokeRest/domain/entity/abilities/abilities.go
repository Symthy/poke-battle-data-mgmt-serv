package abilities

type Abilities struct {
	abilities []Ability
}

func NewAbilities(abilities []Ability) Abilities {
	return Abilities{abilities: abilities}
}

func (a Abilities) Items() []Ability {
	return a.abilities
}
