package parties

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

var _ entity.IDomain[identifier.PartyTagId, uint64] = (*PartyTag)(nil)

type PartyTag struct {
	id           identifier.PartyTagId
	name         string
	isGeneration bool
	isSeason     bool
}

// Todo: factory
func NewPartyTag(id identifier.PartyTagId, name string, isGeneration, isSeason bool) PartyTag {
	return PartyTag{
		id:           id,
		name:         name,
		isGeneration: isGeneration,
		isSeason:     isSeason,
	}
}

func NewPartyTagOfUnregistered(name string) *PartyTag {
	return &PartyTag{
		name:         name,
		isGeneration: false,
		isSeason:     false,
	}
}

func (t PartyTag) Id() identifier.PartyTagId {
	return t.id
}

func (t PartyTag) Notify(note IPartyTagNotification) {
	note.SetId(t.id)
	note.SetName(t.name)
	note.SetIsGeneration(t.isGeneration)
	note.SetIsSeason(t.isSeason)
}
