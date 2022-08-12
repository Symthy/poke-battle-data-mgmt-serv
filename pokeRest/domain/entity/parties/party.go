package parties

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

var _ entity.IDomain[identifier.PartyId, uint64] = (*Party)(nil)

type Party struct {
	id              identifier.PartyId
	name            string
	battleFormat    value.BattleFormat
	isPrivate       bool
	partyResultIds  *value.PartyBattleResultIds
	partyTagIds     *value.PartyTagIds
	trainedPokemons *value.PartyPokemonIds
	userId          identifier.UserId
}

func NewParty(
	id identifier.PartyId, name string, battleFormat value.BattleFormat, isPrivate bool,
	userId identifier.UserId, partyResultIds *value.PartyBattleResultIds,
	partyTagIds *value.PartyTagIds, trainedPokemonIds *value.PartyPokemonIds,
) Party {
	return Party{
		id:              id,
		name:            name,
		battleFormat:    value.BattleFormat(battleFormat),
		partyResultIds:  partyResultIds,
		partyTagIds:     partyTagIds,
		trainedPokemons: trainedPokemonIds,
		isPrivate:       isPrivate,
		userId:          userId,
	}
}

func (p Party) Id() identifier.PartyId {
	return p.id
}

func (p Party) UserId() identifier.UserId {
	return p.userId
}

func (p Party) IsUnregister() bool {
	return p.id.IsEmpty()
}

func (p Party) Notify(note IPartyNotification) {
	note.SetId(p.id)
	note.SetName(p.name)
	note.SetBattleFormat(p.battleFormat)
	note.SetIsPrivate(p.isPrivate)
	// note.SetPartyResultIds(p.partyResultIds)
	// note.SetPartyTagIds(p.partyTagIds)
	// note.SetTrainedPokemons(p.trainedPokemons)
	// note.SetUserId(p.userId)
}
