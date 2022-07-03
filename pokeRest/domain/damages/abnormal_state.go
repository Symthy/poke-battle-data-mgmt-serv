package damages

import (
	"github.com/Symthy/PokeRest/pokeRest/common/fmath"
	"github.com/Symthy/PokeRest/pokeRest/common/lists"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/battles"
)

type AbnormalStateType string

const (
	AbnormalStatePoison    AbnormalStateType = "poison"
	AbnormalStateParalysis AbnormalStateType = "paralysis"
	AbnormalStateBurn      AbnormalStateType = "burn"
	AbnormalStateFreeze    AbnormalStateType = "freeze"
	AbnormalStateNone      AbnormalStateType = ""
)

type AbnormalState struct {
	state      AbnormalStateType
	correction *battles.BattleCorrectionValue
	side       battles.BattleSideType
}

func allAbnormalStates() []AbnormalStateType {
	return []AbnormalStateType{AbnormalStatePoison, AbnormalStateParalysis, AbnormalStateBurn, AbnormalStateFreeze}
}

func NewAbnormalState(state string) AbnormalState {
	if lists.Contains(allAbnormalStates(), state) {
		return AbnormalState{
			state:      AbnormalStateType(state),
			correction: resolveCorrection(AbnormalStateType(state)),
		}
	}
	return AbnormalState{
		state:      AbnormalStateNone,
		correction: battles.NewBattleNonCorrectionValue(),
	}
}

func resolveCorrection(state AbnormalStateType) *battles.BattleCorrectionValue {
	if state == AbnormalStateBurn {
		return battles.NewDefaultCorrectionValue(battles.CorrectionPhysicalMove, 2048, nil)
	}
	if state == AbnormalStateParalysis {
		return battles.NewDefaultCorrectionValue(battles.CorrectionStatusS, 2048, nil)
	}
	return battles.NewBattleNonCorrectionValue()
}

func (s AbnormalState) ApplyCorrection(damage uint16) uint16 {
	correctionValue := s.correction.Apply(4096, nil, s.side)
	return fmath.RoundUpIfDecimalGreaterFive[uint16](float64(damage * correctionValue))
}
