package main

import (
	"math"
)

type Dftable map[Token]map[int]bool

// Idf computes the inverse document frequency of token w.r.t the
// given document frequency table.
func idf(token Token, num_of_docs int, df_table Dftable) float64 {
	df := len(df_table[token])
	if df == 0 {
		return 0
	}
	return math.Log10(float64(num_of_docs) / float64(df))
}
