package entity

type IDomains[T IDomain] interface {
	Items() []T
}

// Todo: temporary
type IDomain interface {
	Id() uint
}
