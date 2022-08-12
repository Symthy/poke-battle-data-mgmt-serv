package parties

type PartyTags struct {
	items []*PartyTag
}

func NewPartyTags(tags []*PartyTag) *PartyTags {
	return &PartyTags{items: tags}
}

func (t PartyTags) Items() []*PartyTag {
	return t.items
}
