package entity

type IDomains[T IDomain[K], K IValueId] interface {
	Items() []T
}

// Todo: temporary
type IDomain[K IValueId] interface {
	Id() K
}

type IValueId interface {
	Value() uint
}
