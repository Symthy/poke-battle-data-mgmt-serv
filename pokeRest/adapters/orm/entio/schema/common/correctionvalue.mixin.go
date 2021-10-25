package common

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/entio/property"
)

type CorrectionValueMixin struct {
	mixin.Schema
}

func (CorrectionValueMixin) Fields() []ent.Field {
	return []ent.Field{
		// 技威力
		field.Int("physical_move_power_correction_value").Default(1),
		field.Int("special_move_power_correction_value").Default(1),
		// 攻撃力
		field.Int("attack_power_correction_value").Default(1),
		field.Int("special_attack_power_correction_value").Default(1),
		// 攻撃威力
		field.Int("attack_correction_value").Default(1),
		field.Int("special_attack_correction_value").Default(1),
		// 防御力
		field.Int("deffense_correction_value").Default(1),
		field.Int("special_deffense_correction_value").Default(1),
		// ダメージ倍率
		field.String("damage_correction_type1").GoType(property.Types("")).NotEmpty(),
		field.Int("damage_correction_value1").Default(1),
		field.String("damage_correction_type2").GoType(property.Types("")).Optional(),
		field.Int("damage_correction_value2").Optional(),
		//　重さ
		field.Int("weight_correction_value").Default(1),
	}
}
