package main

import (
	"math"
)

func tf(term Token, counts map[Token]int) int {
	return counts[term]
}

func idf(term Token, num_of_docs int, df_table map[Token]map[int]bool) float64 {
	df := len(df_table[term])
	if df == 0 {
		return 0
	}
	return math.Log10(float64(num_of_docs) / float64(df))
}
