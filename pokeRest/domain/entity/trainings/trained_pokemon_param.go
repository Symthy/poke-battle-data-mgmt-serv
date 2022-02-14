package trainings

type TrainedPokemonParam struct {
	id           uint
	adjustmentId *uint
	nickname     *string
	gender       Gender
	description  *string
	isPrivate    bool
	userId       uint
}

func NewTrainedPokemonParam(id uint) TrainedPokemonParam {
	// Todo: add and validate
	return TrainedPokemonParam{id: id}
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

func (t TrainedPokemonParam) Gender() Gender {
	return t.gender
}

func (t TrainedPokemonParam) Description() *string {
	return t.description
}

func (t TrainedPokemonParam) IsPrivate() bool {
	return t.isPrivate
}

func (t TrainedPokemonParam) UserId() uint {
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

type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
	None   Gender = "-"
)

func resolveGender(value string) Gender {
	return Gender(value)
}
