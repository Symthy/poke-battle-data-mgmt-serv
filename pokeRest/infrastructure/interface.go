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
