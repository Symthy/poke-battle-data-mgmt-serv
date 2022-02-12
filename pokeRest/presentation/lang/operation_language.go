package lang

import (
	"net/http"

	"golang.org/x/text/language"
)

type OperationLanguage struct {
	lang string
}

func NewRequestLanguage(req http.Request) OperationLanguage {
	return OperationLanguage{
		lang: resolveLanguage(req.Header),
	}
}

func resolveLanguage(header http.Header) string {
	lang := header.Get("Accept-Language")
	if lang != language.Japanese.String() {
		lang = language.English.String()
	}
	return lang
}

func (l OperationLanguage) Lang() string {
	return l.lang
}
