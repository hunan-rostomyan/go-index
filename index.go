package main

func main() {
	col := NewLocalCollection("/Users/hunan/experiments/index/docs-bosh", "(.*).html.md.erb$")

	// Prepare the indexing tables
	documents := col.Documents()
	num_of_docs := 0
	vocab := make(map[Token]bool)
	tf_table := make(map[int]map[Token]int)
	df_table := make(map[Token]map[int]bool)

	// For each document
	for _, doc := range documents {

		num_of_docs += 1

		// Grab its tokens
		tokens := doc.Contents

		// Compute its token frequencies
		counts := Count(tokens)

		// Allocate an entry to hold its token frequencies
		tf_table[doc.Id] = make(map[Token]int)

		// For each token in the document:
		for _, token := range tokens {

			// Add to the vocabulary, unless already present
			if _, ok := vocab[token]; !ok {
				vocab[token] = true
			}

			// Compute the term frequency
			tf_table[doc.Id][token] = tf(token, counts)

			// Compute the document frequency
			_, ok := df_table[token]
			if !ok {
				df_table[token] = make(map[int]bool)
			} else {
				_, ok := df_table[token][doc.Id]
				if !ok {
					df_table[token][doc.Id] = true
				}
			}
		}
	}

	// Index
	index := make(map[Token]map[string]float64)
	for token := range vocab {
		index[token] = make(map[string]float64)
		for _, doc := range documents {
			//tokens := doc.Contents
			_tf := tf_table[doc.Id][token]
			if _tf == 0 {
				continue
			}
			_idf := idf(token, num_of_docs, df_table)
			if _idf == 0 {
				continue
			}
			_tfidf := float64(_tf) * _idf
			index[token][doc.Name] = _tfidf
		}
	}
}
