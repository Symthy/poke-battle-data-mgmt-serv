package spec

import (
	"fmt"

	"github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

type TrainedPokemonSpecification struct {
	userRepo repository.IUserRepository
}

func NewTrainedPokemonSpecification(userRepo repository.IUserRepository) TrainedPokemonSpecification {
	return TrainedPokemonSpecification{userRepo: userRepo}
}

func (s TrainedPokemonSpecification) ValidateForCreate(poke trainings.TrainedPokemon) error {
	if !poke.IsUnregister() {
		// todo: replace err
		return fmt.Errorf("invalid trained pokemon: specified id")
	}
	if err := s.validateSatisfy(poke); err != nil {
		return err
	}
	return nil
}

func (s TrainedPokemonSpecification) ValidateForUpdate(poke trainings.TrainedPokemon) error {
	if poke.IsUnregister() {
		// todo: replace err
		return fmt.Errorf("invalid trained pokemon: non specified id")
	}
	if err := s.validateSatisfy(poke); err != nil {
		return err
	}
	return nil
}

func (s TrainedPokemonSpecification) validateSatisfy(poke trainings.TrainedPokemon) error {
	user, err := s.userRepo.FindById(poke.UserId().Value())
	if err != nil {
		return err
	}
	if user == nil {
		// todo: replace err
		return fmt.Errorf("invalid trained pokemon: user not found")
	}
	return nil
}
