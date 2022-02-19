package value

// 努力値
type EffortValue struct {
	value int
}

func NewEffortValue(value int) EffortValue {
	v := value
	if value < 0 {
		v = 0 // log: warn
	}
	if value > 252 {
		v = 252 // log: warn
	}
	return EffortValue{v}
}

func (e EffortValue) Value() int {
	return e.value
}
