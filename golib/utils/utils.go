package utils

import (
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
