package main

import (
	"math"
)

func Tokens(doc Document) []Token {
	return doc.Contents
}

func tf(term Token, counts map[Token]int) int {
	return counts[term]
}

func idf(term Token, num_of_docs int, df_table map[Token]map[int]bool) float64 {
	df := len(df_table[term])
	return math.Log10(float64(num_of_docs) / float64(df))
}
