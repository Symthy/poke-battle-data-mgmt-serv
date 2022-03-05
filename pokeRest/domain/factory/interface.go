package factory

import "github.com/Symthy/PokeRest/pokeRest/domain/entity"

type IDoaminFactory[K entity.IValueId] interface {
	BuildDomain() (entity.IDomain[K], error)
}
