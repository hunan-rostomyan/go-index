package main

import (
	"strings"
)

type Token struct {
	Datum string `json:"datum"`
}

func (token Token) String() string {
	return token.Datum
}

func Tokenize(text string) []Token {
	tokens := []Token{}
	for _, line := range strings.Split(text, "\n") {
		for _, datum := range strings.Split(line, " ") {
			if len(datum) != 0 {
				tokens = append(tokens, Token{datum})
			}
		}
	}
	return tokens
}
