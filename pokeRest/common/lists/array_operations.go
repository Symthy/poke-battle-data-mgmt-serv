package lists

import (
	"reflect"
)

func Contains(list interface{}, elem interface{}) bool {
	// Todo: refactor
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

func ContainsAll[T any](array []T, elems ...T) bool {
	for _, e := range elems {
		if !Contains(array, e) {
			return false
		}
	}
	return true
}

func Map[TI any, TO any](array []TI, f func(TI) TO) []TO {
	if array == nil {
		return []TO{}
	}
	rets := make([]TO, len(array))
	for _, e := range array {
		rets = append(rets, f(e))
	}
	return rets
}

func Filter[T any](slice []T, f func(T) bool) []T {
	var n []T
	for _, e := range slice {
		if f(e) {
			n = append(n, e)
		}
	}
	return n
}

func AddsToList[T any](array []T, elems ...*T) {
	for _, e := range elems {
		addToList(array, e)
	}
}

func addToList[T any](array []T, elem *T) {
	if elem != nil {
		array = append(array, *elem)
	}
}
