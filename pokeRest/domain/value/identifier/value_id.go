package identifier

type ValueId struct {
	value uint
}

func (id ValueId) Value() uint {
	return id.value
}

func (id ValueId) IsEmpty() bool {
	return id.value == 0
}
