package main

import (
	"fmt"
	"strings"
)

type Token struct {
	Datum string
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

func (token Token) String() string {
	return fmt.Sprintf("Token (%s)", token.Datum)
}
