package tests

import (
	"fmt"
	"testing"
	. "github.com/hunan-rostomyan/go-index/main"
)

func TestDocumentFactory(t *testing.T) {
	name := "Master and Margarita"
	contents := []Token{}

	NewDocument := DocumentFactory()

	doc := NewDocument(name, contents)
	doc2 := NewDocument(name, contents)

	// Ensure document ids are incremented.
	if doc.Id != 0 {
		t.Error(fmt.Sprintf("expected doc1 id to be 0 but was %d", doc.Id))
	}
	if doc2.Id != 1 {
		t.Error(fmt.Sprintf("expected doc2 id to be 1 but was %d", doc2.Id))
	}

	if doc.Name != name {
		t.Error(fmt.Sprintf("expected name '%s' found '%s'", name, doc.Name))
	}
	expectedLength := len(contents)
	actualLength := len(doc.Contents)
	if actualLength != actualLength {
		t.Error(fmt.Sprintf(
			"expected contents to be of size %d but was of size %d",
			expectedLength, actualLength))
	}
}
