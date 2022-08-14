package internal

import "github.com/Symthy/PokeRest/pokeRest/common/lists"

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

func (s StructInfo) NormalImports() []*ImportInfo {
	return lists.Filter(s.imports, func(i *ImportInfo) bool {
		return i.alias == ""
	})
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
	alias string
	path  string
}

func NewImportInfo(alias string, path string) *ImportInfo {
	return &ImportInfo{
		alias: alias,
		path:  exculdeQuotes(path),
	}
}

func exculdeQuotes(quotedPath string) string {
	path := quotedPath
	if path[0] == '"' {
		path = path[1:]
	}
	if path[len(path)-1] == '"' {
		path = path[:len(path)-1]
	}
	return path
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
