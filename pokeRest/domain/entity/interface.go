package entity

type IDomains[T IDomain[K, I], K IValueId[I], I uint16 | uint32 | uint64] interface {
	Items() []T
}

// Todo: temporary
type IDomain[K IValueId[I], I uint16 | uint32 | uint64] interface {
	Id() K
}

type IValueId[I uint16 | uint32 | uint64] interface {
	Value() I
}
