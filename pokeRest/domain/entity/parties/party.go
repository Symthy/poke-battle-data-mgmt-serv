package parties

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity"
)

var _ entity.IDomain = (*Party)(nil)

// Todo: field change to value model
type Party struct {
	id                uint
	name              string
	battleFormat      BattleFormat
	isPrivate         bool
	partyResultIds    *[]uint
	partyTagIds       *[]uint
	trainedPokemonIds []uint
	userId            *uint
}

func NewParty(id uint, name string, battleFormat string, isPrivate bool, userId *uint,
	partyResultIds *[]uint, partyTagIds *[]uint, trainedPokemonIds []uint) Party {
	return Party{
		id:                id,
		name:              name,
		battleFormat:      resolveBattleFormat(battleFormat),
		partyResultIds:    partyResultIds,
		partyTagIds:       partyTagIds,
		trainedPokemonIds: trainedPokemonIds,
		isPrivate:         isPrivate,
		userId:            userId,
	}
}

func NewPartyOfUnregistered(name string, battleFormat string, isPrivate bool, userId uint,
	partyTagIds []uint, trainedPokemonIds []uint) Party {
	// Todo: validate
	return NewParty(0, name, battleFormat, isPrivate, &userId, nil, &partyTagIds, trainedPokemonIds)
}

func NewPartyForUpdated(id uint, name string, battleFormat string, isPrivate bool,
	partyResultIds []uint, partyTagIds []uint, trainedPokemonIds []uint) Party {
	// Todo: validate
	return NewParty(0, name, battleFormat, isPrivate, nil, &partyResultIds, &partyTagIds, trainedPokemonIds)
}

func (p Party) Id() uint {
	return p.id
}

func (p Party) Name() string {
	return p.name
}

func (p Party) BattleFormat() BattleFormat {
	return p.battleFormat
}

func (p Party) PartyResultIds() *[]uint {
	return p.partyResultIds
}

func (p Party) PartyTagIds() *[]uint {
	return p.partyTagIds
}

func (p Party) TrainedPokemonIds() []uint {
	return p.trainedPokemonIds
}

func (p Party) IsPrivate() bool {
	return p.isPrivate
}

func (p Party) UserId() *uint {
	return p.userId
}

type BattleFormat string

const (
	Single BattleFormat = "Single"
	Double BattleFormat = "Double"
)

func resolveBattleFormat(value string) BattleFormat {
	if value == "Double" {
		return Double
	}
	return Single
}
