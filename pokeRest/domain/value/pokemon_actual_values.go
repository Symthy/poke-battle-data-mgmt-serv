package value

type PokemonActualValues struct {
	actualValueH int
	actualValueA int
	actualValueB int
	actualValueC int
	actualValueD int
	actualValueS int
}

func NewPokemonActualValues(h int, a int, b int, c int, d int, s int) PokemonActualValues {
	return PokemonActualValues{
		actualValueH: h,
		actualValueA: a,
		actualValueB: b,
		actualValueC: c,
		actualValueD: d,
		actualValueS: s,
	}
}

func (p PokemonActualValues) H() int {
	return p.actualValueH
}
func (p PokemonActualValues) A() int {
	return p.actualValueA
}
func (p PokemonActualValues) B() int {
	return p.actualValueB
}
func (p PokemonActualValues) C() int {
	return p.actualValueC
}
func (p PokemonActualValues) D() int {
	return p.actualValueD
}
func (p PokemonActualValues) S() int {
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
