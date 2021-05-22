package utils

import (
	"encoding/json"
	"io"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func Normalize(r string) string {
	t := transform.Chain(norm.NFKC, runes.Remove(runes.In(unicode.Mn)))
	s, _, _ := transform.String(t, r)
	return s
}

type apiResponseStruct struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func SendApiResponse(w io.Writer, respCode int, message string, data interface{}) {
	response := new(apiResponseStruct)

	response.StatusCode = respCode
	response.Message = message
	response.Data = data

	json.NewEncoder(w).Encode(response)
}
