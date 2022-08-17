package ast

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/Symthy/PokeRest/tools/codegen/internal"
	"github.com/stretchr/testify/assert"
)

var (
	_, thisPath, _, _ = runtime.Caller(0)
	testInputFilePath = filepath.Join(filepath.Dir(thisPath), "../../test/input", "move.go")
)

func TestPartImport(t *testing.T) {
	cases := []struct {
		inputFilePath string
		expected      []*internal.ImportInfo
	}{
		{
			inputFilePath: testInputFilePath,
			expected: []*internal.ImportInfo{
				internal.NewImportInfo("", "github.com/Symthy/PokeRest/internal/adapters/orm/gormio/mixin"),
				internal.NewImportInfo("", "github.com/Symthy/PokeRest/internal/domain/value/battles"),
				internal.NewImportInfo("ident", "github.com/Symthy/PokeRest/internal/domain/value/identifier"),
			},
		},
	}

	for _, tt := range cases {
		fileNode, err := parseGoFile(tt.inputFilePath)
		assert.NoError(t, err)
		imports := parseImport(fileNode)
		assert.EqualValues(t, tt.expected, imports)
	}
}

func TestParseStruct(t *testing.T) {
	cases := []struct {
		inputFilePath string
		expected      []*internal.StructInfo
	}{
		{
			inputFilePath: testInputFilePath,
			expected: []*internal.StructInfo{
				internal.NewStructInfo(
					"Move",
					[]*internal.ImportInfo{
						internal.NewImportInfo("", "github.com/Symthy/PokeRest/internal/adapters/orm/gormio/mixin"),
						internal.NewImportInfo("", "github.com/Symthy/PokeRest/internal/domain/value/battles"),
						internal.NewImportInfo("ident", "github.com/Symthy/PokeRest/internal/domain/value/identifier"),
					},
					[]*internal.FieldInfo{
						internal.NewFieldInfo("id", internal.NewTypeInfo("ident", "MoveId")),
						internal.NewFieldInfo("name", internal.NewStandardTypeInfo("string")),
						internal.NewFieldInfo("description", internal.NewPointerTypeInfo("", "string")),
						internal.NewFieldInfo("effects", internal.NewPointerTypeInfo("battles", "BattleEffects")),
						internal.NewFieldInfo("acquiredMember", internal.NewPointerArrayTypeInfo("", "Member")),
						internal.NewFieldInfo("characters", internal.NewArrayTypeInfo("", "Character")),
						internal.NewFieldInfo("leanablePokemons", internal.NewPointerArrayTypeInfo("ident", "PokemonId")),
						internal.NewEmbeddedFieldInfo(internal.NewTypeInfo("mixin", "UpdateTimes")),
					},
				),
				internal.NewStructInfo(
					"Test",
					[]*internal.ImportInfo{
						internal.NewImportInfo("", "github.com/Symthy/PokeRest/internal/adapters/orm/gormio/mixin"),
						internal.NewImportInfo("", "github.com/Symthy/PokeRest/internal/domain/value/battles"),
						internal.NewImportInfo("ident", "github.com/Symthy/PokeRest/internal/domain/value/identifier"),
					},
					[]*internal.FieldInfo{
						internal.NewFieldInfo("test", internal.NewStandardTypeInfo("string")),
					},
				),
			},
		},
	}

	for _, tt := range cases {
		fileNode, err := parseGoFile(tt.inputFilePath)
		assert.NoError(t, err)
		structs := parseStruct(fileNode)
		assert.EqualValues(t, tt.expected, structs)
	}
}
