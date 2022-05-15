package infrastructure

type ISchema[T IDomain[K, I], K IValueId[I], I uint16 | uint32 | uint64] interface {
	TableName() string
}

type IDomain[K IValueId[I], I uint16 | uint32 | uint64] interface {
	Id() K
}

type IValueId[I uint16 | uint32 | uint64] interface {
	Value() I
}

type IDomains[T IDomain[K, I], K IValueId[I], I uint16 | uint32 | uint64] interface {
	Items() []T
}

type ISingleRecordFinder[T IDomain[K, I], K IValueId[I], I uint16 | uint32 | uint64] interface {
	FindById(id I) (*T, error)
}
