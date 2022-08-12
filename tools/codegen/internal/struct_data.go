package internal

type StructInfo struct {
	fieldName string
	types     []*TypeInfo
}

type TypeInfo struct {
	isPointer bool
	pkgName   string
	typeName  string
}

func NewStandardTypeInfo(typeName string) *TypeInfo {
	return &TypeInfo{
		isPointer: false,
	}
}
