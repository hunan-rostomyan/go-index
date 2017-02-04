package main

import (
	"bytes"
	"fmt"
	"encoding/json"
)

type TokenCounts map[string]int
type Index map[string]TokenCounts

// BuildIndex extracts from documents an index of terms and their
// relevance to each of the documents they appear in.
func BuildIndex(documents []Document, numOfDocs int) Index {
	var index Index = make(Index)

	// For each document
	for _, doc := range documents {

		// Create an entry for it in the index
		index[doc.Name] = make(map[string]int)

		// Grab its tokens
		tokens := doc.Contents

		// Compute its token frequencies
		counter := NewCounter(tokens)
		for _, counterEntry := range counter.Top(0, "desc") {
			index[doc.Name][counterEntry.Word] = counterEntry.Count
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
func (index Index) MarshalJSON() ([]byte, error) {
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
