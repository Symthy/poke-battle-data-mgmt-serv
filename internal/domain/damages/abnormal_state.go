package damages

import (
	"github.com/Symthy/PokeRest/internal/domain/value/battles"
	"github.com/Symthy/PokeRest/internal/pkg/lists"
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

func NewNotAbnormalState(side battles.BattleSideType) *AbnormalState {
	return &AbnormalState{
		state:      AbnormalStateNone,
		correction: battles.NewNonCorrectionValue(),
		side:       side,
	}
}

func NewAbnormalState(state string, side battles.BattleSideType) *AbnormalState {
	if lists.Contains(allAbnormalStates(), state) {
		return &AbnormalState{
			state:      AbnormalStateType(state),
			correction: resolveCorrection(AbnormalStateType(state)),
			side:       side,
		}
	}
	return &AbnormalState{
		state:      AbnormalStateNone,
		correction: battles.NewNonCorrectionValue(),
		side:       side,
	}
}

func resolveCorrection(state AbnormalStateType) *battles.BattleCorrectionValue {
	if state == AbnormalStateBurn {
		return battles.NewDefaultCorrectionValue(battles.CorrectionPhysicalMove, battles.Correction2048, nil)
	}
	if state == AbnormalStateParalysis {
		return battles.NewDefaultCorrectionValue(battles.CorrectionStatusS, battles.Correction2048, nil)
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
