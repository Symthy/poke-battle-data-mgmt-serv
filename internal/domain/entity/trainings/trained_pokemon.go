package trainings

import (
	"github.com/Symthy/PokeRest/internal/domain/value"
	"github.com/Symthy/PokeRest/internal/domain/value/identifier"
)

type TrainedPokemon struct {
	id          identifier.TrainedPokemonId
	gender      value.Gender
	nickname    string
	description string
	isPrivate   bool
	userId      identifier.UserId
	TrainedPokemonAdjustment
}

func NewTrainedPokemon(
	id identifier.TrainedPokemonId, gender value.Gender, nickname, description string,
	isPrivate bool, userId identifier.UserId, adjustment TrainedPokemonAdjustment) TrainedPokemon {
	entity := TrainedPokemon{
		id:          id,
		nickname:    nickname,
		gender:      gender,
		description: description,
		isPrivate:   isPrivate,
		userId:      userId,
	}
	entity.TrainedPokemonAdjustment = adjustment
	return entity
}

func (t TrainedPokemon) Id() identifier.TrainedPokemonId {
	return t.id
}

func (t TrainedPokemon) UserId() identifier.UserId {
	return t.userId
}

func (t TrainedPokemon) IsUnregister() bool {
	return t.id.IsEmpty()
}

func (t TrainedPokemon) Notify(note ITrainedPokemonNotification) {
	note.SetId(t.id)
	note.SetGender(t.gender)
	note.SetNickname(t.nickname)
	note.SetDescription(t.description)
	note.SetIsPrivate(t.isPrivate)
	note.SetUserId(t.userId)
	// note.SetAdjustment(t.TrainedPokemonAdjustment)
}
