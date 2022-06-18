package identifier

type ValueId[T uint16 | uint32 | uint64] struct {
	value T
}

func NewValueId[T uint16 | uint32 | uint64](value T) ValueId[T] {
	return ValueId[T]{value}
}

func newEmptyId[T uint16 | uint32 | uint64]() ValueId[T] {
	return ValueId[T]{0}
}

func (id ValueId[T]) Value() T {
	return id.value
}

func (id ValueId[T]) IsEmpty() bool {
	return id.value == 0
}
