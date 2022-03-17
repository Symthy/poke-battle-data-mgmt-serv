package mixin

import "github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"

type BattleEffects struct {
	Corrections []Correction `json:"corrections"`
}

type Correction struct {
	Target           enum.CorrectionTarget `json:"target"`
	Value            float32               `json:"value"`
	TriggerCondition *TriggerCondition     `json:"triggerCondition"`
}

type TriggerCondition struct {
	Entry enum.ConditionEntry `json:"entry"`
	Value string              `json:"value"`
}
