package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/property"
)

// Abilities holds the schema definition for the Abilities entity.
type Abilities struct {
	ent.Schema
}

// Fields of the Abilities.
func (Abilities) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("description").Optional(),
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

// Edges of the Abilities.
func (Abilities) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("ability_holder1", Pokemons.Type).
			Ref("ability1"),
		edge.From("ability_holder2", Pokemons.Type).
			Ref("ability2"),
		edge.From("hidden_ability_holder", Pokemons.Type).
			Ref("hidden_ability"),
	}
}
