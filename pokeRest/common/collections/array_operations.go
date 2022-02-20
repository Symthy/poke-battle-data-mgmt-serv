package collections

import (
	"reflect"
)

func ContainsInList(list interface{}, elem interface{}) bool {
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

func MapForList[TI any, TO any](array []TI, mapFunc func(TI) TO) []TO {
	rets := make([]TO, len(array))
	for _, e := range array {
		rets = append(rets, mapFunc(e))
	}
	return rets
}

func AddsToList[T any](array *[]T, elems ...*T) {
	for _, e := range elems {
		addToList(array, e)
	}
}

func addToList[T any](array *[]T, elem *T) {
	if elem != nil {
		*array = append(*array, *elem)
	}
}
