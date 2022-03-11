package factory

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/parties"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type PartyInput struct {
	id                uint
	name              string
	battleFormat      string
	isPrivate         bool
	userId            uint
	partyResultIds    []uint
	partyTagIds       []uint
	trainedPokemonIds []uint
}

func NewPartyInput(id uint, name, battleFormat string, isPrivate bool, userId uint,
	partyResultIds, partyTagIds, trainedPokemonIds []uint) PartyInput {
	return PartyInput{
		id:                id,
		name:              name,
		battleFormat:      battleFormat,
		isPrivate:         isPrivate,
		userId:            userId,
		partyResultIds:    partyResultIds,
		partyTagIds:       partyTagIds,
		trainedPokemonIds: trainedPokemonIds,
	}
}

func NewPartyBuilder() PartyInput {
	return PartyInput{}
}

func (i PartyInput) Id(id uint) {
	i.id = id
}
func (i PartyInput) Name(name string) {
	i.name = name
}
func (i PartyInput) BattleFormat(battleFormat string) {
	i.battleFormat = battleFormat
}
func (i PartyInput) SetIsPrivate(isPrivate bool) {
	i.isPrivate = isPrivate
}
func (i PartyInput) UserId(userId uint) {
	i.userId = userId
}
func (i PartyInput) PartyResultIds(ids []uint) {
	i.partyResultIds = ids
}
func (i PartyInput) PartyTagIds(ids []uint) {
	i.partyTagIds = ids
}
func (i PartyInput) TrainedPokemonIds(ids []uint) {
	i.trainedPokemonIds = ids
}

func (i PartyInput) BuildDomain() (*parties.Party, error) {
	id, err := identifier.NewPartyId(i.id)
	if err != nil {
		return nil, err
	}
	battleFormat := value.BattleFormat(i.battleFormat)
	userId, err := identifier.NewUserId(i.userId)
	if err != nil {
		return nil, err
	}
	partyResultIds := value.NewEmptyPartyBattleResultIds()
	if i.partyResultIds != nil {
		ids, err := value.NewPartyBattleResultIds(i.partyResultIds)
		if err != nil {
			return nil, err
		}
		partyResultIds = *ids
	}
	partyTagIds, err := value.NewPartyTagIds(i.partyTagIds)
	if err != nil {
		return nil, err
	}
	trainedPokemonIds := value.NewPartyPokemonIds(i.trainedPokemonIds)
	if err != nil {
		return nil, err
	}

	domain := parties.NewParty(*id, i.name, battleFormat, i.isPrivate, *userId, partyResultIds,
		*partyTagIds, trainedPokemonIds)
	return &domain, nil
}
