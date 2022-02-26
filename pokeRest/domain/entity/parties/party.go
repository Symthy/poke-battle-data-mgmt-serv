package parties

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

var _ entity.IDomain[identifier.PartyId] = (*Party)(nil)

// Todo: field change to value model
type Party struct {
	id              identifier.PartyId
	name            string
	battleFormat    value.BattleFormat
	isPrivate       bool
	partyResultIds  []identifier.PartyBattleResultId // Todo: do you need aggregation?
	partyTagIds     []identifier.PartyTagId
	trainedPokemons value.PartyPokemonIds
	userId          identifier.UserId
}

func NewParty(
	id identifier.PartyId, name string, battleFormat string, isPrivate bool, userId identifier.UserId,
	partyResultIds []identifier.PartyBattleResultId, partyTagIds []identifier.PartyTagId,
	trainedPokemonIds value.PartyPokemonIds) Party {
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

func NewPartyOfUnregistered(
	name string, battleFormat string, isPrivate bool, userId identifier.UserId,
	partyTagIds []identifier.PartyTagId, trainedPokemonIds value.PartyPokemonIds) Party {
	// Todo: factory
	return NewParty(identifier.NewEmptyPartyId(), name, battleFormat, isPrivate, userId, nil, partyTagIds, trainedPokemonIds)
}

func NewPartyForUpdated(
	id uint, name string, battleFormat string, isPrivate bool, userId identifier.UserId,
	partyResultIds []identifier.PartyBattleResultId, partyTagIds []identifier.PartyTagId,
	trainedPokemonIds value.PartyPokemonIds) Party {
	// Todo: factory
	return NewParty(identifier.NewEmptyPartyId(), name, battleFormat, isPrivate, userId, partyResultIds, partyTagIds, trainedPokemonIds)
}

func (p Party) Id() identifier.PartyId {
	return p.id
}

func (p Party) UserId() identifier.UserId {
	return p.userId
}

func (p Party) Notify(note IPartyNotification) {
	note.SetId(p.id)
	note.SetName(p.name)
	note.SetBattleFormat(p.battleFormat)
	note.SetPartyResultIds(p.partyResultIds)
	note.SetPartyTagIds(p.partyTagIds)
	note.SetTrainedPokemons(p.trainedPokemons)
	note.SetIsPrivate(p.isPrivate)
	note.SetUserId(p.userId)
}
