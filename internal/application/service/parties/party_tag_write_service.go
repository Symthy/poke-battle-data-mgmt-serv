package parties

import (
	"github.com/Symthy/PokeRest/internal/domain/entity/parties"
	"github.com/Symthy/PokeRest/internal/domain/repository"
)

type PartyTagWriteService struct {
	repo repository.IPartyTagRepository
}

func NewPartyTagWriteService(repo repository.IPartyTagRepository) PartyTagWriteService {
	serv := PartyTagWriteService{repo: repo}
	return serv
}

func (s PartyTagWriteService) SavePartyTag(name string) (*parties.PartyTag, error) {
	partyTag := parties.NewPartyTagOfUnregistered(name)
	return s.repo.Create(partyTag)
}

func (s PartyTagWriteService) DeletePartyTag(id uint64) (*parties.PartyTag, error) {
	return s.repo.Delete(id)
}
