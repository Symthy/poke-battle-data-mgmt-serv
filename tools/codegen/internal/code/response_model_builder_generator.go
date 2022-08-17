package code

import (
	"fmt"
	"path/filepath"

	"github.com/Symthy/PokeRest/tools/codegen/internal"
	"github.com/Symthy/PokeRest/tools/codegen/internal/pkg/strs"
	"github.com/dave/jennifer/jen"
)

const serverPkgPath = "github.com/Symthy/PokeRest/internal/adapters/rest/autogen/server"

func GenerateResponseModelBuilderStruct(homePath, outPath string, responseModelStructs []*internal.StructInfo) error {
	isErr := false
	for _, responseModelStruct := range responseModelStructs {
		f := jen.NewFile("dto")
		f.ImportName(serverPkgPath, "")

		funcStatements := []*jen.Statement{}
		builderTypeName := responseModelStruct.Name() + "Builder"
		for _, field := range responseModelStruct.Fields() {
			typeStatement := resolveTypeStatementReplacedPkgName(field, responseModelStruct, serverPkgPath)
			funcStatement := buildSetterMethodStatementForEmbeddedStruct("b", builderTypeName,
				field.Name(), field.Name(), field.TypeName(), typeStatement, field, responseModelStruct.Name())
			funcStatements = append(funcStatements, funcStatement)
		}

		f.Type().Id(builderTypeName).Struct(jen.Op("*").Qual(serverPkgPath, responseModelStruct.Name()))

		f.Add(buildBuilderConstructorStatement(responseModelStruct.Name(), builderTypeName))

		for _, funcStatement := range funcStatements {
			f.Add(funcStatement)
		}

		err := f.Save(filepath.Join(outPath, strs.ToSnakeCase(builderTypeName)+".gen.go"))
		if err != nil {
			fmt.Printf("[Error] failed to simple builder code gen (%s): %v\n", responseModelStruct.Name(), err)
			isErr = true
		}
	}

	if isErr {
		return fmt.Errorf("failed to response model builder structs code generation.")
	}
	fmt.Println("success to response model builder structs code generation.")
	return nil
}

func buildSetterMethodStatementForEmbeddedStruct(
	receiverName, targetTypeName, methodName, fieldName, fieldTypeName string, typeStatement *jen.Statement,
	fieldInfo *internal.FieldInfo, embeddedStructName string,
) *jen.Statement {
	structMemberName := fieldName
	if fieldInfo.IsEmbedded() {
		structMemberName = fieldTypeName
	}

	methodName = strs.ToTopUpper(methodName)
	if len(methodName) > 2 && methodName[:2] == "Is" {
		methodName = "Set" + methodName
	}
	return jen.Func().
		Params(jen.Id(receiverName).Op("*").Id(targetTypeName)).       // pointer receiver
		Id(methodName).                                                // func
		Params(jen.Id(strs.ToTopLower(fieldName)).Add(typeStatement)). // arguments
		Op("*").Qual("", targetTypeName).                              // return type
		Block(
			jen.Id(receiverName).Op(".").Id(embeddedStructName).Op(".").Id(structMemberName).Op("=").Id(strs.ToTopLower(fieldName)),
			jen.Return(jen.Id(receiverName)),
		)
}
