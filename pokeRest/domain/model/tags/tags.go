package tags

type Tags struct {
	items []Tag
}

func (t Tags) Items() []Tag {
	return t.items
}
