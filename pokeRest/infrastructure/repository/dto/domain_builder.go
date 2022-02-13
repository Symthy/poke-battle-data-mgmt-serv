package dto

import "github.com/Symthy/PokeRest/pokeRest/infrastructure"

func BuildDomains[T infrastructure.IDomain, TL infrastructure.IDomains[T]](
	domainArray []T, domainsBuilder func([]T) TL) *TL {
	if len(domainArray) == 0 {
		return nil
	}
	domains := domainsBuilder(domainArray)
	return &domains
}
