package database

import "strings"

func ConvertToDbField(fields ...string) []string {
	var convertedFields = make([]string, len(fields))
	for i, field := range fields {
		convertedFields[i] = convertFieldName(field)
	}
	return convertedFields
}

const (
	upperA = rune('A')
	upperZ = rune('Z')
)

func convertFieldName(field string) string {
	ret := field
	for _, c := range field {
		if c >= upperA && c <= upperZ {
			lower := strings.ToLower(string(c))
			ret = strings.Replace(ret, string(c), "_"+lower, 1)
		}
	}
	return ret
}
