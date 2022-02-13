package parties

type PartyTag struct {
	id              uint
	name            string
	isGenerationTag bool
	isSeasonTag     bool
}

func NewPartyTag(id uint, name string, isGenerationTag, isSeasonTag bool) PartyTag {
	return PartyTag{
		id:              id,
		name:            name,
		isGenerationTag: isGenerationTag,
		isSeasonTag:     isSeasonTag,
	}
}

func NewPartyTagOfUnregistered(name string) PartyTag {
	return PartyTag{
		name:            name,
		isGenerationTag: false,
		isSeasonTag:     false,
	}
}

func (t PartyTag) Id() uint {
	return t.id
}

func (t PartyTag) Name() string {
	return t.name
}

func (t PartyTag) IsGenerationTag() bool {
	return t.isGenerationTag
}

func (t PartyTag) IsSeasonTag() bool {
	return t.isSeasonTag
}
