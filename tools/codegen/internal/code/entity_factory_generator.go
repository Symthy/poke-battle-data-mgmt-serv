package code

import (
	"fmt"
	"path/filepath"

	"github.com/Symthy/PokeRest/tools/codegen/internal"
	"github.com/Symthy/PokeRest/tools/codegen/internal/ast"
	"github.com/Symthy/PokeRest/tools/codegen/internal/pkg/strs"
	"github.com/dave/jennifer/jen"
)

const (
	factoryTypeSuffix = "Input"
)

var (
	entityToFactory = map[string]string{
		"Ability":                     "Ability" + factoryTypeSuffix,
		"BattleRecord":                "BattleRecord" + factoryTypeSuffix,
		"BattleOpponentParty":         "BattleOpponentParty" + factoryTypeSuffix,
		"Season":                      "Season" + factoryTypeSuffix,
		"HeldItem":                    "HeldItem" + factoryTypeSuffix,
		"Move":                        "Move" + factoryTypeSuffix,
		"Party":                       "Party" + factoryTypeSuffix,
		"PartyTag":                    "PartyTag" + factoryTypeSuffix,
		"PartyBattleResult":           "PartyBattleResult" + factoryTypeSuffix,
		"Pokemon":                     "Pokemon" + factoryTypeSuffix,
		"TrainedPokemon":              "TraindPokemon" + factoryTypeSuffix,
		"TrainedPokemonAdjustment":    "TraindPokemonAdjustment" + factoryTypeSuffix,
		"TrainedPokemonAttackTarget":  "TraindPokemonAttackTarget" + factoryTypeSuffix,
		"TrainedPokemonDefenseTarget": "TraindPokemonDefenseTarget" + factoryTypeSuffix,
		"User":                        "User" + factoryTypeSuffix,
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
		"PokemonTypeSet":      "internal/domain/value/pokemon_type_set.go",
		"PokemonAbilityIdSet": "internal/domain/value/pokemon_ability_id_set.go",
		"EffortValues":        "internal/domain/value/effort_values.go",
		"PokemonMoveIdSet":    "internal/domain/value/pokemon_move_id_set.go",
	}
)

func GenerateEntityFactoryStructs(
	rootPath, homePath, outPath string, entityStructs []*internal.StructInfo,
) error {
	typeNameToMultiFieldStruct, err := parseMultiFieldStructs(rootPath)
	if err != nil {
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
		addImportStatement(f, st)
		addEntityFactoryStructStatement(f, factoryTypeName, st, typeNameToMultiFieldStruct)

		err := f.Save(filepath.Join(outPath, strs.ToSnakeCase(factoryTypeName)+".gen.go"))
		if err != nil {
			fmt.Printf("[Error] failed to entity factory code gen (%s): %v\n", st.Name(), err)
			isErr = true
		}
	}

	if isErr {
		return fmt.Errorf("failed to entity factory structs code generation.")
	}
	fmt.Println("success to entity factory structs code generation.")
	return nil
}

func addEntityFactoryStructStatement(
	f *jen.File, factoryTypeName string, entityStruct *internal.StructInfo,
	typeNameToMultiFieldStruct map[string]*internal.StructInfo,
) {
	fieldStatements := []*jen.Statement{}
	funcStatements := []*jen.Statement{}
	for _, field := range entityStruct.Fields() {
		if multiFieldStruct, ok := typeNameToMultiFieldStruct[field.TypeName()]; ok {
			for _, innerField := range multiFieldStruct.Fields() {
				fieldName := ""
				typeName := field.TypeName()
				if !field.IsEmbedded() {
					fieldName = strs.ToTopLower(typeName) + strs.ToTopUpper(innerField.Name())
				}
				fieldStatement := jen.Id(fieldName)
				fieldTypeStatement := resolveEntityFactoryTypeStatement(innerField, entityStruct)
				fieldStatement.Add(fieldTypeStatement)
				fieldStatements = append(fieldStatements, fieldStatement)

				funcName := strs.ToTopUpper(fieldName)
				funcStatement := buildSetterMethodStatement("b", factoryTypeName, funcName, fieldName, typeName, fieldTypeStatement, field)
				funcStatements = append(funcStatements, funcStatement)
			}
			continue
		}

		if convertedTypeName, ok := entityToFactory[field.TypeName()]; ok {
			fieldName := strs.ToTopLower(field.Name())
			fieldStatement := jen.Id(fieldName)
			typeName := convertedTypeName
			typeStatement := jen.Empty().Op("*").Qual("", typeName)
			fieldStatement.Add(typeStatement)
			fieldStatements = append(fieldStatements, fieldStatement)

			funcName := strs.ToTopUpper(typeName)[:len(typeName)-len(factoryTypeSuffix)]
			if field.IsEmbedded() {
				fieldName = strs.ToTopLower(typeName)
			}
			funcStatement := buildSetterMethodStatement("b", factoryTypeName, funcName, fieldName, typeName, typeStatement, field)
			funcStatements = append(funcStatements, funcStatement)
			continue
		}

		fieldName := strs.ToTopLower(field.Name())
		fieldStatement := jen.Id(fieldName)
		typeName := field.TypeName()
		typeStatement := resolveEntityFactoryTypeStatement(field, entityStruct)
		fieldStatement.Add(typeStatement)
		fieldStatements = append(fieldStatements, fieldStatement)

		funcName := strs.ToTopUpper(field.Name())
		if field.IsEmbedded() {
			funcName = strs.ToTopUpper(typeName)
			fieldName = strs.ToTopLower(typeName)
		}
		funcStatement := buildSetterMethodStatement("b", factoryTypeName, funcName, fieldName, typeName, typeStatement, field)
		funcStatements = append(funcStatements, funcStatement)
	}

	f.Type().Id(factoryTypeName).StructFunc(func(g *jen.Group) {
		for _, fieldStatement := range fieldStatements {
			g.Add(fieldStatement)
		}
	})

	f.Add(buildBuilderConstructorStatement(entityStruct.Name(), factoryTypeName))
	for _, funcStatement := range funcStatements {
		f.Add(funcStatement)
	}
}

func resolveEntityFactoryTypeStatement(fieldInfo *internal.FieldInfo, structInfo *internal.StructInfo) *jen.Statement {
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
		return resolveTypeStatement(fieldInfo, structInfo)
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
