package main

import (
	"regexp"
)

func Depunct(text string) string {
	re := regexp.MustCompile("[^a-z-_[:xdigit:]]")
	return re.ReplaceAllString(text, " ")
}
