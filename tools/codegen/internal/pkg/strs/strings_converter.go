package strs

import (
	"strings"
	"unicode"
)

func ToTopUpper(name string) string {
	b := &strings.Builder{}
	for i, r := range name {
		if i == 0 {
			b.WriteRune(unicode.ToUpper(r))
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}

func ToTopLower(name string) string {
	b := &strings.Builder{}
	for i, r := range name {
		if i == 0 {
			b.WriteRune(unicode.ToLower(r))
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}

func ToSnakeCase(pascal string) string {
	b := &strings.Builder{}
	for i, r := range pascal {
		if i == 0 {
			b.WriteRune(unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			b.WriteRune('_')
			b.WriteRune(unicode.ToLower(r))
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}
