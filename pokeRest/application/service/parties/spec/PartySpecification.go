package spec

import (
	"fmt"

	"github.com/Symthy/PokeRest/pokeRest/domain/entity/parties"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

type PartySpecification struct {
	userRepo repository.IUserRepository
}

func NewPartySpecification(userRepo repository.IUserRepository) PartySpecification {
	return PartySpecification{userRepo: userRepo}
}

func (s PartySpecification) ValidateForCreate(party *parties.Party) error {
	if !party.IsUnregister() {
		// todo: replace err
		return fmt.Errorf("invalid party: specified id")
	}
	if err := s.validateSatisfy(party); err != nil {
		return err
	}
	return nil
}

func (s PartySpecification) ValidateForUpdate(party *parties.Party) error {
	if party.IsUnregister() {
		// todo: replace err
		return fmt.Errorf("invalid party: non specified id")
	}
	if err := s.validateSatisfy(party); err != nil {
		return err
	}
	return nil
}

func (s PartySpecification) validateSatisfy(party *parties.Party) error {
	user, err := s.userRepo.FindById(party.UserId().Value())
	if err != nil {
		return err
	}
	if user == nil {
		// todo: replace err
		return fmt.Errorf("invalid party: user not found")
	}
	return nil
}
