package internal

import (
	"strings"

	"github.com/Symthy/PokeRest/pokeRest/common/lists"
)

type StructInfo struct {
	name    string
	imports []*ImportInfo
	fields  []*FieldInfo
}

func NewStructInfo(name string, imports []*ImportInfo, fields []*FieldInfo) *StructInfo {
	return &StructInfo{
		name:    name,
		imports: imports,
		fields:  fields,
	}
}

func (s StructInfo) Name() string {
	return s.name
}

func (s StructInfo) AliasImports() []*ImportInfo {
	return lists.Filter(s.imports, func(i *ImportInfo) bool {
		return i.alias != ""
	})
}

func (s StructInfo) NormalImportPaths() []string {
	filteredImports := lists.Filter(s.imports, func(i *ImportInfo) bool {
		return i.alias == ""
	})
	return lists.Map(filteredImports, func(i *ImportInfo) string {
		return i.path
	})
}

func (s StructInfo) ResolvePkgPath(pkgName string) string {
	if pkgName == "" {
		// byte等の標準型や同パッケージの型はそのままにする
		return pkgName
	}
	for _, importInfo := range s.imports {
		if importInfo.alias == pkgName || importInfo.pkgName == pkgName {
			return importInfo.path
		}
	}
	return pkgName
}

func (s StructInfo) Fields() []*FieldInfo {
	return s.fields
}

func (s StructInfo) NormalFields() []*FieldInfo {
	return lists.Filter(s.fields, func(f *FieldInfo) bool {
		return f.isEmbedded == false
	})
}

func (s StructInfo) FieldNames() []string {
	return lists.Map(s.NormalFields(), func(f *FieldInfo) string {
		return f.fieldName
	})
}

type ImportInfo struct {
	alias   string
	pkgName string
	path    string
}

func NewImportInfo(alias string, path string) *ImportInfo {
	unquotedPath := exculdeQuotes(path)
	return &ImportInfo{
		alias:   alias,
		pkgName: resolvePkgName(unquotedPath),
		path:    unquotedPath,
	}
}

func resolvePkgName(unquotedPath string) string {
	pkgElems := strings.Split(unquotedPath, "/")
	return pkgElems[len(pkgElems)-1]
}

func exculdeQuotes(quotedPath string) string {
	// ASTから取得したパスは”で括られているため外す
	path := quotedPath
	if path[0] == '"' {
		path = path[1:]
	}
	if path[len(path)-1] == '"' {
		path = path[:len(path)-1]
	}
	return path
}

func (i ImportInfo) AliasName() string {
	return i.alias
}

func (i ImportInfo) Path() string {
	return i.path
}

type FieldInfo struct {
	fieldName  string
	isEmbedded bool
	typeInfo   *TypeInfo
}

func NewFieldInfo(fieldName string, typeInfo *TypeInfo) *FieldInfo {
	return &FieldInfo{
		fieldName:  fieldName,
		typeInfo:   typeInfo,
		isEmbedded: false,
	}
}

func NewEmbeddedFieldInfo(typeInfo *TypeInfo) *FieldInfo {
	return &FieldInfo{
		fieldName:  "",
		typeInfo:   typeInfo,
		isEmbedded: true,
	}
}

func (f FieldInfo) Name() string {
	return f.fieldName
}

func (f FieldInfo) TypeName() string {
	return f.typeInfo.typeName
}

func (f FieldInfo) TypePkgName() string {
	return f.typeInfo.pkgName
}

func (f FieldInfo) IsPointer() bool {
	return f.typeInfo.isPointer
}

func (f FieldInfo) IsArray() bool {
	return f.typeInfo.isArray
}

func (f FieldInfo) IsEmbedded() bool {
	return f.isEmbedded
}

type TypeInfo struct {
	isArray   bool
	isPointer bool
	pkgName   string
	typeName  string
}

func NewStandardTypeInfo(typeName string) *TypeInfo {
	return &TypeInfo{
		isPointer: false,
		pkgName:   "",
		typeName:  typeName,
	}
}

func NewTypeInfo(packageName string, typeName string) *TypeInfo {
	return &TypeInfo{
		pkgName:  packageName,
		typeName: typeName,
	}
}

func NewPointerTypeInfo(packageName string, typeName string) *TypeInfo {
	return &TypeInfo{
		pkgName:   packageName,
		typeName:  typeName,
		isPointer: true,
		isArray:   false,
	}
}

func NewArrayTypeInfo(packageName string, typeName string) *TypeInfo {
	return &TypeInfo{
		pkgName:   packageName,
		typeName:  typeName,
		isPointer: false,
		isArray:   true,
	}
}

func NewPointerArrayTypeInfo(packageName string, typeName string) *TypeInfo {
	return &TypeInfo{
		pkgName:   packageName,
		typeName:  typeName,
		isPointer: true,
		isArray:   true,
	}
}
