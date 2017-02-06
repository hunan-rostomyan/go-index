package main

import (
	"bytes"
	"fmt"
	"encoding/json"
	"math"
)

type TokenCounts map[string]int
type InvertedIndex map[string]map[string]float64

// TfTable associates a term and a document with the number of times
// the term occurs in the document.
type TfTable map[string]map[string]int

// DcTable associates terms with their document counts.
type DcTable map[string]map[string]int

// BuildInvertedIndex extracts from documents an index of terms and their
// relevance to each of the documents they appear in. It belongs to the
// family of tf.idf algorithms. Returns the inverted index.
func BuildInvertedIndex(documents []Document, numOfDocs int) InvertedIndex {	
	vocab := make(map[string]bool)
	var tfTable TfTable = make(TfTable)  // map[string]map[string]int
	var dcTable DcTable = make(DcTable)  // map[string]map[string]int

	// Collect statistics
	for _, doc := range documents {

		tokens := doc.Contents
		counter := NewCounter(tokens)
		tfTable[doc.Name] = make(map[string]int)

		for _, counterEntry := range counter.Top(0, "desc") {

			// Add term to the vocabulary, unless already present
			if _, ok := vocab[counterEntry.Word]; !ok {
				vocab[counterEntry.Word] = true
			}

			tfTable[doc.Name][counterEntry.Word] = counterEntry.Count

			// Clip the count, since we're interested in membership, not count.
			wordCount := counterEntry.Count
			if wordCount > 0 {
				wordCount = 1
			}
			dcTable[counterEntry.Word] = make(map[string]int)
			dcTable[counterEntry.Word][doc.Name] = wordCount
		}

	}

	// Index
	var index InvertedIndex = make(InvertedIndex)
	for term := range vocab {

		index[term] = make(map[string]float64)

		// the number of documents the term occurs in
		// (not document-dependent, so it's outside the loop)
		dc := 0
		for _, docCount := range dcTable[term] {
			dc += docCount
		}

		for _, doc := range documents {

			// term frequency
			tf := tfTable[doc.Name][term]
			if tf == 0 {
				continue
			}

			// inverse document frequency
			idf := math.Log10(float64(numOfDocs) / float64(dc))
			if idf == 0 {
				continue
			}

			// final relevance score
			index[term][doc.Name] = float64(tf) * idf
		}
	}
	return index
}

func (tokenCounts TokenCounts) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("{")
	i := 0
	for word, count := range tokenCounts {
		buffer.WriteString(fmt.Sprintf("\"%s\":%d", word, count))
		i += 1
		if i < len(tokenCounts) {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("}")
	return buffer.Bytes(), nil
}

// From https://gist.github.com/mdwhatcott/8dd2eef0042f7f1c0cd8
func (index InvertedIndex) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("{")
	i := 0
	for key, value := range index {
		jsonValue, err := json.Marshal(value)
		if err != nil {
			return nil, err
		}
		buffer.WriteString(fmt.Sprintf("\"%s\":%s", key, jsonValue))
		i += 1
		if i < len(index) {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("}")
	return buffer.Bytes(), nil
}

