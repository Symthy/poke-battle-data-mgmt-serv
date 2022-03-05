package trainings

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

var _ entity.IDomain[identifier.TrainedPokemonId] = (*TrainedPokemonParam)(nil)

type TrainedPokemonParam struct {
	id           identifier.TrainedPokemonId
	gender       value.Gender
	nickname     string
	description  string
	adjustmentId identifier.TrainedAdjustmentId
	isPrivate    bool
	userId       identifier.UserId
}

func NewTrainedPokemonParam(id identifier.TrainedPokemonId, gender value.Gender, nickname, description string,
	adjustmentId identifier.TrainedAdjustmentId, isPrivate bool, userId identifier.UserId,
) TrainedPokemonParam {
	// Todo: factory
	return TrainedPokemonParam{
		id:           id,
		gender:       gender,
		adjustmentId: adjustmentId,
		nickname:     nickname,
		description:  description,
		isPrivate:    isPrivate,
		userId:       userId,
	}
}

// Todo: refactor Notification
func (t TrainedPokemonParam) Id() identifier.TrainedPokemonId {
	return t.id
}

func (t TrainedPokemonParam) AdjustmentId() identifier.TrainedAdjustmentId {
	return t.adjustmentId
}

func (t TrainedPokemonParam) ApplyAdjustmentId(adjustmentId identifier.TrainedAdjustmentId) {
	t.adjustmentId = adjustmentId
}

func (t TrainedPokemonParam) Notify(note ITrainedPokemonParamNotification) {
	note.SetId(t.id)
	note.SetGender(t.gender)
	note.SetNickname(t.nickname)
	note.SetDescription(t.description)
	note.SetAdjustmentId(t.adjustmentId)
	note.SetIsPrivate(t.isPrivate)
	note.SetUserId(t.userId)
}
