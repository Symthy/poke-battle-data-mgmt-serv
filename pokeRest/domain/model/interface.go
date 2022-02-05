package model

type IDomains[T IDomain] interface {
	Items() []T
}

type IDomain interface {
	Id() uint
}
