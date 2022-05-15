package value

type PokemonActualValues struct {
	actualValueH uint16
	actualValueA uint16
	actualValueB uint16
	actualValueC uint16
	actualValueD uint16
	actualValueS uint16
}

func NewPokemonActualValues(h, a, b, c, d, s uint16) PokemonActualValues {
	return PokemonActualValues{
		actualValueH: h,
		actualValueA: a,
		actualValueB: b,
		actualValueC: c,
		actualValueD: d,
		actualValueS: s,
	}
}

func (p PokemonActualValues) H() uint16 {
	return p.actualValueH
}
func (p PokemonActualValues) A() uint16 {
	return p.actualValueA
}
func (p PokemonActualValues) B() uint16 {
	return p.actualValueB
}
func (p PokemonActualValues) C() uint16 {
	return p.actualValueC
}
func (p PokemonActualValues) D() uint16 {
	return p.actualValueD
}
func (p PokemonActualValues) S() uint16 {
	return p.actualValueS
}

func (p PokemonActualValues) ApplyCorrection(appliers map[PokemonParam]correctionApplier) PokemonActualValues {
	corrected := NewPokemonActualValues(
		p.H(),
		appliers[A](p.A()),
		appliers[B](p.B()),
		appliers[C](p.C()),
		appliers[D](p.D()),
		appliers[S](p.S()),
	)
	return corrected
}
