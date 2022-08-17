package battles

import "github.com/Symthy/PokeRest/internal/pkg/lists"

type CorrectionTarget string

func (c CorrectionTarget) ToString() string {
	return string(c)
}

const (
	// 技威力への補正
	CorrectionPhysicalMove CorrectionTarget = "PhysicalMove" // 物理技
	CorrectionSpecialMove  CorrectionTarget = "SpecialMove"
	// 威力補正
	CorrectionPhysicalPower CorrectionTarget = "PhysicalPower"
	CorrectionSpecialPower  CorrectionTarget = "SpecialPower"
	// ステータス補正
	CorrectionStatusA CorrectionTarget = "StatusA"
	CorrectionStatusC CorrectionTarget = "StatusC"
	CorrectionStatusB CorrectionTarget = "StatusB"
	CorrectionStatusD CorrectionTarget = "StatusD"
	CorrectionStatusS CorrectionTarget = "StatusS"
	// ダメージ
	CorrectionDamage CorrectionTarget = "Damage"
	// その他
	CorrectionWeight CorrectionTarget = "Weight"
	CorrectionNone   CorrectionTarget = ""
)

func allCorrectionTarget() []CorrectionTarget {
	return []CorrectionTarget{
		CorrectionPhysicalPower,
		CorrectionSpecialMove,
		CorrectionPhysicalPower,
		CorrectionSpecialPower,
		CorrectionStatusA,
		CorrectionStatusC,
		CorrectionStatusB,
		CorrectionStatusD,
		CorrectionStatusS,
		CorrectionWeight,
	}
}

func NewCorrectionTarget(target string) CorrectionTarget {
	if lists.Contains(allCorrectionTarget(), target) {
		return CorrectionTarget(target)
	}
	return CorrectionNone
}

func (c CorrectionTarget) String() string {
	return string(c)
}

func GetStatusCorrectionTargets() []CorrectionTarget {
	return []CorrectionTarget{
		CorrectionStatusA,
		CorrectionStatusC,
		CorrectionStatusB,
		CorrectionStatusD,
		CorrectionStatusS,
		CorrectionWeight,
	}
}

func GetPowerCorrectionTargets() []CorrectionTarget {
	return []CorrectionTarget{
		CorrectionPhysicalPower,
		CorrectionSpecialPower,
	}
}

func GetMovePowerCorrectionTargets() []CorrectionTarget {
	return []CorrectionTarget{
		CorrectionPhysicalMove,
		CorrectionSpecialMove,
	}
}

func GetDamageCorrectionTargets() []CorrectionTarget {
	return []CorrectionTarget{
		CorrectionDamage,
	}
}
