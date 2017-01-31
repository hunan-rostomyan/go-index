package main

import (
	"bytes"
	"fmt"
	"encoding/json"
)

type Dftable map[Token]map[int]bool
type Index map[Token]map[string]float64
type Tftable map[int]map[Token]int
type Vocab map[Token]bool

// BuildIndex extracts from documents an index of terms and their
// relevance to each of the documents they appear in.
func BuildIndex(documents []Document, numOfDocs int) Index {

	// Prepare the indexing tables
	vocab := make(Vocab)
	tf_table := make(Tftable)
	df_table := make(Dftable)

	// For each document
	for _, doc := range documents {

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
			tf_table[doc.Id][token] = counts[token]

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

	// Index, Rank
	var index Index = make(Index)

	for token := range vocab {
		index[token] = make(map[string]float64)
		for _, doc := range documents {
			_tf := tf_table[doc.Id][token]
			if _tf == 0 {
				continue
			}
			_idf := idf(token, numOfDocs, df_table)
			if _idf == 0 {
				continue
			}
			_tfidf := float64(_tf) * _idf
			index[token][doc.Name] = _tfidf
		}
	}
	return index
}

// From https://gist.github.com/mdwhatcott/8dd2eef0042f7f1c0cd8
func (index Index) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("{")
	size := len(index)
	i := 0
	for key, value := range index {
		jsonValue, err := json.Marshal(value)
		if err != nil {
			return nil, err
		}
		buffer.WriteString(fmt.Sprintf("\"%s\":%s", key, jsonValue))
		i += 1
		if i < size {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("}")
	return buffer.Bytes(), nil
}
