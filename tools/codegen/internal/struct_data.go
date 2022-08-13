package internal

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
