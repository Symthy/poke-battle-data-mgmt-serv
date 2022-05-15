package factory

import "github.com/Symthy/PokeRest/pokeRest/domain/entity"

type IDoaminFactory[K entity.IValueId[I], I uint16 | uint32 | uint64] interface {
	BuildDomain() (entity.IDomain[K, I], error)
}
