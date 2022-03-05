package infrastructure

type ISchema[T IDomain[K], K IValueId] interface {
	TableName() string
}

type IDomain[K IValueId] interface {
	Id() K
}

type IValueId interface {
	Value() uint
}

type IDomains[T IDomain[K], K IValueId] interface {
	Items() []T
}

type ISingleRecordFinder[T IDomain[K], K IValueId] interface {
	FindById(id uint) (*T, error)
}
