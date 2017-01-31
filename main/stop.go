package main

// Stop filters 10 of the most common words out of the list
// of tokens. It returns the filtered list.
func Stop(tokens []Token) []Token {
	stops := map[string]bool{
		"the": true,
		"be": true,
		"and": true,
		"of": true,
		"a": true,
		"in": true,
		"to": true,
		"have": true,
		"it": true,
		"for": true,
	}

	toRet := []Token{}
	for _, token := range tokens {
		if _, ok := stops[string(token.Datum)]; !ok {
			toRet = append(toRet, token)
		}
	}
	return toRet
}
