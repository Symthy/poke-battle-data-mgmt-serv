package mixin

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Symthy/PokeRest/internal/adapters/orm/gormio/enum"
)

type BattleEffects struct {
	Corrections []*Correction `json:"corrections"`
	Overrides   []*Override   `json:"overrides"`
}

func NewBattleEffects() *BattleEffects {
	return &BattleEffects{
		Corrections: []*Correction{},
		Overrides:   []*Override{},
	}
}

func (b *BattleEffects) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		// Todo: ok?
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	effects := BattleEffects{}
	err := json.Unmarshal(bytes, &effects)
	return err
}

func (b BattleEffects) Value() (driver.Value, error) {
	return json.Marshal(b)
}

type Correction struct {
	Target           enum.CorrectionTarget `json:"target"`
	Value            uint16                `json:"value"`
	TriggerCondition *TriggerCondition     `json:"triggerCondition"`
}

type Override struct {
	Target           enum.OverrideTarget `json:"target"`
	Value            string              `json:"value"`
	TriggerCondition *TriggerCondition   `json:"triggerCondition"`
}

type TriggerCondition struct {
	Entry enum.ConditionEntry `json:"entry"`
	Value string              `json:"value"`
}
