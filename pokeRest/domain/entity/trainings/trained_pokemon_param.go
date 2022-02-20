package trainings

import "github.com/Symthy/PokeRest/pokeRest/domain/value"

type TrainedPokemonParam struct {
	id           uint
	gender       value.Gender
	nickname     *string
	description  *string
	adjustmentId *uint
	isPrivate    bool
	userId       *uint
}

func NewTrainedPokemonParam(id uint, gender string, nickname, description *string,
	adjustmentId *uint, isPrivate bool, userId *uint) TrainedPokemonParam {
	// Todo: add and validate
	return TrainedPokemonParam{
		id:           id,
		gender:       value.Gender(gender),
		adjustmentId: adjustmentId,
		nickname:     nickname,
		description:  description,
		isPrivate:    isPrivate,
		userId:       userId,
	}
}

func (t TrainedPokemonParam) Id() uint {
	return t.id
}

func (t TrainedPokemonParam) AdjustmentId() *uint {
	return t.adjustmentId
}

func (t TrainedPokemonParam) Nickname() *string {
	return t.nickname
}

func (t TrainedPokemonParam) Gender() value.Gender {
	return t.gender
}

func (t TrainedPokemonParam) Description() *string {
	return t.description
}

func (t TrainedPokemonParam) IsPrivate() bool {
	return t.isPrivate
}

func (t TrainedPokemonParam) UserId() *uint {
	return t.userId
}

func (t TrainedPokemonParam) ApplyAdjustmentId(adjustmentId uint) TrainedPokemonParam {
	return TrainedPokemonParam{
		id:           t.id,
		adjustmentId: &adjustmentId,
		nickname:     t.nickname,
		gender:       t.gender,
		description:  t.description,
		isPrivate:    t.isPrivate,
		userId:       t.userId,
	}
}
