package main

import (
	"fmt"
)

type Document struct {
	Id       int
	Name     string
	Contents []Token
}

func (doc Document) String() string {
	return fmt.Sprintf("Document #%d ('%s') %d", doc.Id, doc.Name, len(doc.Contents))
}

func DocumentFactory() func(string, []Token) Document {
	var ordinal = 0
	return func(name string, contents []Token) Document {
		defer func() { ordinal += 1 }()
		return Document{ordinal, name, contents}
	}
}
