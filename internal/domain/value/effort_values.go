package value

import "github.com/Symthy/PokeRest/internal/errs"

type EffortValues struct {
	h EffortValue
	a EffortValue
	b EffortValue
	c EffortValue
	d EffortValue
	s EffortValue
}

func NewEffortValues(valueH, valueA, valueB, valueC, valueD, valueS uint64) (*EffortValues, error) {
	total := uint16(valueH + valueA + valueB + valueC + valueD + valueS)
	if total > 510 {
		return nil, errs.ThrowErrorInvalidValue("EffortValues", "total", string(rune(total)))
	}
	return NewEffortValuesIfValidIndividual(valueH, valueA, valueB, valueC, valueD, valueS)
}

func NewEffortValuesIfValidIndividual(valueH, valueA, valueB, valueC, valueD, valueS uint64) (*EffortValues, error) {
	effortValueH, err := NewEffortValue(valueH, H)
	if err != nil {
		return nil, err
	}
	effortValueA, err := NewEffortValue(valueA, A)
	if err != nil {
		return nil, err
	}
	effortValueB, err := NewEffortValue(valueB, B)
	if err != nil {
		return nil, err
	}
	effortValueC, err := NewEffortValue(valueC, C)
	if err != nil {
		return nil, err
	}
	effortValueD, err := NewEffortValue(valueD, D)
	if err != nil {
		return nil, err
	}
	effortValueS, err := NewEffortValue(valueS, S)
	if err != nil {
		return nil, err
	}

	return &EffortValues{*effortValueH, *effortValueA, *effortValueB, *effortValueC, *effortValueD, *effortValueS}, nil
}

func (e EffortValues) H() EffortValue {
	return e.h
}
func (e EffortValues) A() EffortValue {
	return e.a
}
func (e EffortValues) B() EffortValue {
	return e.b
}
func (e EffortValues) C() EffortValue {
	return e.c
}
func (e EffortValues) D() EffortValue {
	return e.d
}
func (e EffortValues) S() EffortValue {
	return e.s
}
