package tags

type Tags struct {
	items []Tag
}

func NewTags(tags []Tag) Tags {
	return Tags{items: tags}
}

func (t Tags) Items() []Tag {
	return t.items
}
