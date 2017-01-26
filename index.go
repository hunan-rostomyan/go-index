package main

import (
	"fmt"
)

func main() {
	col := NewLocalCollection("/Users/hunan/experiments/index/docs-bosh", "(.*).html.md.erb$")

	// Build the vocabulary of the entire collection.
	vocab := make(map[Token]bool)
	for _, doc := range col.Documents() {
		for _, token := range doc.Contents {
			if _, ok := vocab[token]; !ok {
				vocab[token] = true
			}
		}
	}

	fmt.Printf("Vocabulary size: %d", len(vocab))


	// Explore the counts of a specific document.
	doc := col.Documents()[3]
	keys, invCounts, _ := Count(doc.Contents)
	for i, key := range keys {
		if i > 2 {
			break;
		}
		fmt.Println(key)
		for _, val := range invCounts[key] {
			fmt.Printf("\t%s\n", val)
		}
	}
}
