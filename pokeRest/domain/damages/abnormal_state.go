package damages

import (
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
		correction: battles.NewNonCorrectionValue(),
	}
}

func resolveCorrection(state AbnormalStateType) *battles.BattleCorrectionValue {
	if state == AbnormalStateBurn {
		return battles.NewDefaultCorrectionValue(battles.CorrectionPhysicalMove, 2048, nil)
	}
	if state == AbnormalStateParalysis {
		return battles.NewDefaultCorrectionValue(battles.CorrectionStatusS, 2048, nil)
	}
	return battles.NewNonCorrectionValue()
}

func (s AbnormalState) CorrectedValue() uint16 {

	return s.correction.Apply(4096, nil, s.side)
}

func (s AbnormalState) IsParalysis() bool {
	return s.state == AbnormalStateParalysis
}

func (s AbnormalState) IsBurn() bool {
	return s.state == AbnormalStateBurn
}
