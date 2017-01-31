package main

import (
	"strings"
)

// Token represents a word in a Document. It may also contain meta
// information about the word, such as which documents it belongs to,
// which lines (w.r.t. a particular document) it occurs on, and so on.
type Token struct {
	Datum string `json:"datum"`
	LineNum int
}

func (token Token) String() string {
	return token.Datum
}

// TextToTokens transforms a string into a list of Tokens. Expects text
// to be a string of lines joined by newlines. It splits the text into
// lines, lowercases them, and removes punctuation. The list of non-stop
// tokens is returned.
func TextToTokens(text string) []Token {
	tokens := []Token{}

	// Tokenize
	for lineNum, line := range strings.Split(text, "\n") {

		// Lowercase
		line = strings.ToLower(line)

		// Remove punctuation
		line = Depunct(line)

		for _, datum := range strings.Split(line, " ") {
			if len(datum) > 1 {  // ignore one-letter words
				tokens = append(tokens, Token{datum, lineNum})
			}
		}
	}

	// Remove stop words
	tokens = Stop(tokens)

	return tokens
}
