package test

import (
	"reflect"
	"strings"
)

type ObjectStructure struct {
	fields []string
	values []string
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

func (o ObjectStructure) Values() []string {
	return o.values
}

func ConvertToObjectStructure(s interface{}) ObjectStructure {
	reflectValue := reflect.ValueOf(s)
	reflectType := reflect.TypeOf(s)
	objectStructure := NewObjectStructure(reflectType.NumField())
	for i := 0; i < reflectType.NumField(); i++ {
		field := reflectType.Field(i)
		value := reflectValue.FieldByName(field.Name)

		trimName := strings.TrimSpace(field.Name)
		topStr := trimName[0:]
		fieldName := strings.Replace(trimName, topStr, strings.ToLower(topStr), 1)
		objectStructure.fields = append(objectStructure.fields, fieldName)
		objectStructure.values = append(objectStructure.values, value.String())
	}
	return objectStructure
}
