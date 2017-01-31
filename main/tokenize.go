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

// Text2Tokens transforms a string into a list of Tokens. Expects text
// to be a string of lines joined by newlines. It splits the text into
// lines, lowercases them, and removes punctuation. The list of non-stop
// tokens is returned.
func Text2Tokens(text string) []Token {
	tokens := []Token{}

	// Tokenize
	for lineNum, line := range strings.Split(text, "\n") {

		// Lowercase
		text = strings.ToLower(text)

		// Remove punctuation
		line = Depunct(line)

		for _, datum := range strings.Split(line, " ") {
			if len(datum) != 0 {
				tokens = append(tokens, Token{datum, lineNum})
			}
		}
	}

	// Remove stop words
	tokens = Stop(tokens)

	return tokens
}
