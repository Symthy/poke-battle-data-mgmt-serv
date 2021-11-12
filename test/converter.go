package test

import (
	"reflect"
	"strings"
)

type ObjectStructure struct {
	fields []string
	values []interface{}
}

func NewObjectStructure(length int) ObjectStructure {
	return ObjectStructure{
		//fields: make([]string, length),
		//values: make([]string, length),
	}
}

func (o ObjectStructure) Fields() []string {
	return o.fields
}

func (o ObjectStructure) Values() []interface{} {
	return o.values
}

func Convert(s interface{}) ObjectStructure {
	return ConvertInEnum(s, []int{})
}

func ConvertInEnum(s interface{}, enumIndexes []int) ObjectStructure {
	reflectValue := reflect.ValueOf(s)
	reflectType := reflect.TypeOf(s)
	objectStructure := NewObjectStructure(reflectType.NumField())
	for i := 0; i < reflectType.NumField(); i++ {
		field := reflectType.Field(i)
		value := reflectValue.FieldByName(field.Name)

		trimName := strings.TrimSpace(field.Name)
		// 先頭を小文字にする必要があるかと思い実施したが全て小文字にしていた。
		// これが原因でsqlmockからabilityIdの値が渡らずnilになっていた
		//topStr := trimName[0:]
		//fieldName := strings.Replace(trimName, topStr, strings.ToLower(topStr), 1)
		objectStructure.fields = append(objectStructure.fields, trimName)
		if isContains := contains(i, enumIndexes); isContains {
			objectStructure.values = append(objectStructure.values, value.Convert(reflect.TypeOf([]byte{})).Interface())
		} else {
			objectStructure.values = append(objectStructure.values, value.Interface())
		}
	}
	return objectStructure
}

func contains(i int, indexes []int) bool {
	for _, v := range indexes {
		if i == v {
			return true
		}
	}
	return false
}
