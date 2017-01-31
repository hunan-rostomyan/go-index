package main

import (
	"regexp"
)

// Depunct removes punctuation from text. Assumes text is
// all lowercase. Returns text with all unacceptable characters
// replaced by spaces.
func Depunct(text string) string {
	re := regexp.MustCompile("[^a-z-_[:xdigit:]]")
	return re.ReplaceAllString(text, " ")
}
