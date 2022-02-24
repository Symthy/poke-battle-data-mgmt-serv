package infrastructure

type ISchema[T IDomain] interface {
	ConvertToDomain() T
}

type IDomain interface {
	Id() uint
}

type IDomains[T IDomain] interface {
	Items() []T
}

type ISingleRecordFinder[T IDomain] interface {
	FindById(id uint) (*T, error)
}
