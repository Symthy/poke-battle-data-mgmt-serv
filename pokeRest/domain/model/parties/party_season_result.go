package parties

type PartySeasonResult struct {
	id uint
}

func (p PartySeasonResult) Id() uint {
	return p.id
}
