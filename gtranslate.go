package gtranslate

import (
	"golang.org/x/text/language"
	"net/http"
)

// FromTo is a util struct to pass as parameter to indicate how to translate
type FromTo struct {
	From string
	To   string
}

// Translate translate a text using native tags offer by go language
func Translate(client *http.Client, text string, from language.Tag, to language.Tag) (string, error) {
	translated, err := translate(client, text, from.String(), to.String(), false)
	if err != nil {
		return "", err
	}

	return translated, nil
}

// TranslateWithFromTo translate a text with simple params as string
func TranslateWithFromTo(client *http.Client, text string, fromto FromTo) (string, error) {
	translated, err := translate(client, text, fromto.From, fromto.To, true)
	if err != nil {
		return "", err
	}

	return translated, nil

}
