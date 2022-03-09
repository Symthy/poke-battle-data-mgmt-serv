package factory

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/parties"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type PartyBattleResultInput struct {
	id               uint
	generation       int
	series           int
	season           int
	maxRate          int
	seasonRanking    int
	maxSeasonRanking int
	winCount         int
	loseCount        int
}

func NewPartyBattleResultInput(id uint, generation, series, season, maxRate, seasonRanking,
	maxSeasonRanking, winCount, loseCount int) PartyBattleResultInput {
	return PartyBattleResultInput{
		id, generation, series, season, maxRate, seasonRanking, maxSeasonRanking, winCount, loseCount,
	}
}

func (i PartyBattleResultInput) BuildDomain() (*parties.PartyBattleResult, error) {
	id, err := identifier.NewPartyBattleResultId(i.id)
	if err != nil {
		return nil, err
	}
	season, err := battles.NewSeason(i.generation, i.series, i.season)
	if err != nil {
		return nil, err
	}
	domain := parties.NewPartyBattleResult(*id, i.maxRate, i.seasonRanking,
		i.maxSeasonRanking, i.winCount, i.loseCount, *season)
	return &domain, nil
}
