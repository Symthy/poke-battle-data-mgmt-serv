package dto

import "github.com/Symthy/PokeRest/pokeRest/infrastructure"

func BuildDomains[T infrastructure.IDomain[K], TM infrastructure.IDomains[T, K], K infrastructure.IValueId](
	domainArray []T, domainsBuilder func([]T) TM) *TM {
	if len(domainArray) == 0 {
		return nil
	}
	domains := domainsBuilder(domainArray)
	return &domains
}
