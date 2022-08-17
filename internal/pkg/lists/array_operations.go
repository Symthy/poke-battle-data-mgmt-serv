package lists

import "reflect"

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
	for i, e := range array {
		rets[i] = f(e)
	}
	return rets
}

func Filter[T any](slice []T, f func(T) bool) []T {
	n := []T{}
	for _, e := range slice {
		if f(e) {
			n = append(n, e)
		}
	}
	return n
}

func AddLiterals[T uint16 | uint32 | uint64 | string](array []T, elems ...*T) {
	for _, e := range elems {
		if e != nil {
			array = append(array, *e)
		}
	}
}

func Add[T any](array []T, elems ...T) []T {
	rets := append(array, elems...)
	return rets
}

func ConvertTypeUint64To16(values []uint64) []uint16 {
	return Map(values, func(value uint64) uint16 { return uint16(value) })
}

func ConvertTypeUint16To64(values []uint16) []uint64 {
	return Map(values, func(value uint16) uint64 { return uint64(value) })
}
