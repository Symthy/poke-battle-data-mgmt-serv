package parties

import (
	"github.com/Symthy/PokeRest/pokeRest/application/service/parties/command"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/parties"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

type PartyWriteService struct {
	repo repository.IPartyRepository
}

func NewPartyWriteService(repo repository.IPartyRepository) PartyWriteService {
	serv := PartyWriteService{repo: repo}
	return serv
}

func (s PartyWriteService) SaveParty(cmd command.CreatePartyCommand) (*parties.Party, error) {
	domain, err := cmd.BuildDomain()
	if err != nil {
		return nil, err
	}
	return s.repo.Create(*domain)
}

func (s PartyWriteService) UpdateParty(cmd command.UpdatePartyCommand) (*parties.Party, error) {
	domain, err := cmd.BuildDomain()
	if err != nil {
		return nil, err
	}
	return s.repo.Update(*domain)
}

func (s PartyWriteService) DeleteParty(id uint) (*parties.Party, error) {
	return s.repo.Delete(id)
}
