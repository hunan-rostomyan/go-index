package main

import (
	"fmt"
)

// Document represents a file to be indexed. Each document
// has a unique identifier, a name, and consists of a list
// of Tokens.
type Document struct {
	Id       int
	Name     string
	Contents []Token
}

// DocumentFactory closes over an ordinal (initialized to 0 and used as a
// document identifier), returning a factory for making documents given a
// name and a list of tokens. Each time the factory is called, the ordinal
// is incremented, ensuring that each document has a unique identifier.
func DocumentFactory() func(string, []Token) Document {
	var ordinal = 0
	return func(name string, contents []Token) Document {
		defer func() { ordinal += 1 }()
		return Document{ordinal, name, contents}
	}
}

func (doc Document) String() string {
	return fmt.Sprintf("Document #%d ('%s')", doc.Id, doc.Name)
}
