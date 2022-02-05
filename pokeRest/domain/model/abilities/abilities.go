package abilities

type Abilities struct {
	abilities []Ability
}

func (a Abilities) Items() []Ability {
	return a.abilities
}
