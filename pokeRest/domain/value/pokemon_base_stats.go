package value

type PokemonBaseStats struct {
	value int
}

func NewPokemonBaseStats(value int) PokemonBaseStats {
	if value < 0 {
		return PokemonBaseStats{0}
	} else if value > 252 {
		return PokemonBaseStats{252}
	}
	return PokemonBaseStats{value}
}

func (s PokemonBaseStats) Value() int {
	return s.value
}
