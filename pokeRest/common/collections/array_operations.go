package collections

import (
	"reflect"

	"github.com/Symthy/PokeRest/pokeRest/infrastructure"
)

func ListContains(list interface{}, elem interface{}) bool {
	listV := reflect.ValueOf(list)

	if listV.Kind() == reflect.Slice {
		for i := 0; i < listV.Len(); i++ {
			item := listV.Index(i).Interface()
			// 型変換可能か確認する
			if !reflect.TypeOf(elem).ConvertibleTo(reflect.TypeOf(item)) {
				continue
			}
			// 型変換する
			target := reflect.ValueOf(elem).Convert(reflect.TypeOf(item)).Interface()
			// 等価判定をする
			if ok := reflect.DeepEqual(item, target); ok {
				return true
			}
		}
	}
	return false
}

func ListMap[TS infrastructure.ISchema[TD], TD infrastructure.IDomain](schemas []TS) []TD {
	domains := make([]TD, len(schemas), len(schemas))
	for i, s := range schemas {
		domains[i] = s.ConvertToDomain()
	}
	return domains
}
