package tags

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

func (t Tag) Name() string {
	return t.name
}

func (t Tag) IsGenerationTag() bool {
	return t.isGenerationTag
}

func (t Tag) IsSeasonTag() bool {
	return t.isSeasonTag
}
