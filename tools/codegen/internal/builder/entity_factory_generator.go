package builder

import (
	"fmt"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/Symthy/PokeRest/tools/codegen/filesystem"
	"github.com/Symthy/PokeRest/tools/codegen/internal"
	"github.com/Symthy/PokeRest/tools/codegen/internal/ast"
	"github.com/dave/jennifer/jen"
)

const (
	typeUint8       = "uint8"
	typeUint16      = "uint16"
	typeUint32      = "uint32"
	typeUint64      = "uint64"
	typeArrayUint8  = "[]" + typeUint8
	typeArrayUint16 = "[]" + typeUint16
	typeArrayUint32 = "[]" + typeUint32
	typeArrayUint64 = "[]" + typeUint64
	typeString      = "string"
	typeArrayString = "[]string"
)

var (
	entityToFactory = map[string]string{
		"Ability":                     "AbilityInput",
		"BattleRecord":                "BattleRecordInput",
		"BattleOpponentParty":         "BattleOpponentPartyInput",
		"Season":                      "SeasonInput",
		"HeldItem":                    "HeldItemInput",
		"Move":                        "MoveInput",
		"Party":                       "PartyInput",
		"PartyTag":                    "PartyTagInput",
		"PartyBattleResult":           "PartyBattleResultInput",
		"Pokemon":                     "PokemonInput",
		"TrainedPokemon":              "TraindPokemonInput",
		"TrainedPokemonAdjustment":    "TraindPokemonAdjustmentInput",
		"TrainedPokemonAttackTarget":  "TraindPokemonAttackTargetInput",
		"TrainedPokemonDefenseTarget": "TraindPokemonDefenseTargetInput",
		"User":                        "UserInput",
	}

	// Todo: auto gen
	fieldMapping = map[string]string{
		// identifier
		"AbilityId":              typeUint64,
		"BattleOpponentPartyId":  typeUint64,
		"BattleRecordId":         typeUint64,
		"HeldItemId":             typeUint64,
		"MoveId":                 typeUint64,
		"PartyId":                typeUint64,
		"PartyTagId":             typeUint64,
		"PartyBattleResultId":    typeUint64,
		"PokemonId":              typeUint64,
		"TrainedPokemonId":       typeUint64,
		"TrainedAdjustmentId":    typeUint64,
		"TrainedAttackTargetId":  typeUint64,
		"TrainedDefenseTargetId": typeUint64,
		"UserId":                 typeUint64,
		// battles
		"ElectionPokemons":        typeArrayUint64,
		"ElectionTrainedPokemons": typeArrayUint64,
		// value
		"PartyPokemonIds":      typeArrayUint64,
		"MoveSpecies":          typeString,
		"PokemonType":          typeString,
		"BattleFormat":         typeString,
		"BattleResult":         typeString,
		"PartyBattleResultIds": typeArrayUint64,
		"PartyTagIds":          typeArrayUint64,
		"PokedexId":            typeUint64,
		"PokemonBaseStats":     typeUint64,
		"PokemonNature":        typeString,
		"EffortValue":          typeUint64,
		"Gender":               typeString,
		"UserName":             typeString,
		"Role":                 typeString,
	}

	multiFieldStructToPath = map[string]string{
		"PokemonTypeSet":      "pokeRest/domain/value/pokemon_type_set.go",
		"PokemonAbilityIdSet": "pokeRest/domain/value/pokemon_ability_id_set.go",
		"EffortValues":        "pokeRest/domain/value/effort_values.go",
		"PokemonMoveIdSet":    "pokeRest/domain/value/pokemon_move_id_set.go",
	}
)

func GenerateEntityFactoryStructs(rootPath string, homePath string, entityStructs []*internal.StructInfo) error {
	typeNameToMultiFieldStruct, err := parseMultiFieldStructs(rootPath)
	if err != nil {
		return err
	}

	outPath := filepath.Join(homePath, "output/factory")
	if err := filesystem.MakeDirIfNotExists(outPath); err != nil {
		return err
	}

	isErr := false
	for _, st := range entityStructs {
		factoryTypeName, ok := entityToFactory[st.Name()]
		if !ok {
			fmt.Printf("[Warn]  Skip entity: %s\n", st.Name())
			continue
		}

		f := jen.NewFile("factory")
		for _, importPath := range st.NormalImportPaths() {
			f.ImportName(importPath, "")
		}
		for _, importInfo := range st.AliasImports() {
			f.ImportAlias(importInfo.Path(), importInfo.AliasName())
		}

		f.Type().Id(factoryTypeName).StructFunc(func(g *jen.Group) {
			for _, field := range st.Fields() {
				if multiFieldStruct, ok := typeNameToMultiFieldStruct[field.TypeName()]; ok {
					for _, innerField := range multiFieldStruct.Fields() {
						innerFieldTypeName := toTopLower(innerField.TypeName()) + toTopUpper(innerField.Name())
						innerFieldStatement := jen.Id(innerFieldTypeName)
						innerFieldStatement.Add(resolveTypeStatement(innerField, st))
						g.Add(innerFieldStatement)
					}
					continue
				}

				fieldStatement := jen.Id(field.Name())
				if tn, ok := entityToFactory[field.TypeName()]; ok {
					fieldStatement.Op("*").Qual("", tn)
					g.Add(fieldStatement)
					continue
				}
				fieldStatement.Add(resolveTypeStatement(field, st))
				g.Add(fieldStatement)
			}
		})

		err := f.Save(filepath.Join(outPath, toSnakeCase(factoryTypeName)+".go"))
		if err != nil {
			fmt.Printf("[Error] failed to factory code gen (%s): %v\n", st.Name(), err)
			isErr = true
		}
	}

	if isErr {
		return fmt.Errorf("[Error] failed to factory structs code generation.")
	}
	return nil
}

func resolveTypeStatement(fieldInfo *internal.FieldInfo, structInfo *internal.StructInfo) *jen.Statement {
	typeName := fieldInfo.TypeName()
	if tn, ok := fieldMapping[fieldInfo.TypeName()]; ok {
		typeName = tn
	}
	switch typeName {
	case typeUint64, typeUint32, typeUint16, typeUint8:
		return jen.Uint64()
	case typeArrayUint64, typeArrayUint32, typeArrayUint16, typeArrayUint8:
		return jen.Index().Uint64()
	case typeString:
		return jen.String()
	case typeArrayString:
		return jen.Index().String()
	default:
		statement := jen.Empty()
		if fieldInfo.IsPointer() {
			statement.Op("*")
		}
		if fieldInfo.IsArray() {
			statement.Index()
		}
		statement.Qual(structInfo.ResolvePkgPath(fieldInfo.TypePkgName()), fieldInfo.TypeName())
		return statement
	}
}

func parseMultiFieldStructs(rootPath string) (map[string]*internal.StructInfo, error) {
	ret := map[string]*internal.StructInfo{}
	for typeName, path := range multiFieldStructToPath {
		stInfo, err := ast.ParseSingleStruct(filepath.Join(rootPath, path), typeName)
		if err != nil {
			fmt.Println(err)
			continue
		}
		ret[typeName] = stInfo
	}
	return ret, nil
}

func toTopUpper(name string) string {
	b := &strings.Builder{}
	for i, r := range name {
		if i == 0 {
			b.WriteRune(unicode.ToUpper(r))
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}

func toTopLower(name string) string {
	b := &strings.Builder{}
	for i, r := range name {
		if i == 0 {
			b.WriteRune(unicode.ToLower(r))
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}

func toSnakeCase(pascal string) string {
	b := &strings.Builder{}
	for i, r := range pascal {
		if i == 0 {
			b.WriteRune(unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			b.WriteRune('_')
			b.WriteRune(unicode.ToLower(r))
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}
