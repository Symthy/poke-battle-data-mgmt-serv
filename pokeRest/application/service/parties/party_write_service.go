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
	return s.repo.Create(cmd.ToDomain())
}

func (s PartyWriteService) UpdateParty(cmd command.UpdatePartyCommand) (*parties.Party, error) {
	return s.repo.Update(cmd.ToDomain())
}

func (s PartyWriteService) DeleteParty(id uint) (*parties.Party, error) {
	return s.repo.Delete(id)
}
