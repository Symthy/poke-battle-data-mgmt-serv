package parties

import (
	"math"

	"github.com/Symthy/PokeRest/internal/domain/entity"
	"github.com/Symthy/PokeRest/internal/domain/entity/battles"
	"github.com/Symthy/PokeRest/internal/domain/value/identifier"
	"github.com/Symthy/PokeRest/internal/errs"
)

var _ entity.IDomain[identifier.PartyBattleResultId, uint64] = (*PartyBattleResult)(nil)

type PartyBattleResult struct {
	id               identifier.PartyBattleResultId
	maxRate          uint16
	seasonRanking    uint16
	maxSeasonRanking uint16
	winCount         uint16
	loseCount        uint16
	battles.Season
}

func NewPartyBattleResult(
	id identifier.PartyBattleResultId, maxRate, seasonRanking, maxSeasonRanking, winCount, loseCount uint64,
	season battles.Season) (*PartyBattleResult, error) {
	if maxRate > 2500 {
		return nil, errs.ThrowErrorInvalidValue("PartyBattleResult", "maxRate", string(rune(maxRate)))
	}
	if seasonRanking > math.MaxUint16 {
		return nil, errs.ThrowErrorInvalidValue("PartyBattleResult", "seasonRanking", string(rune(seasonRanking)))
	}
	if maxSeasonRanking > math.MaxUint16 {
		return nil, errs.ThrowErrorInvalidValue("PartyBattleResult", "maxSeasonRanking", string(rune(maxSeasonRanking)))
	}
	if winCount > math.MaxUint16 {
		return nil, errs.ThrowErrorInvalidValue("PartyBattleResult", "winCount", string(rune(winCount)))
	}
	if loseCount > math.MaxUint16 {
		return nil, errs.ThrowErrorInvalidValue("PartyBattleResult", "loseCount", string(rune(loseCount)))
	}
	return &PartyBattleResult{
		id:               id,
		maxRate:          uint16(maxRate),
		seasonRanking:    uint16(seasonRanking),
		maxSeasonRanking: uint16(maxSeasonRanking),
		winCount:         uint16(winCount),
		loseCount:        uint16(loseCount),
		Season:           season,
	}, nil
}

func (p PartyBattleResult) Id() identifier.PartyBattleResultId {
	return p.id
}

func (p PartyBattleResult) getWinCount() uint16 {
	return p.winCount
}

func (p PartyBattleResult) getLoseCount() uint16 {
	return p.loseCount
}
