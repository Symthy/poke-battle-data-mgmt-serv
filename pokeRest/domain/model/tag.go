package model

type Tag struct {
	id              uint
	name            string
	isGenerationTag bool
	isSeasonTag     bool
}

func (t Tag) Id() uint {
	return t.id
}
