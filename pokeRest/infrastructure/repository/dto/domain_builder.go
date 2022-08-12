package dto

import "github.com/Symthy/PokeRest/pokeRest/infrastructure"

func BuildDomains[TD infrastructure.IDomain[K, I], TM infrastructure.IDomains[TD, K, I], K infrastructure.IValueId[I], I uint16 | uint32 | uint64](
	domainArray []*TD, domainsBuilder func([]*TD) *TM) *TM {
	if len(domainArray) == 0 {
		return nil
	}
	domains := domainsBuilder(domainArray)
	return domains
}
