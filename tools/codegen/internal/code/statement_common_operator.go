package code

import (
	"github.com/Symthy/PokeRest/tools/codegen/internal"
	"github.com/Symthy/PokeRest/tools/codegen/internal/pkg/strs"
	"github.com/dave/jennifer/jen"
)

func addImportStatement(f *jen.File, st *internal.StructInfo) {
	for _, importPath := range st.NormalImportPaths() {
		f.ImportName(importPath, "")
	}
	for _, importInfo := range st.AliasImports() {
		f.ImportAlias(importInfo.Path(), importInfo.AliasName())
	}
}

func resolveTypeStatement(fieldInfo *internal.FieldInfo, structInfo *internal.StructInfo) *jen.Statement {
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

func resolveTypeStatementReplacedPkgName(
	fieldInfo *internal.FieldInfo, structInfo *internal.StructInfo, replacedPkgName string,
) *jen.Statement {
	statement := jen.Empty()
	if fieldInfo.IsPointer() {
		statement.Op("*")
	}
	if fieldInfo.IsArray() {
		statement.Index()
	}
	var pkgName string
	switch fieldInfo.TypeName() {
	case typeUint8, typeUint16, typeUint32, typeUint64, typeUint,
		typeArrayUint8, typeArrayUint16, typeArrayUint32, typeArrayUint64, typeArrayUint,
		typeInt8, typeInt16, typeInt32, typeInt64, typeInt,
		typeArrayInt8, typeArrayInt16, typeArrayInt32, typeArrayInt64, typeArrayInt,
		typeFloat32, typeFloat64, typeArrayFloat32, typeArrayFloat64,
		typeByte, typeArrayByte, typeString, typeArrayString, typeBool:
		pkgName = ""
	default:
		pkgName = replacedPkgName
	}
	statement.Qual(pkgName, fieldInfo.TypeName())
	return statement
}

func buildBuilderConstructorStatement(baseStructName, typeName string) *jen.Statement {
	return jen.Func().Id("New"+baseStructName+"Builder").
		Params().
		Op("*").Qual("", typeName).
		Block(
			jen.Return(jen.Op("&").Qual("", typeName).Block()),
		)
}

func buildSetterMethodStatement(
	receiverName, targetTypeName, methodName, fieldName, fieldTypeName string, typeStatement *jen.Statement, fieldInfo *internal.FieldInfo,
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
			jen.Id(receiverName).Op(".").Id(structMemberName).Op("=").Id(strs.ToTopLower(fieldName)),
			jen.Return(jen.Id(receiverName)),
		)
}
