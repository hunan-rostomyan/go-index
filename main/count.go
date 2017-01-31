package main

// Count tallies up the number of occurrences of tokens.
// Returns a map from tokens to their counts.
func Count(tokens []Token) map[Token]int {
	counts := make(map[Token]int)

	for _, tok := range tokens {
		found, item := findCount(&counts, tok)
		if !found {
			counts[tok] = 1
		} else {
			counts[item]++
		}
	}
	return counts
}

func findCount(counts *map[Token]int, token Token) (bool, Token) {
	for tok, _ := range *counts {
		if tok.Datum == token.Datum {
			return true, tok
		}
	}
	return false, Token{}
}
