package parties

import (
	"github.com/Symthy/PokeRest/pokeRest/application/service/parties/command"
	"github.com/Symthy/PokeRest/pokeRest/application/service/parties/spec"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/parties"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

type PartyWriteService struct {
	repo repository.IPartyRepository
	spec spec.PartySpecification
}

func NewPartyWriteService(
	repo repository.IPartyRepository, userRepo repository.IUserRepository) PartyWriteService {
	serv := PartyWriteService{
		repo: repo,
		spec: spec.NewPartySpecification(userRepo),
	}
	return serv
}

func (s PartyWriteService) SaveParty(cmd command.CreatePartyCommand) (*parties.Party, error) {
	domain, err := cmd.BuildDomain()
	if err != nil {
		return nil, err
	}
	if err := s.spec.ValidateForCreate(domain); err != nil {
		return nil, err
	}
	return s.repo.Create(domain)
}

func (s PartyWriteService) UpdateParty(cmd command.UpdatePartyCommand) (*parties.Party, error) {
	domain, err := cmd.BuildDomain()
	if err != nil {
		return nil, err
	}
	if err := s.spec.ValidateForUpdate(domain); err != nil {
		return nil, err
	}
	return s.repo.Update(domain)
}

func (s PartyWriteService) DeleteParty(cmd command.DeletePartyCommand) (*parties.Party, error) {
	domain, err := cmd.BuildDomain()
	if err != nil {
		return nil, err
	}
	if err := s.spec.ValidateForUpdate(domain); err != nil {
		return nil, err
	}
	return s.repo.Delete(domain.Id().Value())
}
