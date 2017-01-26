package main

import (
	"sort"
)

// Tally the tokens in the decreasing order of frequency. Sorting is unstable.
// Returns the sorted keys, the inverted counts, and the total number of occurrences.
func Count(tokens []Token) ([]int, map[int][]Token, int) {
	// Count the occurrences.
	counts := make(map[Token]int)
	for _, token := range tokens {
		_, ok := counts[token];
		if !ok {
			counts[token] = 1
		} else {
			counts[token] += 1
		}
	}

	// Invert the counts.
	invCounts := make(map[int][]Token)
	for k, v := range counts {
		invCounts[v] = append(invCounts[v], k)
	}

	// Accomulate the keys.
	var keys []int
	for k := range invCounts {
		keys = append(keys, k)
	}

	// Sort the keys in decreasing order.
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	return keys, invCounts, len(tokens)
}

func Freq(m int, n int) float32 {
	return float32(m) / float32(n)
}
