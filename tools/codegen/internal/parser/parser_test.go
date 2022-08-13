package parser

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/Symthy/PokeRest/tools/codegen/internal"
	"github.com/stretchr/testify/assert"
)

var (
	_, thisPath, _, _ = runtime.Caller(0)
)

func TestParse(t *testing.T) {
	cases := []struct {
		inputFilePath string
		expected      []*internal.FieldInfo
	}{
		{
			inputFilePath: filepath.Join(filepath.Dir(thisPath), "../../test/input", "move.go"),
			expected: []*internal.FieldInfo{
				internal.NewFieldInfo("id", internal.NewTypeInfo("identifier", "MoveId")),
				internal.NewFieldInfo("name", internal.NewStandardTypeInfo("string")),
				internal.NewFieldInfo("description", internal.NewPointerTypeInfo("", "string")),
				internal.NewFieldInfo("effects", internal.NewPointerTypeInfo("battles", "BattleEffects")),
				internal.NewFieldInfo("acquiredMember", internal.NewPointerArrayTypeInfo("", "Member")),
				internal.NewFieldInfo("characters", internal.NewArrayTypeInfo("", "Character")),
				internal.NewFieldInfo("leanablePokemons", internal.NewPointerArrayTypeInfo("identifier", "PokemonId")),
				internal.NewEmbeddedFieldInfo(internal.NewTypeInfo("mixin", "UpdateTimes")),
			},
		},
	}

	for _, tt := range cases {
		fields, err := Parse(tt.inputFilePath)
		assert.NoError(t, err)
		assert.EqualValues(t, tt.expected, fields)
	}
}
