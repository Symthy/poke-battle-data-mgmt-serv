package factory

import (
	"github.com/Symthy/PokeRest/internal/domain/entity/battles"
	"github.com/Symthy/PokeRest/internal/domain/entity/parties"
	"github.com/Symthy/PokeRest/internal/domain/value/identifier"
)

type PartyBattleResultInput struct {
	id               uint64
	generation       uint64
	series           uint64
	season           uint64
	maxRate          uint64
	seasonRanking    uint64
	maxSeasonRanking uint64
	winCount         uint64
	loseCount        uint64
}

func NewPartyBattleResultInput(id, generation, series, season, maxRate, seasonRanking,
	maxSeasonRanking, winCount, loseCount uint64) PartyBattleResultInput {
	return PartyBattleResultInput{
		id, generation, series, season, maxRate, seasonRanking, maxSeasonRanking, winCount, loseCount,
	}
}

func NewPartyBattleResultBuilder() *PartyBattleResultInput {
	return &PartyBattleResultInput{}
}

func (i *PartyBattleResultInput) Id(id uint64) {
	i.id = id
}
func (i *PartyBattleResultInput) Generation(generation uint64) {
	i.generation = generation
}
func (i *PartyBattleResultInput) Series(series uint64) {
	i.series = series
}
func (i *PartyBattleResultInput) Season(season uint64) {
	i.season = season
}
func (i *PartyBattleResultInput) MaxRate(maxRate uint64) {
	i.maxRate = maxRate
}
func (i *PartyBattleResultInput) SeasonRanking(seasonRanking uint64) {
	i.seasonRanking = seasonRanking
}
func (i *PartyBattleResultInput) MaxSeasonRanking(maxSeasonRanking uint64) {
	i.maxSeasonRanking = maxSeasonRanking
}
func (i *PartyBattleResultInput) WinCount(winCount uint64) {
	i.winCount = winCount
}
func (i *PartyBattleResultInput) LoseCount(loseCount uint64) {
	i.loseCount = loseCount
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
	domain, err := parties.NewPartyBattleResult(*id, i.maxRate, i.seasonRanking,
		i.maxSeasonRanking, i.winCount, i.loseCount, *season)
	if err != nil {
		return nil, err
	}
	return domain, nil
}
