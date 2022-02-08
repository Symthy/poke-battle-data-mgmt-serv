package parties

type PartySeasonResult struct {
	id uint
}

func NewPartySeasonResult(id uint) PartySeasonResult {
	return PartySeasonResult{
		id: id,
	}
}

func (p PartySeasonResult) Id() uint {
	return p.id
}
