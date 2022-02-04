package model

type Tag struct {
	id              uint
	name            string
	isGenerationTag bool
	isSeasonTag     bool
}

func NewTag(id uint, name string, isGenerationTag, isSeasonTag bool) Tag {
	return Tag{
		id:              id,
		name:            name,
		isGenerationTag: isGenerationTag,
		isSeasonTag:     isSeasonTag,
	}
}

func (t Tag) Id() uint {
	return t.id
}
