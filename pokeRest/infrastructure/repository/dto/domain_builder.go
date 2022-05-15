package dto

import "github.com/Symthy/PokeRest/pokeRest/infrastructure"

func BuildDomains[T infrastructure.IDomain[K, I], TM infrastructure.IDomains[T, K, I], K infrastructure.IValueId[I], I uint16 | uint32 | uint64](
	domainArray []T, domainsBuilder func([]T) TM) *TM {
	if len(domainArray) == 0 {
		return nil
	}
	domains := domainsBuilder(domainArray)
	return &domains
}
