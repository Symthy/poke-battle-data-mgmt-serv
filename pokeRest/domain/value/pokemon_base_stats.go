package value

type PokemonBaseStats struct {
	value uint8
}

func NewPokemonBaseStats(value uint64) PokemonBaseStats {
	if value < 0 {
		return PokemonBaseStats{0}
	} else if value > 252 {
		return PokemonBaseStats{252}
	}
	return PokemonBaseStats{uint8(value)}
}

func (s PokemonBaseStats) Value() uint8 {
	return s.value
}
